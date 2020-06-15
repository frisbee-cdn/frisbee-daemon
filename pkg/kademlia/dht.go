package kademlia

import (
	"context"
	"time"

	"github/frisbee-cdn/frisbee-daemon/internal"
	kb "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"

	config "github/frisbee-cdn/frisbee-daemon/internal"

	log "github.com/sirupsen/logrus"

	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"

	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
)

var logger = log.New()

//type mode int
//
//const (
//	modeServer mode = iota + 1
//	modeClient
//)

// FrisbeeNode represents the node inside our network.
type FrisbeeNode struct {
	*proto.Node

	self peer.ID

	cfg *config.Configuration

	service *rpc.NetworkService

	createdAt time.Time

	ctx context.Context

	routingTable *kb.RoutingTable

	auto internal.ModeOpt
}

// New initializes the Frisbee Node
func New(selfId peer.ID, port uint32, conf *config.Configuration) (*FrisbeeNode, error) {

	var cfg *config.Configuration

	if conf != nil {
		cfg = conf
	} else {
		cfg = config.Defaults
	}

	node := &FrisbeeNode{
		Node:      new(proto.Node),
		cfg:       cfg,
		createdAt: time.Now(),
		self:      selfId,
	}

	// Hash IP Address and create Identifier
	id, err := kb.HashKey(node.self)
	if err != nil {
		return nil, err
	}

	node.Id = id
	node.Port = cfg.Server.Port
	node.Addr = cfg.Server.Addr

	node.routingTable, err = kb.NewRoutingTable(cfg.BucketSize, id, time.Minute)

	if err != nil {
		logger.Fatalf("Failed creating routing table: ", err)
	}

	service, err := rpc.NewNetworkService(cfg.Server)

	if err != nil {
		panic(err)
	}

	node.service = service

	proto.RegisterFrisbeeServer(node.service.GetServer(), node)

	// Start service connections
	go node.service.Start()

	logger.Infof("Peer %x just started listening on: %v:%v", node.Id, node.Addr, node.Port)

	return node, nil
}

func (n *FrisbeeNode) Join(bootpNode string) {

	// logger.Printf("Trying to connect to: %s:%d...", addr, port)
	// client, err := n.service.Connect(fmt.Sprintf("%s:%d", addr, port))
	// if err != nil {
	// 	logger.Fatal("Failed to create client")
	// }
	// r, err := client.Ping(context.Background(), &proto.CheckStatusRequest{Message: "Ping"})
	// if err != nil {
	// 	logger.Fatalf("Node not alive: %v", err)
	// }
	// logger.Printf("Recevied from Bootstrap node: ", r.Status)
}

func (n *FrisbeeNode) shutdown() error {
	return nil
}

// RPC Interface Implementation
