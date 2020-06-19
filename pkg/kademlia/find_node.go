package kademlia

import (
	"context"
	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindNode RPC Abstract Call
func (n *FrisbeeDHT) FindNode(ctx context.Context, reqBody *proto.FindNodeRequest) (*proto.FindNodeReply, error) {

	senderID := reqBody.Id
	logger.Infof("FindNode Called by sender with ID: %x", senderID)
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

// FindNodeRequest Client call operation
func (n *FrisbeeDHT) FindNodeRequest(ctx context.Context, target peer.NodeID, addr string) ([]*proto.Node, error) {
	client, err := n.service.Connect(addr)

	if err != nil {
		return nil, err
	}
	defer client.Close()

	logger.Infof("Trying to collect closest Peers from Node with address = %s ", addr)
	r, err := client.GetClient().FindNode(ctx, &proto.FindNodeRequest{Id: target})
	if err != nil {
		return nil, err
	}

	logger.Infof("Succesfully received %d clossest peers", len(r.Nodes))
	return r.Nodes, nil

}
func (n *FrisbeeDHT) iterativeFindNode() {

}
