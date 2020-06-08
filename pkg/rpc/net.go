package rpc

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Transport abstraction of RPC service functionalities
type Network struct {

}

//
func (n *Network)CreateServer(){

	listen, err := net.Listen("tcp", fmt.Sprintf("%v:%v",  n.Addr, n.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	model.RegisterFrisbeeServer(server, &Network{})

	if err := server.Serve(listen); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}

//
func (n *Network)CreateClient(serverAddr string) model.FrisbeeClient{

	conn, err := grpc.Dial(serverAddr)
	if err != nil{

	}
	defer conn.Close()

	client := model.NewFrisbeeClient(conn)

	return client
}

//
func (n *Network)Ping(ctx context.Context, in *model.CheckStatusRequest) (*model.CheckStatusReply, error) {
	return &model.CheckStatusReply{Status: "Pong"}, nil
}
