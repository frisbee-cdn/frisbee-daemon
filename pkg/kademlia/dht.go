package kademlia

import (
	"context"
	"fmt"
	"time"

	"github/frisbee-cdn/frisbee-daemon/internal"
	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"

	config "github/frisbee-cdn/frisbee-daemon/internal"

	log "github.com/sirupsen/logrus"

	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
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
	*model.Node

	ownerID string

	cfg *config.Configuration

	service *rpc.NetworkService

	createdAt time.Time

	ctx context.Context

	routingTable *peer.RoutingTable

	auto internal.ModeOpt
}

// BootStrap creates a new node in the newtork
// func BootStrap(ownId string) *FrisbeeNode {

// 	node, err := New(config)
// 	if err != nil {
// 		panic(fmt.Errorf("FrisbeeNode startup failed"))
// 	}

// 	return node
// }

// New initializes the Frisbee Node
func New(ownId string, port uint32, conf *config.Configuration) (*FrisbeeNode, error) {

	var cfg *config.Configuration

	if conf != nil {
		cfg = conf
	} else {
		cfg = config.Defaults
	}

	node := &FrisbeeNode{
		Node:      new(model.Node),
		service:   rpc.NewNetworkService(cfg.Server),
		cfg:       cfg,
		createdAt: time.Now(),
		ownerID:   ownId,
	}

	// Hash IP Address and create Identifier
	id, err := peer.HashKey(cfg.Server.Addr)
	if err != nil {
		return nil, err
	}

	node.Id = id
	node.Port = cfg.Server.Port
	node.Addr = cfg.Server.Addr

	node.routingTable, err = peer.NewRoutingTable(cfg.BucketSize, id, time.Minute)

	if err != nil {
		logger.Fatalf("Failed creating routing table: ", err)
	}

	logger.Infof("Peer %x just started listening on: %v:%v", node.Id, node.Addr, node.Port)
	node.service.Start()

	return node, nil
}

func (n *FrisbeeNode) Join(addr string, port uint32) {

	logger.Printf("Trying to connect to: %s:%d...", addr, port)
	client, err := n.service.CreateClient(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		logger.Fatal("Failed to create client")
	}
	r, err := client.Ping(context.Background(), &model.CheckStatusRequest{Message: "Ping"})
	if err != nil {
		logger.Fatalf("Node not alive: %v", err)
	}
	logger.Printf("Recevied from Bootstrap node: ", r.Status)
}

func (n *FrisbeeNode) leave() error {

	return nil
}

func (n *FrisbeeNode) stabilize() {

}
