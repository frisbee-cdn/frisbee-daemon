package kademlia

import (
	"context"
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Ping RPC Abstract Call
func (n *FrisbeeDHT) Ping(ctx context.Context, reqBody *proto.PingRequest) (*proto.PingReply, error) {
	logger.Info("Ping: ", reqBody.GetMessage())

	_, err := n.routingTable.Add(common.NewNode(reqBody.Origin.Addr, reqBody.Origin.Port, reqBody.Origin.Id), true, false)
	if err != nil {
		return &proto.PingReply{Status: "Error", Error: &proto.Error{
			Message: fmt.Sprintf("%s", err),
		}}, nil
	}
	return &proto.PingReply{Status: "Pong"}, nil
}

// Ping
func (n *FrisbeeDHT) PingRequest(ctx context.Context, addr string) error {

	client, err := n.service.Connect(addr)
	if err != nil {
		return err
	}

	req := &proto.PingRequest{
		Message: "PING",
		Origin: &proto.Node{
			Id:   n.node.SelfKey,
			Port: n.node.GetAddressPort(),
			Addr: n.node.GetHostAddress(),
		},
	}

	r, err := client.Ping(ctx, req)
	if err != nil {
		return err
	} else {
		//n.routingTable.Add()
	}
	logger.Infof("Received message from Peer: %s", r.Status)
	return nil
}
