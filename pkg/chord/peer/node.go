package peer

import (
	cfg "github/frisbee-cdn/frisbee-daemon/internal"
	trans "github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Node represents the node inside our network.
type Node struct {
	*model.Node

	Config    *cfg.Configuration
	Transport trans.Transport
}

// InitNode creates a new node in the newtork
func InitNode() {

}

func newNode(config *cfg.Configuration) (*Node, error) {

	node := &Node{
		Node:   new(model.Node),
		Config: config,
	}

	id, err := HashKey(config.Server.Addr)
	if err != nil {
		return nil, err
	}
	node.Node.Id = id
	node.Node.Addr = config.Server.Addr

	return node, nil
}

func (p *Node) join() error {
	return nil
}

func (p *Node) leave() error {

	return nil
}

func (p *Node) stabilize() {

}
