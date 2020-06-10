package kademlia

import (
	"context"
	"fmt"
	"time"

	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"

	log "github.com/sirupsen/logrus"
	"github/frisbee-cdn/frisbee-daemon/internal"

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

	config    *internal.Configuration

	net rpc.NetworkService

	createdAt time.Time

	ctx context.Context

	routingTable *peer.RoutingTable

	auto internal.ModeOpt


}


// BootStrap creates a new node in the newtork
func BootStrap(config *internal.Configuration) *FrisbeeNode {

	node, err := New(config)
	if err != nil {
		panic(fmt.Errorf("FrisbeeNode startup failed"))
	}

	return node
}

func New(cfg *internal.Configuration) (*FrisbeeNode, error) {

	node := &FrisbeeNode{
		Node:         new(model.Node),
		config:       cfg,
		createdAt:	  time.Now(),
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

	if err != nil{
		logger.Fatalf("Failed creating routing table: ",err)
	}

	service := rpc.NetworkService{}

	service.CreateServer(cfg.Server.Addr, cfg.Server.Port)

	logger.Infof("Peer %o just started listening on: %v:%v", node.Id, node.Addr, node.Port)
	return node, nil
}

func (n *FrisbeeNode) join(addr string, port uint32) error {
	return nil
}

func (n *FrisbeeNode) leave() error {

	return nil
}

func (n *FrisbeeNode) stabilize() {

}

