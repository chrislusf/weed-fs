package broker

import (
	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/mq/pub_balancer"
	"github.com/seaweedfs/seaweedfs/weed/pb/mq_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PublisherToPubBalancer receives connections from brokers and collects stats
func (broker *MessageQueueBroker) PublisherToPubBalancer(stream mq_pb.SeaweedMessaging_ConnectToBalancerServer) error {
	if !broker.lockAsBalancer.IsLocked() {
		return status.Errorf(codes.Unavailable, "not current broker balancer")
	}
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	// process init message
	initMessage := req.GetInit()
	var brokerStats *pub_balancer.BrokerStats
	if initMessage != nil {
		var found bool
		brokerStats, found = broker.Balancer.Brokers.Get(initMessage.Broker)
		if !found {
			brokerStats = pub_balancer.NewBrokerStats()
			if !broker.Balancer.Brokers.SetIfAbsent(initMessage.Broker, brokerStats) {
				brokerStats, _ = broker.Balancer.Brokers.Get(initMessage.Broker)
			}
		}
	} else {
		return status.Errorf(codes.InvalidArgument, "balancer init message is empty")
	}
	defer func() {
		broker.Balancer.Brokers.Remove(initMessage.Broker)
	}()

	// process stats message
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if !broker.lockAsBalancer.IsLocked() {
			return status.Errorf(codes.Unavailable, "not current broker balancer")
		}
		if receivedStats := req.GetStats(); receivedStats != nil {
			brokerStats.UpdateStats(receivedStats)

			glog.V(4).Infof("broker %s stats: %+v", initMessage.Broker, brokerStats)
			glog.V(4).Infof("received stats: %+v", receivedStats)
		}
	}

	return nil
}
