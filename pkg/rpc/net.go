package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"

	config "github/frisbee-cdn/frisbee-daemon/internal"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

var logger = log.New()

// NetworkService abstraction of RPC service functionalities
type NetworkService struct {
	cfg    config.ServerConfiguration
	server *grpc.Server
	sock   *net.TCPListener
}

// NewNetworkService creates a new service and initializes all server info
func NewNetworkService(cfg config.ServerConfiguration) (*NetworkService, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port))
	if err != nil {
		return nil, err
	}

	proto := &NetworkService{
		sock:   listener.(*net.TCPListener),
		cfg:    cfg,
		server: grpc.NewServer(),
	}

	return proto, nil

}

// CreateServer
func (n *NetworkService) Start() {

	if err := n.server.Serve(n.sock); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

// Connect
func (n *NetworkService) Connect(serverAddr string) (proto.FrisbeeProtocolClient, error) {

	conn, err := grpc.Dial(serverAddr)
	if err != nil {
		return nil, nil
	}
	defer conn.Close()

	client := proto.NewFrisbeeProtocolClient(conn)

	return client, nil
}

// GetServer returns the server
func (n *NetworkService) GetServer() *grpc.Server {
	return n.server
}
