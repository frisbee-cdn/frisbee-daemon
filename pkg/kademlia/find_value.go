package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindValue RPC Abstract Call
func (n *FrisbeeDHT) FindValue(ctx context.Context, reqBody *proto.FindValueRequest) (*proto.FindValueReply , error) {

	value, err := n.datastore.Get(reqBody.Key)
	logger.Info("FindValue Called by sender with ID: %x", reqBody.Id)
	if err != nil{
		logger.Errorf("Value lookup error: %s", err)
	}

	if value != nil {
		return &proto.FindValueReply{Content: value}, nil
	}

	logger.Info("No content found for key, returning closest peers")
	res := make( []*proto.Node, 0, n.cfg.BucketSize)
	closestPeers := n.routingTable.FindClosestPeers(reqBody.Id, n.cfg.BucketSize)
	for _, p := range closestPeers{
		res = append(res, &proto.Node{
			Id: p.Node.ID,
			Addr: p.Node.GetHostAddress(),
			Port: p.Node.GetAddressPort(),
		})
	}
	logger.Infof("Succesfully found %d closest peers from the senders ID", len(closestPeers))
	return &proto.FindValueReply{Nodes: res}, nil
}

// FindValueRequest
func (n *FrisbeeDHT) FindValueRequest(ctx context.Context, key string, addr string) ([]*proto.Node, []byte, error){

	client, err := n.service.Connect(addr)

	if err != nil{
		return nil, nil, err
	}

	defer client.Close()

	logger.Infof("Trying to find value from Node with address = %s", addr)

	r, err := client.GetClient().FindValue(ctx, &proto.FindValueRequest{Key: key, Id: n.node.ID})

	if r.Content != nil{
		return nil, r.Content, nil
	}
	return r.Nodes, nil, nil

}

func (n *FrisbeeDHT) iterativeFindValue() {

}
