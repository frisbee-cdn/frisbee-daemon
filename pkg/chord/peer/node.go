package peer

import (
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
	"github/frisbee-cdn/frisbee-daemon/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Node represents the node inside our network.
type Node struct {
	*model.Node

	Config    *internal.Configuration
	Transport rpc.Transport

	Routes *RoutingTable
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
		Routes: NewRoutingTable(),
	}

	// Hash IP Address and create Identifier
	id, err := util.HashKey(config.Server.Addr)
	if err != nil {
		return nil, err
	}
	node.Node.Id = id
	node.Node.Port = config.Server.Port
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


func (n *Node)StartServer(){

	listen, err := net.Listen("tcp", fmt.Sprintf("%v:%v",  n.Addr, n.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	model.RegisterFrisbeeServer(server, n)

	if err := server.Serve(listen); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (n *Node) Ping(ctx context.Context, in *model.CheckStatusRequest) (*model.CheckStatusReply, error) {
	return &model.CheckStatusReply{Status: "Pong"}, nil
}
