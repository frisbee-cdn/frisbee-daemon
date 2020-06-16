package kademlia

import (
	"context"
	"time"

	kb "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"

	config "github/frisbee-cdn/frisbee-daemon/internal"

	log "github.com/sirupsen/logrus"

	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"

	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
)

var logger = log.New()

// FrisbeeDHT represents the node inside our network.
type FrisbeeDHT struct {
	// self node triplet
	node *peer.Node

	parallelismDegree uint32

	cfg *config.Configuration

	service *rpc.NetworkService

	createdAt time.Time

	ctx context.Context

	routingTable *kb.RoutingTable
}

// New initializes the Frisbee Node
func New(selfID peer.ID, port uint32, conf *config.Configuration) (*FrisbeeDHT, error) {

	var cfg *config.Configuration

	if conf != nil {
		cfg = conf
	} else {
		cfg = config.Defaults
	}

	dht := &FrisbeeDHT{
		node:              peer.NewNode(cfg.Server.Addr, port, selfID),
		parallelismDegree: cfg.ParallelismDegree,
		cfg:               cfg,
		createdAt:         time.Now(),
	}
	// Hash IP Address and create Identifier
	id, err := kb.HashKey(selfID)
	if err != nil {
		return nil, err
	}

	dht.node.SelfKey = id
	dht.routingTable, err = kb.NewRoutingTable(cfg.BucketSize, id, time.Minute)

	if err != nil {
		logger.Fatalf("Failed creating routing table: ", err)
	}

	service, err := rpc.NewNetworkService(cfg.Server)

	if err != nil {
		panic(err)
	}

	dht.service = service

	proto.RegisterFrisbeeProtocolServer(dht.service.GetServer(), dht)

	// Start service connections
	go dht.service.Start()

	logger.Infof("Peer %x just started listening on: %v:%v", dht.node.SelfKey, dht.node.GetHostAddress(), dht.node.GetAddressPort())

	return dht, nil
}

func (n *FrisbeeDHT) Join() {

	// The first bootstrap node is selected for testing purposes

	/*TODO:
	1. Get The First BootstrapNode
	2. Ping The Bootstrap Node ( if everything goes ok, they will add each other so BN will knwo NN and vice versa)
	3. Sender Will perform FIND_NODE RPC to BN to find k clossest nodes from it
	4. Ping every node recieved to learn about them and they as well
	5. NN is now connected inside the network, and other nodes know about it
	6. *Optionally you can perform an iterative_finde_node to get a better idea of the nodes inside the network
	*/

}

func (n *FrisbeeDHT) shutdown() error {
	return nil
}

// RPC Interface Implementation
// NodeLookup
func (n *FrisbeeDHT) NodeLookup(ctx context.Context, target kb.ID, addr string, done chan []*proto.Node) {

	client, err := n.service.Connect(addr)
	if err != nil {
		done <- nil
		return
	}

	r, err := client.FindNode(ctx, &proto.ID{Id: target})
	if err != nil {
		done <- nil
		return
	}

	//TODO: Add Sender ID to routingTable
	//n.routingTable.ADD(true, false)

	done <- r.Nodes
}
