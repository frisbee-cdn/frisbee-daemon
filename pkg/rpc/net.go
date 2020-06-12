package rpc

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"

	config "github/frisbee-cdn/frisbee-daemon/internal"
	service "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

var logger = log.New()

// NetworkService abstraction of RPC service functionalities
type NetworkService struct {
	cfg    config.ServerConfiguration
	server *grpc.Server
}

// NewNetworkService creates a new service and initializes all server info
func NewNetworkService(cfg config.ServerConfiguration) *NetworkService {

	proto := &NetworkService{
		cfg:    cfg,
		server: grpc.NewServer(),
	}

	service.RegisterFrisbeeServer(proto.server, proto)

	reflection.Register(proto.server)

	return proto

}

// CreateServer
func (n *NetworkService) Start() {

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", n.cfg.Addr, n.cfg.Port))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	if err := n.server.Serve(listen); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

// CreateClient
// TODO: Rename CreateClient into Connect(context, address)
func (n *NetworkService) Connect(serverAddr string) (service.FrisbeeClient, error) {

	conn, err := grpc.Dial(serverAddr)
	if err != nil {
		return nil, nil
	}
	defer conn.Close()

	client := service.NewFrisbeeClient(conn)

	return client, nil
}

// Ping
func (n *NetworkService) Ping(ctx context.Context, reqBody *service.CheckStatusRequest) (*service.CheckStatusReply, error) {
	logger.Info("Ping: ", reqBody.GetMessage())
	return &service.CheckStatusReply{Status: "Pong"}, nil
}

// Store
func (n *NetworkService) Store(ctx context.Context, reqBody *service.StoreRequest) (*service.Error, error) {
	return nil, nil
}

// FindNode
func (n *NetworkService) FindNode(ctx context.Context, reqBody *service.ID) (*service.NodeResponse, error) {
	return nil, nil
}

// FindValue
func (n *NetworkService) FindValue(ctx context.Context, reqBody *service.ID) (*service.StorageResponse, error) {
	return nil, nil
}
