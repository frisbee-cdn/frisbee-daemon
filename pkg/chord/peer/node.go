package peer

import (
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
	"github/frisbee-cdn/frisbee-daemon/pkg/util"
)

// Node represents the node inside our network.
type Node struct {
	*model.Node

	Config    *internal.Configuration
	Transport rpc.Transport
}

// BootStrap creates a new node in the newtork
func BootStrap(config *internal.Configuration) *Node {

	node, err := newNode(config)
	if err != nil {
		panic(fmt.Errorf("Node startup failed"))
	}

	return node
}

func newNode(config *internal.Configuration) (*Node, error) {

	node := &Node{
		Node:   new(model.Node),
		Config: config,
	}

	// Hash IP Address and create Identifier
	id, err := util.HashKey(config.Server.Addr)
	if err != nil {
		return nil, err
	}
	node.Node.Id = id
	node.Node.Addr = config.Server.Addr

	return node, nil
}

func (n *Node) join() error {
	return nil
}

func (n *Node) leave() error {

	return nil
}

func (n *Node) stabilize() {

}

func (n *Node) findCloserNode() {

}

func findNode() {

}

func discoverPeer() {

}
