package kademlia

import (
	"container/list"
	"context"
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
	net "github/frisbee-cdn/frisbee-daemon/pkg/rpc"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindValue RPC Abstract Call
func (n *FrisbeeDHT) FindValue(ctx context.Context, reqBody *proto.FindValueRequest) (*proto.FindValueReply , error) {

	value, err := n.datastore.Get(reqBody.Key)
	logger.Info("FindValue Called by sender with ID: %s", reqBody.Id)
	if err != nil{
		logger.Errorf("Value lookup error: %s", err)
		return &proto.FindValueReply{}, nil
	}

	if value != nil {
		return &proto.FindValueReply{Content: value}, nil
	}

	logger.Info("No content found for key, returning closest peers")

	closestPeers := n.routingTable.FindClosestPeers(reqBody.Id, n.cfg.BucketSize)
	if len(closestPeers) > 0 {
		res := make( []*proto.Node, 0, n.cfg.BucketSize)
		for _, p := range closestPeers {
			res = append(res, &proto.Node{
				Id:   p.Node.ID,
				Addr: p.Node.GetHostAddress(),
				Port: p.Node.GetAddressPort(),
			})
		}
		logger.Infof("Succesfully found %d closest peers from the senders ID", len(closestPeers))
		return &proto.FindValueReply{Nodes: res}, nil
	}
	logger.Error("No closest peers found")
	return &proto.FindValueReply{}, nil

}

// FindValueRequest
func (n *FrisbeeDHT) FindValueRequest(ctx context.Context, key string, addr string,
	value chan []byte, contacts chan kbucket.Contacts){

	client, err := net.Connect(addr)

	if err != nil{
		value <- nil
		contacts <- nil
		return
	}

	defer client.Close()

	logger.Infof("Trying to find value from Node with address = %s", addr)

	r, err := client.GetClient().FindValue(ctx, &proto.FindValueRequest{Key: key, Id: n.node.ID})

	if r != nil{
		if r.Content != nil{

			logger.Info("Succesfully found key from contact")
			value <- r.Content
			contacts <- nil
			return
		}

		if r.Nodes !=nil{
			cts := make(kbucket.Contacts, 0, len(r.Nodes))
			for _, n := range r.Nodes {
				cts = append(cts, &kbucket.Contact{
					Node: common.NewNode(n.Addr, n.Port, n.Id),
				})
			}
			logger.Infof("Succesfully found %d closest peers from the senders ID", len(cts))
			contacts <- cts
			value <- nil
			return
		}
	}
	contacts <- nil
	value <- nil
}

// FindValueProxy
func (n *FrisbeeDHT)FindValueProxy(ctx context.Context, reqBody *proto.FindValueProxyRequest)(*proto.FindValueProxyReply, error){

	k := reqBody.Key
	value, err := n.datastore.Get(k)
	if err != nil{

		logger.Errorf("No value found for key: %s, trying to find on other peers", k)
		valueChan := make(chan []byte)
		contactsChan := make(chan kbucket.Contacts)


		queue := list.New()
		seen := make(map[string]bool)

		delta := int(n.cfg.ParallelismDegree)
		for _, c := range n.routingTable.FindClosestPeers(n.node.ID, delta) {
			queue.PushBack(c)
			seen[c.Node.ID.String()] = true
		}

		logger.Infof("Found %d contacts that may containt the value", len(seen))
		pending := 0

		for i := 0; i < delta && queue.Len() > 0 ; i++ {
			pending++
			c := queue.Front()
			go n.FindValueRequest(context.Background(), k,
				fmt.Sprintf("%s:%d", c.Value.(*kbucket.Contact).Node.GetHostAddress(),
					c.Value.(*kbucket.Contact).Node.GetAddressPort()),
				valueChan, contactsChan)
			queue.Remove(c)
		}

		for pending > 0 {
			pending--

			select {

			case val := <-valueChan:
				return &proto.FindValueProxyReply{Value: val}, nil
			case contacts := <- contactsChan:
				for _, c := range contacts{
					if _, ok := seen[c.Node.ID.String()]; !ok {
						queue.PushBack(c)
						seen[c.Node.ID.String()] = true
					}
				}
				for pending < delta && queue.Len() > 0{
					pending++
					c := queue.Front()
					go n.FindValueRequest(context.Background(), k,
						fmt.Sprintf("%s:%d", c.Value.(*kbucket.Contact).Node.GetHostAddress(),
							c.Value.(*kbucket.Contact).Node.GetAddressPort()),
						valueChan, contactsChan)
					queue.Remove(c)
				}
			default:
				return &proto.FindValueProxyReply{}, nil
			}
		}
	}

	if value != nil{
		return &proto.FindValueProxyReply{Value: value}, nil
	}

	return nil, nil
}

