package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Ping RPC Abstract Call
func (n *FrisbeeNode) Ping(ctx context.Context, reqBody *proto.CheckStatusRequest) (*proto.CheckStatusReply, error) {
	logger.Info("Ping: ", reqBody.GetMessage())
	return &proto.CheckStatusReply{Status: "Pong"}, nil
}

// Ping
func (n *FrisbeeNode) PingRequest(ctx context.Context, message string, addr string) error {

	client, err := n.service.Connect(addr)
	if err != nil {
		return err
	}

	r, err := client.Ping(ctx, &proto.CheckStatusRequest{Message: message})
	if err != nil {
		return err
	}
	logger.Infof("Received message from Peer: %s", r.Status)
	return nil
}
