package rpc

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	service "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

var logger = log.New()


// NetworkService abstraction of RPC service functionalities
type NetworkService struct {}

// CreateServer
func (n *NetworkService)CreateServer(addr string, port uint32){

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service.RegisterFrisbeeServer(server, n)

	//reflection.Register(server)

	if err := server.Serve(listen); err != nil{
		logger.Fatalf("Failed to serve: %v", err)
	}

}

// CreateClient
func (n *NetworkService)CreateClient(serverAddr string) service.FrisbeeClient{

	conn, err := grpc.Dial(serverAddr)
	if err != nil{
		logger.Fatalln("Failed creating client: ", err)
	}
	defer conn.Close()

	client := service.NewFrisbeeClient(conn)

	return client
}

// Ping
func (n *NetworkService)Ping(ctx context.Context, reqBody *service.CheckStatusRequest) (*service.CheckStatusReply, error) {
	return &service.CheckStatusReply{Status: "Pong"}, nil
}

// Store
func (n *NetworkService)Store(ctx context.Context, reqBody  *service.StoreRequest) ( *service.Error, error){
	return nil,nil
}

// FindNode
func (n *NetworkService)FindNode(ctx context.Context, reqBody *service.ID) ( *service.NodeResponse, error){
	return nil, nil
}

// FindValue
func (n *NetworkService)FindValue(ctx context.Context, reqBody *service.ID) ( *service.StorageResponse, error){
	return nil, nil
}