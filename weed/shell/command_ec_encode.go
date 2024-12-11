package shell

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/pb"

	"google.golang.org/grpc"

	"github.com/seaweedfs/seaweedfs/weed/operation"
	"github.com/seaweedfs/seaweedfs/weed/pb/master_pb"
	"github.com/seaweedfs/seaweedfs/weed/pb/volume_server_pb"
	"github.com/seaweedfs/seaweedfs/weed/storage/erasure_coding"
	"github.com/seaweedfs/seaweedfs/weed/storage/needle"
	"github.com/seaweedfs/seaweedfs/weed/wdclient"
)

func init() {
	Commands = append(Commands, &commandEcEncode{})
}

type commandEcEncode struct {
}

func (c *commandEcEncode) Name() string {
	return "ec.encode"
}

func (c *commandEcEncode) Help() string {
	return `apply erasure coding to a volume

	ec.encode [-collection=""] [-fullPercent=95 -quietFor=1h]
	ec.encode [-collection=""] [-volumeId=<volume_id>]

	This command will:
	1. freeze one volume
	2. apply erasure coding to the volume
	3. (optionally) re-balance encoded shards across multiple volume servers

	The erasure coding is 10.4. So ideally you have more than 14 volume servers, and you can afford
	to lose 4 volume servers.

	If the number of volumes are not high, the worst case is that you only have 4 volume servers,
	and the shards are spread as 4,4,3,3, respectively. You can afford to lose one volume server.

	If you only have less than 4 volume servers, with erasure coding, at least you can afford to
	have 4 corrupted shard files.

	Re-balancing algorithm:
	` + ecBalanceAlgorithmDescription
}

func (c *commandEcEncode) HasTag(CommandTag) bool {
	return false
}

func (c *commandEcEncode) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {

	encodeCommand := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	volumeId := encodeCommand.Int("volumeId", 0, "the volume id")
	collection := encodeCommand.String("collection", "", "the collection name")
	fullPercentage := encodeCommand.Float64("fullPercent", 95, "the volume reaches the percentage of max volume size")
	quietPeriod := encodeCommand.Duration("quietFor", time.Hour, "select volumes without no writes for this period")
	parallelize := encodeCommand.Bool("parallelize", true, "parallelize operations whenever possible")
	forceChanges := encodeCommand.Bool("force", false, "force the encoding even if the cluster has less than recommended 4 nodes")
	shardReplicaPlacement := encodeCommand.String("shardReplicaPlacement", "", "replica placement for EC shards, or master default if empty")
	applyBalancing := encodeCommand.Bool("rebalance", false, "re-balance EC shards after creation")

	if err = encodeCommand.Parse(args); err != nil {
		return nil
	}
	if err = commandEnv.confirmIsLocked(args); err != nil {
		return
	}
	rp, err := parseReplicaPlacementArg(commandEnv, *shardReplicaPlacement)
	if err != nil {
		return err
	}

	// collect topology information
	topologyInfo, _, err := collectTopologyInfo(commandEnv, 0)
	if err != nil {
		return err
	}

	if !*forceChanges {
		var nodeCount int
		eachDataNode(topologyInfo, func(dc DataCenterId, rack RackId, dn *master_pb.DataNodeInfo) {
			nodeCount++
		})
		if nodeCount < erasure_coding.ParityShardsCount {
			glog.V(0).Infof("skip erasure coding with %d nodes, less than recommended %d nodes", nodeCount, erasure_coding.ParityShardsCount)
			return nil
		}
	}

	var volumeIds []needle.VolumeId
	if vid := needle.VolumeId(*volumeId); vid != 0 {
		// volumeId is provided
		volumeIds = append(volumeIds, vid)
	} else {
		// apply to all volumes in the collection
		volumeIds, err = collectVolumeIdsForEcEncode(commandEnv, *collection, *fullPercentage, *quietPeriod)
		if err != nil {
			return err
		}
	}

	var collections []string
	if *collection != "" {
		collections = []string{*collection}
	} else {
		// TODO: should we limit this to collections associated with the provided volume ID?
		collections, err = ListCollectionNames(commandEnv, false, true)
		if err != nil {
			return err
		}
	}

	// encode all requested volumes...
	for _, vid := range volumeIds {
		if err = doEcEncode(commandEnv, *collection, vid); err != nil {
			return fmt.Errorf("ec encode for volume %d: %v", vid, err)
		}
	}
	// ...then re-balance ec shards.
	if err := EcBalance(commandEnv, collections, "", rp, *parallelize, *applyBalancing); err != nil {
		return fmt.Errorf("re-balance ec shards for collection(s) %v: %v", collections, err)
	}

	return nil
}

func doEcEncode(commandEnv *CommandEnv, collection string, vid needle.VolumeId) error {
	if !commandEnv.isLocked() {
		return fmt.Errorf("lock is lost")
	}

	// find volume location
	locations, found := commandEnv.MasterClient.GetLocationsClone(uint32(vid))
	if !found {
		return fmt.Errorf("volume %d not found", vid)
	}

	// fmt.Printf("found ec %d shards on %v\n", vid, locations)

	// mark the volume as readonly
	if err := markVolumeReplicasWritable(commandEnv.option.GrpcDialOption, vid, locations, false, false); err != nil {
		return fmt.Errorf("mark volume %d as readonly on %s: %v", vid, locations[0].Url, err)
	}

	// generate ec shards
	if err := generateEcShards(commandEnv.option.GrpcDialOption, vid, collection, locations[0].ServerAddress()); err != nil {
		return fmt.Errorf("generate ec shards for volume %d on %s: %v", vid, locations[0].Url, err)
	}

	return nil
}

func generateEcShards(grpcDialOption grpc.DialOption, volumeId needle.VolumeId, collection string, sourceVolumeServer pb.ServerAddress) error {

	fmt.Printf("generateEcShards %s %d on %s ...\n", collection, volumeId, sourceVolumeServer)

	err := operation.WithVolumeServerClient(false, sourceVolumeServer, grpcDialOption, func(volumeServerClient volume_server_pb.VolumeServerClient) error {
		_, genErr := volumeServerClient.VolumeEcShardsGenerate(context.Background(), &volume_server_pb.VolumeEcShardsGenerateRequest{
			VolumeId:   uint32(volumeId),
			Collection: collection,
		})
		return genErr
	})

	return err

}

// TODO: delete this (now unused) shard spread logic.
func spreadEcShards(commandEnv *CommandEnv, volumeId needle.VolumeId, collection string, existingLocations []wdclient.Location, parallelCopy bool) (err error) {

	allEcNodes, totalFreeEcSlots, err := collectEcNodes(commandEnv)
	if err != nil {
		return err
	}

	if totalFreeEcSlots < erasure_coding.TotalShardsCount {
		return fmt.Errorf("not enough free ec shard slots. only %d left", totalFreeEcSlots)
	}
	allocatedDataNodes := allEcNodes
	if len(allocatedDataNodes) > erasure_coding.TotalShardsCount {
		allocatedDataNodes = allocatedDataNodes[:erasure_coding.TotalShardsCount]
	}

	// calculate how many shards to allocate for these servers
	allocatedEcIds := balancedEcDistribution(allocatedDataNodes)

	// ask the data nodes to copy from the source volume server
	copiedShardIds, err := parallelCopyEcShardsFromSource(commandEnv.option.GrpcDialOption, allocatedDataNodes, allocatedEcIds, volumeId, collection, existingLocations[0], parallelCopy)
	if err != nil {
		return err
	}

	// unmount the to be deleted shards
	err = unmountEcShards(commandEnv.option.GrpcDialOption, volumeId, existingLocations[0].ServerAddress(), copiedShardIds)
	if err != nil {
		return err
	}

	// ask the source volume server to clean up copied ec shards
	err = sourceServerDeleteEcShards(commandEnv.option.GrpcDialOption, collection, volumeId, existingLocations[0].ServerAddress(), copiedShardIds)
	if err != nil {
		return fmt.Errorf("source delete copied ecShards %s %d.%v: %v", existingLocations[0].Url, volumeId, copiedShardIds, err)
	}

	// ask the source volume server to delete the original volume
	for _, location := range existingLocations {
		fmt.Printf("delete volume %d from %s\n", volumeId, location.Url)
		err = deleteVolume(commandEnv.option.GrpcDialOption, volumeId, location.ServerAddress(), false)
		if err != nil {
			return fmt.Errorf("deleteVolume %s volume %d: %v", location.Url, volumeId, err)
		}
	}

	return err

}

func parallelCopyEcShardsFromSource(grpcDialOption grpc.DialOption, targetServers []*EcNode, allocatedEcIds [][]uint32, volumeId needle.VolumeId, collection string, existingLocation wdclient.Location, parallelCopy bool) (actuallyCopied []uint32, err error) {

	fmt.Printf("parallelCopyEcShardsFromSource %d %s\n", volumeId, existingLocation.Url)

	var wg sync.WaitGroup
	shardIdChan := make(chan []uint32, len(targetServers))
	copyFunc := func(server *EcNode, allocatedEcShardIds []uint32) {
		defer wg.Done()
		copiedShardIds, copyErr := oneServerCopyAndMountEcShardsFromSource(grpcDialOption, server,
			allocatedEcShardIds, volumeId, collection, existingLocation.ServerAddress())
		if copyErr != nil {
			err = copyErr
		} else {
			shardIdChan <- copiedShardIds
			server.addEcVolumeShards(volumeId, collection, copiedShardIds)
		}
	}
	cleanupFunc := func(server *EcNode, allocatedEcShardIds []uint32) {
		if err := unmountEcShards(grpcDialOption, volumeId, pb.NewServerAddressFromDataNode(server.info), allocatedEcShardIds); err != nil {
			fmt.Printf("unmount aborted shards %d.%v on %s: %v\n", volumeId, allocatedEcShardIds, server.info.Id, err)
		}
		if err := sourceServerDeleteEcShards(grpcDialOption, collection, volumeId, pb.NewServerAddressFromDataNode(server.info), allocatedEcShardIds); err != nil {
			fmt.Printf("remove aborted shards %d.%v on target server %s: %v\n", volumeId, allocatedEcShardIds, server.info.Id, err)
		}
		if err := sourceServerDeleteEcShards(grpcDialOption, collection, volumeId, existingLocation.ServerAddress(), allocatedEcShardIds); err != nil {
			fmt.Printf("remove aborted shards %d.%v on existing server %s: %v\n", volumeId, allocatedEcShardIds, existingLocation.ServerAddress(), err)
		}
	}

	// maybe parallelize
	for i, server := range targetServers {
		if len(allocatedEcIds[i]) <= 0 {
			continue
		}

		wg.Add(1)
		if parallelCopy {
			go copyFunc(server, allocatedEcIds[i])
		} else {
			copyFunc(server, allocatedEcIds[i])
		}
	}
	wg.Wait()
	close(shardIdChan)

	if err != nil {
		for i, server := range targetServers {
			if len(allocatedEcIds[i]) <= 0 {
				continue
			}
			cleanupFunc(server, allocatedEcIds[i])
		}
		return nil, err
	}

	for shardIds := range shardIdChan {
		actuallyCopied = append(actuallyCopied, shardIds...)
	}

	return
}

func balancedEcDistribution(servers []*EcNode) (allocated [][]uint32) {
	allocated = make([][]uint32, len(servers))
	allocatedShardIdIndex := uint32(0)
	serverIndex := rand.Intn(len(servers))
	for allocatedShardIdIndex < erasure_coding.TotalShardsCount {
		if servers[serverIndex].freeEcSlot > 0 {
			allocated[serverIndex] = append(allocated[serverIndex], allocatedShardIdIndex)
			allocatedShardIdIndex++
		}
		serverIndex++
		if serverIndex >= len(servers) {
			serverIndex = 0
		}
	}

	return allocated
}

func collectVolumeIdsForEcEncode(commandEnv *CommandEnv, selectedCollection string, fullPercentage float64, quietPeriod time.Duration) (vids []needle.VolumeId, err error) {
	// collect topology information
	topologyInfo, volumeSizeLimitMb, err := collectTopologyInfo(commandEnv, 0)
	if err != nil {
		return
	}

	quietSeconds := int64(quietPeriod / time.Second)
	nowUnixSeconds := time.Now().Unix()

	fmt.Printf("collect volumes quiet for: %d seconds and %.1f%% full\n", quietSeconds, fullPercentage)

	vidMap := make(map[uint32]bool)
	eachDataNode(topologyInfo, func(dc DataCenterId, rack RackId, dn *master_pb.DataNodeInfo) {
		for _, diskInfo := range dn.DiskInfos {
			for _, v := range diskInfo.VolumeInfos {
				// ignore remote volumes
				if v.RemoteStorageName != "" && v.RemoteStorageKey != "" {
					continue
				}
				if v.Collection == selectedCollection && v.ModifiedAtSecond+quietSeconds < nowUnixSeconds {
					if float64(v.Size) > fullPercentage/100*float64(volumeSizeLimitMb)*1024*1024 {
						if good, found := vidMap[v.Id]; found {
							if good {
								if diskInfo.FreeVolumeCount < 2 {
									glog.V(0).Infof("skip %s %d on %s, no free disk", v.Collection, v.Id, dn.Id)
									vidMap[v.Id] = false
								}
							}
						} else {
							if diskInfo.FreeVolumeCount < 2 {
								glog.V(0).Infof("skip %s %d on %s, no free disk", v.Collection, v.Id, dn.Id)
								vidMap[v.Id] = false
							} else {
								vidMap[v.Id] = true
							}
						}
					}
				}
			}
		}
	})

	for vid, good := range vidMap {
		if good {
			vids = append(vids, needle.VolumeId(vid))
		}
	}

	return
}
