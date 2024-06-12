package command

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	hashicorpRaft "github.com/hashicorp/raft"

	"golang.org/x/exp/slices"

	"github.com/gorilla/mux"
	"github.com/seaweedfs/raft/protobuf"
	"github.com/spf13/viper"
	"google.golang.org/grpc/reflection"

	stats_collect "github.com/seaweedfs/seaweedfs/weed/stats"

	"github.com/seaweedfs/seaweedfs/weed/util/grace"

	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/pb"
	"github.com/seaweedfs/seaweedfs/weed/pb/master_pb"
	"github.com/seaweedfs/seaweedfs/weed/security"
	weed_server "github.com/seaweedfs/seaweedfs/weed/server"
	"github.com/seaweedfs/seaweedfs/weed/storage/backend"
	"github.com/seaweedfs/seaweedfs/weed/util"
)

var (
	m MasterOptions
)

type MasterOptions struct {
	port              *int
	portGrpc          *int
	ip                *string
	ipBind            *string
	metaFolder        *string
	peers             *string
	volumeSizeLimitMB *uint
	volumePreallocate *bool
	// pulseSeconds       *int
	defaultReplication *string
	garbageThreshold   *float64
	whiteList          *string
	disableHttp        *bool
	metricsAddress     *string
	metricsIntervalSec *int
	raftResumeState    *bool
	metricsHttpPort    *int
	heartbeatInterval  *time.Duration
	electionTimeout    *time.Duration
	raftHashicorp      *bool
	raftBootstrap      *bool
	username           *string
	password           *string
}

func init() {
	cmdMaster.Run = runMaster // break init cycle
	m.port = cmdMaster.Flag.Int("port", 9333, "http listen port")
	m.portGrpc = cmdMaster.Flag.Int("port.grpc", 0, "grpc listen port")
	m.ip = cmdMaster.Flag.String("ip", util.DetectedHostAddress(), "master <ip>|<server> address, also used as identifier")
	m.ipBind = cmdMaster.Flag.String("ip.bind", "", "ip address to bind to. If empty, default to same as -ip option.")
	m.metaFolder = cmdMaster.Flag.String("mdir", os.TempDir(), "data directory to store meta data")
	m.peers = cmdMaster.Flag.String("peers", "", "all master nodes in comma separated ip:port list, example: 127.0.0.1:9093,127.0.0.1:9094,127.0.0.1:9095")
	m.volumeSizeLimitMB = cmdMaster.Flag.Uint("volumeSizeLimitMB", 30*1000, "Master stops directing writes to oversized volumes.")
	m.volumePreallocate = cmdMaster.Flag.Bool("volumePreallocate", false, "Preallocate disk space for volumes.")
	// m.pulseSeconds = cmdMaster.Flag.Int("pulseSeconds", 5, "number of seconds between heartbeats")
	m.defaultReplication = cmdMaster.Flag.String("defaultReplication", "", "Default replication type if not specified.")
	m.garbageThreshold = cmdMaster.Flag.Float64("garbageThreshold", 0.3, "threshold to vacuum and reclaim spaces")
	m.whiteList = cmdMaster.Flag.String("whiteList", "", "comma separated Ip addresses having write permission. No limit if empty.")
	m.disableHttp = cmdMaster.Flag.Bool("disableHttp", false, "disable http requests, only gRPC operations are allowed.")
	m.metricsAddress = cmdMaster.Flag.String("metrics.address", "", "Prometheus gateway address <host>:<port>")
	m.metricsIntervalSec = cmdMaster.Flag.Int("metrics.intervalSeconds", 15, "Prometheus push interval in seconds")
	m.metricsHttpPort = cmdMaster.Flag.Int("metricsPort", 0, "Prometheus metrics listen port")
	m.raftResumeState = cmdMaster.Flag.Bool("resumeState", false, "resume previous state on start master server")
	m.heartbeatInterval = cmdMaster.Flag.Duration("heartbeatInterval", 300*time.Millisecond, "heartbeat interval of master servers, and will be randomly multiplied by [1, 1.25)")
	m.electionTimeout = cmdMaster.Flag.Duration("electionTimeout", 10*time.Second, "election timeout of master servers")
	m.raftHashicorp = cmdMaster.Flag.Bool("raftHashicorp", false, "use hashicorp raft")
	m.raftBootstrap = cmdMaster.Flag.Bool("raftBootstrap", false, "Whether to bootstrap the Raft cluster")
	m.username = cmdMaster.Flag.String("username", "", "username for authentication")
	m.password = cmdMaster.Flag.String("password", "", "password for authentication")
}

var cmdMaster = &Command{
	UsageLine: "master -port=9333",
	Short:     "start a master server",
	Long: `start a master server to provide volume=>location mapping service and sequence number of file ids

	The configuration file "security.toml" is read from ".", "$HOME/.seaweedfs/", "/usr/local/etc/seaweedfs/", or "/etc/seaweedfs/", in that order.

	The example security.toml configuration file can be generated by "weed scaffold -config=security"

  `,
}

var (
	masterCpuProfile = cmdMaster.Flag.String("cpuprofile", "", "cpu profile output file")
	masterMemProfile = cmdMaster.Flag.String("memprofile", "", "memory profile output file")
)

func runMaster(cmd *Command, args []string) bool {

	util.LoadConfiguration("security", false)
	util.LoadConfiguration("master", false)

	grace.SetupProfiling(*masterCpuProfile, *masterMemProfile)

	parent, _ := util.FullPath(*m.metaFolder).DirAndName()
	if util.FileExists(string(parent)) && !util.FileExists(*m.metaFolder) {
		os.MkdirAll(*m.metaFolder, 0755)
	}
	if err := util.TestFolderWritable(util.ResolvePath(*m.metaFolder)); err != nil {
		glog.Fatalf("Check Meta Folder (-mdir) Writable %s : %s", *m.metaFolder, err)
	}

	masterWhiteList := util.StringSplit(*m.whiteList, ",")
	if *m.volumeSizeLimitMB > util.VolumeSizeLimitGB*1000 {
		glog.Fatalf("volumeSizeLimitMB should be smaller than 30000")
	}

	go stats_collect.StartMetricsServer(*m.ipBind, *m.metricsHttpPort)
	startMaster(m, masterWhiteList)

	return true
}

func startMaster(masterOption MasterOptions, masterWhiteList []string) {

	backend.LoadConfiguration(util.GetViper())

	if *masterOption.portGrpc == 0 {
		*masterOption.portGrpc = 10000 + *masterOption.port
	}
	if *masterOption.ipBind == "" {
		*masterOption.ipBind = *masterOption.ip
	}

	myMasterAddress, peers := checkPeers(*masterOption.ip, *masterOption.port, *masterOption.portGrpc, *masterOption.peers)

	masterPeers := make(map[string]pb.ServerAddress)
	for _, peer := range peers {
		masterPeers[string(peer)] = peer
	}

	r := mux.NewRouter()
	ms := weed_server.NewMasterServer(r, masterOption.toMasterOption(masterWhiteList), masterPeers)
	listeningAddress := util.JoinHostPort(*masterOption.ipBind, *masterOption.port)
	glog.V(0).Infof("Start Seaweed Master %s at %s", util.Version(), listeningAddress)
	masterListener, masterLocalListener, e := util.NewIpAndLocalListeners(*masterOption.ipBind, *masterOption.port, 0)
	if e != nil {
		glog.Fatalf("Master startup error: %v", e)
	}

	// start raftServer
	metaDir := path.Join(*masterOption.metaFolder, fmt.Sprintf("m%d", *masterOption.port))
	raftServerOption := &weed_server.RaftServerOption{
		GrpcDialOption:    security.LoadClientTLS(util.GetViper(), "grpc.master"),
		Peers:             masterPeers,
		ServerAddr:        myMasterAddress,
		DataDir:           util.ResolvePath(metaDir),
		Topo:              ms.Topo,
		RaftResumeState:   *masterOption.raftResumeState,
		HeartbeatInterval: *masterOption.heartbeatInterval,
		ElectionTimeout:   *masterOption.electionTimeout,
		RaftBootstrap:     *masterOption.raftBootstrap,
	}
	var raftServer *weed_server.RaftServer
	var err error
	if *masterOption.raftHashicorp {
		if raftServer, err = weed_server.NewHashicorpRaftServer(raftServerOption); err != nil {
			glog.Fatalf("NewHashicorpRaftServer: %s", err)
		}
		ms.Topo.HashicorpRaftTransportManager = raftServer.TransportManager
	} else {
		raftServer, err = weed_server.NewRaftServer(raftServerOption)
		if raftServer == nil {
			glog.Fatalf("please verify %s is writable, see https://github.com/seaweedfs/seaweedfs/issues/717: %s", *masterOption.metaFolder, err)
		}
	}
	ms.SetRaftServer(raftServer)
	r.HandleFunc("/cluster/status", raftServer.StatusHandler).Methods("GET")
	r.HandleFunc("/cluster/healthz", raftServer.HealthzHandler).Methods("GET", "HEAD")
	if *masterOption.raftHashicorp {
		r.HandleFunc("/raft/stats", raftServer.StatsRaftHandler).Methods("GET")
	}
	// starting grpc server
	grpcPort := *masterOption.portGrpc
	grpcL, grpcLocalL, err := util.NewIpAndLocalListeners(*masterOption.ipBind, grpcPort, 0)
	if err != nil {
		glog.Fatalf("master failed to listen on grpc port %d: %v", grpcPort, err)
	}
	grpcS := pb.NewGrpcServer(security.LoadServerTLS(util.GetViper(), "grpc.master"))
	master_pb.RegisterSeaweedServer(grpcS, ms)
	if *masterOption.raftHashicorp {
		raftServer.TransportManager.Register(grpcS)
	} else {
		protobuf.RegisterRaftServer(grpcS, raftServer)
	}
	reflection.Register(grpcS)
	glog.V(0).Infof("Start Seaweed Master %s grpc server at %s:%d", util.Version(), *masterOption.ipBind, grpcPort)
	if grpcLocalL != nil {
		go grpcS.Serve(grpcLocalL)
	}
	go grpcS.Serve(grpcL)

	timeSleep := 1500 * time.Millisecond
	if !*masterOption.raftHashicorp {
		go func() {
			time.Sleep(timeSleep)

			ms.Topo.RaftServerAccessLock.RLock()
			isEmptyMaster := ms.Topo.RaftServer.Leader() == "" && ms.Topo.RaftServer.IsLogEmpty()
			if isEmptyMaster && isTheFirstOne(myMasterAddress, peers) && ms.MasterClient.FindLeaderFromOtherPeers(myMasterAddress) == "" {
				raftServer.DoJoinCommand()
			}
			ms.Topo.RaftServerAccessLock.RUnlock()
		}()
	}

	go ms.MasterClient.KeepConnectedToMaster()

	// start http server
	var (
		clientCertFile,
		certFile,
		keyFile string
	)
	useTLS := false
	useMTLS := false

	if viper.GetString("https.master.key") != "" {
		useTLS = true
		certFile = viper.GetString("https.master.cert")
		keyFile = viper.GetString("https.master.key")
	}

	if viper.GetString("https.master.ca") != "" {
		useMTLS = true
		clientCertFile = viper.GetString("https.master.ca")
	}

	httpS := &http.Server{Handler: r}
	if masterLocalListener != nil {
		go httpS.Serve(masterLocalListener)
	}

	if useMTLS {
		httpS.TLSConfig = security.LoadClientTLSHTTP(clientCertFile)
	}

	if useTLS {
		go httpS.ServeTLS(masterListener, certFile, keyFile)
	} else {
		go httpS.Serve(masterListener)
	}

	grace.OnInterrupt(ms.Shutdown)
	grace.OnInterrupt(grpcS.Stop)
	grace.OnReload(func() {
		if ms.Topo.HashicorpRaft != nil && ms.Topo.HashicorpRaft.State() == hashicorpRaft.Leader {
			ms.Topo.HashicorpRaft.LeadershipTransfer()
		}
	})
	select {}
}

func checkPeers(masterIp string, masterPort int, masterGrpcPort int, peers string) (masterAddress pb.ServerAddress, cleanedPeers []pb.ServerAddress) {
	glog.V(0).Infof("current: %s:%d peers:%s", masterIp, masterPort, peers)
	masterAddress = pb.NewServerAddress(masterIp, masterPort, masterGrpcPort)
	cleanedPeers = pb.ServerAddresses(peers).ToAddresses()

	hasSelf := false
	for _, peer := range cleanedPeers {
		if peer.ToHttpAddress() == masterAddress.ToHttpAddress() {
			hasSelf = true
			break
		}
	}

	if !hasSelf {
		cleanedPeers = append(cleanedPeers, masterAddress)
	}
	if len(cleanedPeers)%2 == 0 {
		glog.Fatalf("Only odd number of masters are supported: %+v", cleanedPeers)
	}
	return
}

func isTheFirstOne(self pb.ServerAddress, peers []pb.ServerAddress) bool {
	slices.SortFunc(peers, func(a, b pb.ServerAddress) int {
		return strings.Compare(string(a), string(b))
	})
	if len(peers) <= 0 {
		return true
	}
	return self == peers[0]
}

func (m *MasterOptions) toMasterOption(whiteList []string) *weed_server.MasterOption {
	masterAddress := pb.NewServerAddress(*m.ip, *m.port, *m.portGrpc)
	return &weed_server.MasterOption{
		Master:            masterAddress,
		MetaFolder:        *m.metaFolder,
		VolumeSizeLimitMB: uint32(*m.volumeSizeLimitMB),
		VolumePreallocate: *m.volumePreallocate,
		// PulseSeconds:            *m.pulseSeconds,
		DefaultReplicaPlacement: *m.defaultReplication,
		GarbageThreshold:        *m.garbageThreshold,
		WhiteList:               whiteList,
		DisableHttp:             *m.disableHttp,
		MetricsAddress:          *m.metricsAddress,
		MetricsIntervalSec:      *m.metricsIntervalSec,
		Username:                *m.username,
		Password:                *m.password,
	}
}
