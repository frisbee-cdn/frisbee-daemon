package kademlia

import (
	"context"
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
	"time"
)

type mode int

const (
	modeServer mode = iota + 1
	modeClient
)

// Peer represents the node inside our network.
type Peer struct {
	*model.Node

	config    *internal.Configuration
	net rpc.Network

	createdAt time.Time

	ctx context.Context

	routingTable *peer.RoutingTable

	auto internal.ModeOpt
	mode mode

	bucketSize int


}


// BootStrap creates a new node in the newtork
func BootStrap(config *internal.Configuration) *Peer {

	node, err := New(config)
	if err != nil {
		panic(fmt.Errorf("Peer startup failed"))
	}

	return node
}

func New(cfg *internal.Configuration) (*Peer, error) {

	node := &Peer{
		Node:         new(model.Node),
		config:       cfg,
		createdAt:	  time.Now(),
		bucketSize:   cfg.BucketSize,
	}

	// Hash IP Address and create Identifier
	id, err := peer.HashKey(cfg.Server.Addr)
	if err != nil {
		return nil, err
	}

	node.Node.Id = id
	node.Node.Port = cfg.Server.Port
	node.Node.Addr = cfg.Server.Addr

	node.routingTable, err = peer.NewRoutingTable(cfg.BucketSize, id, time.Minute)

	if err != nil{

	}

	return node, nil
}

func (n *Peer) join() error {
	return nil
}

func (n *Peer) leave() error {

	return nil
}

func (n *Peer) stabilize() {

}

func (n *Peer) findCloserNode() {

}

func findNode() {

}

func discoverPeer() {

}


