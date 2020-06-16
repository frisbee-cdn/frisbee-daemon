package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Ping RPC Abstract Call
func (n *FrisbeeDHT) Ping(ctx context.Context, reqBody *proto.CheckStatusRequest) (*proto.CheckStatusReply, error) {
	logger.Info("Ping: ", reqBody.GetMessage())

	//TODO: Try To Add Sender to routingTable
	return &proto.CheckStatusReply{Status: "Pong"}, nil
}

// Ping
func (n *FrisbeeDHT) PingRequest(ctx context.Context, message string, addr string) error {

	client, err := n.service.Connect(addr)
	if err != nil {
		return err
	}

	r, err := client.Ping(ctx, &proto.CheckStatusRequest{Message: message})
	if err != nil {
		return err
	} else {
		//TODO: Add Probbed Node to routingTable
	}
	logger.Infof("Received message from Peer: %s", r.Status)
	return nil
}
