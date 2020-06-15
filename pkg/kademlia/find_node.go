package kademlia

import (
	"context"
	kb "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindNode RPC Abstract Call
func (n *FrisbeeNode) FindNode(ctx context.Context, reqBody *proto.ID) (*proto.NodeResponse, error) {

	senderID := reqBody.Id
	logger.Info("FindNode Called by sender with ID: %x", senderID)
	res := make([]*proto.Node, 0, n.cfg.BucketSize)
	closestPeers := n.routingTable.FindClosestPeers(senderID, n.cfg.BucketSize)
	for _, p := range closestPeers {
		res = append(res, p.Node)
	}
	logger.Infof("Succesfully found %d closest peers from the senders ID", len(closestPeers))
	return &proto.NodeResponse{Nodes: res}, nil
}

// FindNodeRequest
func (n *FrisbeeNode) FindNodeRequest(ctx context.Context, target kb.ID, addr string) ([]*proto.Node, error) {
	client, err := n.service.Connect(addr)

	if err != nil {
		return nil, err
	}

	r, err := client.FindNode(ctx, &proto.ID{Id: target})
	if err != nil {
		return nil, err
	}

	return r.Nodes, nil

}

// NodeLookup
func (n *FrisbeeNode) NodeLookup(ctx context.Context, target kb.ID, addr string, done chan []*proto.Node) {

	client, err := n.service.Connect(addr)
	if err != nil {
		done <- nil
		return
	}

	r, err := client.FindNode(ctx, &proto.ID{Id: target})
	if err != nil {
		done <- nil
		return
	}

	//TODO: Add Sender ID to routingTable
	//n.routingTable.ADD(true, false)

	done <- r.Nodes
}
