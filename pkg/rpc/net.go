package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc/reflection"

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

// ClientConn represents the open connection of the client
type ClientConn struct {
	hostAddr string
	client   proto.FrisbeeProtocolClient
	conn     *grpc.ClientConn
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

	reflection.Register(n.server)
	if err := n.server.Serve(n.sock); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

// Connect
func Connect(serverAddr string) (*ClientConn, error) {

	c, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clt := proto.NewFrisbeeProtocolClient(c)

	cc := &ClientConn{hostAddr: serverAddr, client: clt, conn: c}
	return cc, nil
}

// GetServer returns the server
func (n *NetworkService) GetServer() *grpc.Server {
	return n.server
}

// Close is used to shutdown the client connection
func (c *ClientConn) Close() {
	c.conn.Close()
}

// GetClient returns the client
func (c *ClientConn) GetClient() proto.FrisbeeProtocolClient {
	return c.client
}
