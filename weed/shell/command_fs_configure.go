package shell

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/chrislusf/seaweedfs/weed/filer"
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
	"github.com/chrislusf/seaweedfs/weed/storage/super_block"
	"io"
)

func init() {
	Commands = append(Commands, &commandFsConfigure{})
}

type commandFsConfigure struct {
}

func (c *commandFsConfigure) Name() string {
	return "fs.configure"
}

func (c *commandFsConfigure) Help() string {
	return `configure and apply storage options for each location

	# see the current configuration file content
	fs.configure

	# trying the changes and see the possible configuration file content
	fs.configure -locationPrefix=/my/folder -collection=abc
	fs.configure -locationPrefix=/my/folder -collection=abc -ttl=7d

	# example: configure adding only 1 physical volume for each bucket collection
	fs.configure -locationPrefix=/buckets/ -volumeGrowthCount=1

	# apply the changes
	fs.configure -locationPrefix=/my/folder -collection=abc -apply

	# delete the changes
	fs.configure -locationPrefix=/my/folder -delete -apply

`
}

func (c *commandFsConfigure) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {

	fsConfigureCommand := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	locationPrefix := fsConfigureCommand.String("locationPrefix", "", "path prefix, required to update the path-specific configuration")
	collection := fsConfigureCommand.String("collection", "", "assign writes to this collection")
	replication := fsConfigureCommand.String("replication", "", "assign writes with this replication")
	ttl := fsConfigureCommand.String("ttl", "", "assign writes with this ttl")
	diskType := fsConfigureCommand.String("disk", "", "[hdd|ssd|<tag>] hard drive or solid state drive or any tag")
	fsync := fsConfigureCommand.Bool("fsync", false, "fsync for the writes")
	isReadOnly := fsConfigureCommand.Bool("readOnly", false, "disable writes")
	dataCenter := fsConfigureCommand.String("dataCenter", "", "assign writes to this dataCenter")
	rack := fsConfigureCommand.String("rack", "", "assign writes to this rack")
	dataNode := fsConfigureCommand.String("dataNode", "", "assign writes to this dataNode")
	volumeGrowthCount := fsConfigureCommand.Int("volumeGrowthCount", 0, "the number of physical volumes to add if no writable volumes")
	isDelete := fsConfigureCommand.Bool("delete", false, "delete the configuration by locationPrefix")
	apply := fsConfigureCommand.Bool("apply", false, "update and apply filer configuration")
	if err = fsConfigureCommand.Parse(args); err != nil {
		return nil
	}

	fc, err := filer.ReadFilerConf(commandEnv.option.FilerAddress, commandEnv.option.GrpcDialOption, commandEnv.MasterClient)
	if err != nil {
		return err
	}

	if *locationPrefix != "" {
		infoAboutSimulationMode(writer, *apply, "-apply")
		locConf := &filer_pb.FilerConf_PathConf{
			LocationPrefix:    *locationPrefix,
			Collection:        *collection,
			Replication:       *replication,
			Ttl:               *ttl,
			Fsync:             *fsync,
			DiskType:          *diskType,
			VolumeGrowthCount: uint32(*volumeGrowthCount),
			ReadOnly:          *isReadOnly,
			DataCenter:        *dataCenter,
			Rack:              *rack,
			DataNode:          *dataNode,
		}

		// check replication
		if *replication != "" {
			_, err := super_block.NewReplicaPlacementFromString(*replication)
			if err != nil {
				return fmt.Errorf("parse replication %s: %v", *replication, err)
			}
		}

		// save it
		if *isDelete {
			fc.DeleteLocationConf(*locationPrefix)
		} else {
			fc.AddLocationConf(locConf)
		}
	}

	var buf2 bytes.Buffer
	fc.ToText(&buf2)

	fmt.Fprintf(writer, string(buf2.Bytes()))
	fmt.Fprintln(writer)

	if *apply {

		if err = commandEnv.WithFilerClient(false, func(client filer_pb.SeaweedFilerClient) error {
			return filer.SaveInsideFiler(client, filer.DirectoryEtcSeaweedFS, filer.FilerConfName, buf2.Bytes())
		}); err != nil && err != filer_pb.ErrNotFound {
			return err
		}

	}

	return nil

}

func infoAboutSimulationMode(writer io.Writer, forceMode bool, forceModeOption string) {
	if forceMode {
		return
	}
	fmt.Fprintf(writer, "Running in simulation mode. Use \"%s\" option to apply the changes.\n", forceModeOption)
}
