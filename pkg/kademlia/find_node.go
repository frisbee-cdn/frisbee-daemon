package kademlia

import (
	"context"
	kb "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindNode RPC Abstract Call
func (n *FrisbeeDHT) FindNode(ctx context.Context, reqBody *proto.FindNodeRequest) (*proto.FindNodeReply, error) {

	senderID := reqBody.Id
	logger.Info("FindNode Called by sender with ID: %x", senderID)
	res := make([]*proto.Node, 0, n.cfg.BucketSize)
	closestPeers := n.routingTable.FindClosestPeers(senderID, n.cfg.BucketSize)
	for _, p := range closestPeers {
		res = append(res, &proto.Node{
			Id:   p.Node.ID,
			Addr: p.Node.GetHostAddress(),
			Port: p.Node.GetAddressPort(),
		})
	}
	logger.Infof("Succesfully found %d closest peers from the senders ID", len(closestPeers))
	return &proto.FindNodeReply{Nodes: res}, nil
}

// FindNodeRequest
func (n *FrisbeeDHT) FindNodeRequest(ctx context.Context, target kb.ID, addr string) ([]*proto.Node, error) {
	client, err := n.service.Connect(addr)

	if err != nil {
		return nil, err
	}

	r, err := client.FindNode(ctx, &proto.FindNodeRequest{Id: target})
	if err != nil {
		return nil, err
	}

	return r.Nodes, nil

}
func (n *FrisbeeDHT) iterativeFindNode() {

}
