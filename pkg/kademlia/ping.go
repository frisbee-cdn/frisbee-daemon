package kademlia

import (
	"context"
	"fmt"

	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Ping RPC Abstract Call
func (n *FrisbeeDHT) Ping(ctx context.Context, reqBody *proto.PingRequest) (*proto.PingReply, error) {

	p := reqBody.Origin.Id

	logger.Infof("Ping from Sender: %x with Message: %s", p, reqBody.GetMessage())

	_, err := n.routingTable.Add(common.NewNode(reqBody.Origin.Addr,
		reqBody.Origin.Port,
		reqBody.Origin.Id), true, false)
	if err != nil {
		return &proto.PingReply{Status: "Error", Error: &proto.Error{
			Message: fmt.Sprintf("%s", err),
		}}, nil
	}

	n.routingTable.PrintInfo()

	recipient := &proto.Node{
		Id:   n.node.ID,
		Addr: n.node.GetHostAddress(),
		Port: n.node.GetAddressPort(),
	}
	return &proto.PingReply{Status: "Pong", Recipient: recipient}, nil
}

// Ping
func (n *FrisbeeDHT) PingRequest(ctx context.Context, addr string) error {

	client, err := n.service.Connect(addr)
	//print(client)
	if err != nil {
		return err
	}

	req := &proto.PingRequest{
		Message: "PING",
		Origin: &proto.Node{
			Id:   n.node.ID,
			Port: n.node.GetAddressPort(),
			Addr: n.node.GetHostAddress(),
		},
	}

	logger.Infof("Trying to Ping Node with address = %s", addr)
	r, err := client.GetClient().Ping(ctx, req)
	if err != nil {
		return err
	}
	defer client.Close()

	if r.Error != nil {
		logger.Errorf("Error Pinging: %s", r.Error.Message)
	} else {
		_, err = n.routingTable.Add(common.NewNode(r.Recipient.Addr,
			r.Recipient.Port,
			r.Recipient.Id), true, false)
		if err != nil {
			logger.Errorf("Error Adding Contact: %s", err)
		}
	}
	logger.Infof("Received message from Peer: %s", r.Status)
	return nil
}
