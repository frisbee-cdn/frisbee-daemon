package kademlia

import (
	"context"
	"fmt"
	"time"

	kb "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"

	config "github/frisbee-cdn/frisbee-daemon/internal"

	log "github.com/sirupsen/logrus"

	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"

	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"

	ds "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/storage"
)

var logger = log.New()

// FrisbeeDHT represents the node inside our network.
type FrisbeeDHT struct {
	// self node triplet
	node *peer.Node

	StringRepr string

	parallelismDegree uint32

	cfg *config.Configuration

	service *rpc.NetworkService

	createdAt time.Time

	ctx context.Context

	routingTable *kb.RoutingTable

	datastore *ds.MapDatastore
}

// New initializes the Frisbee Node
func New(selfID string, port uint32, conf *config.Configuration) (*FrisbeeDHT, error) {

	var cfg *config.Configuration

	if conf != nil {
		cfg = conf
	} else {
		cfg = config.Defaults
		cfg.Server.Port = port
	}

	// Hash IP Address and create Identifier
	id, err := peer.HashKey(selfID)
	if err != nil {
		return nil, err
	}

	dht := &FrisbeeDHT{
		node:              peer.NewNode(cfg.Server.Addr, port, id),
		parallelismDegree: cfg.ParallelismDegree,
		cfg:               cfg,
		createdAt:         time.Now(),
		datastore:         ds.NewMapDatastore(),
	}

	dht.node.ID = id
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

	go func() {
		ticker := time.NewTicker(cfg.RefreshTimeout)

		for {
			select {
			case <-ticker.C:
				fmt.Printf("New Refresh")
			}
		}
	}()

	logger.Infof("Peer %x just started listening on: %v:%v", dht.node.ID, dht.node.GetHostAddress(), dht.node.GetAddressPort())

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

	bootNode := n.cfg.DefaultBootstrapPeers[0]
	//ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)

	addr := fmt.Sprintf("%s:%d", bootNode.Addr, bootNode.Port)
	err := n.PingRequest(context.Background(), addr)
	if err != nil {
		logger.Fatalf("Error Pinging Node: %s", err)
	}

	clsPeers, err := n.FindNodeRequest(context.Background(), n.node.ID, addr)
	if err != nil {
		logger.Errorf("Error Finding Node: %s", err)
	}

	for _, p := range clsPeers {
		err := n.PingRequest(context.Background(), fmt.Sprintf("%s:%d", p.Addr, p.Port))
		if err != nil {
			logger.Fatalf("Error Pinging Node: %s", err)
		}
	}

	n.routingTable.PrintInfo()

}

func (n *FrisbeeDHT) shutdown() error {
	return nil
}

// RPC Interface Implementation
// NodeLookup
func (n *FrisbeeDHT) NodeLookup(ctx context.Context, target peer.NodeID, addr string, done chan []*proto.Node) {

	client, err := n.service.Connect(addr)
	if err != nil {
		done <- nil
		return
	}

	r, err := client.GetClient().FindNode(ctx, &proto.FindNodeRequest{Id: target})
	if err != nil {
		done <- nil
		return
	}

	//TODO: Add Sender ID to routingTable
	//n.routingTable.ADD(true, false)

	done <- r.Nodes
}
