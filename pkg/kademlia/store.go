package kademlia

import (
	"fmt"
	"container/list"
	"context"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
	net "github/frisbee-cdn/frisbee-daemon/pkg/rpc"
)

// Store
func (n *FrisbeeDHT) Store(ctx context.Context, reqBody *proto.StoreRequest) (*proto.Error, error) {
	err := n.datastore.Put(string(reqBody.Key), reqBody.Content)
	if err != nil {
		logger.Errorf("Error adding in datastore: %s", err)
	}

	logger.Info("Successfully stored content on local datastore")
	return &proto.Error{}, nil
}

// Store
func (n *FrisbeeDHT) StoreRequest(ctx context.Context, key string, value []byte, addr string, done chan bool){

	client,err := net.Connect(addr)
	if err != nil{
		done <- false
		return
	}
	defer client.Close()

	logger.Infof("Store to Key: %s Node with address = %s", key, addr)
	_, err = client.GetClient().Store(ctx, &proto.StoreRequest{Key: key, Content: value})
	if err != nil{
		done <- false
		return
	}
	logger.Info("Successfully stored value on contact")
	done <- true


}

func (n *FrisbeeDHT) iterativeStore() {

}


func (n *FrisbeeDHT) StoreProxy(ctx context.Context, reqBody *proto.StoreProxyRequest)(*proto.Error, error){

	k := reqBody.Key
	val := reqBody.Content

	err := n.datastore.Put(k, val)
	if err != nil{
		logger.Errorf("Error adding in datastore: %s", err)
	}

	done := make(chan bool)
	queue := list.New()
	seen := make(map[string]bool)

	delta := int(n.cfg.ParallelismDegree)
	for _, c := range n.routingTable.FindClosestPeers(n.node.ID, delta) {
		queue.PushBack(c)
		seen[c.Node.ID.String()] = true
	}

	logger.Info("Mass store on nearest contacts")
	pending := 0
	for i := 0; i < delta && queue.Len() > 0 ; i++ {
		pending++
		c := queue.Front()
		go n.StoreRequest(context.Background(), k, val, fmt.Sprintf("%s:%d", c.Value.(*kbucket.Contact).Node.GetHostAddress(),
			c.Value.(*kbucket.Contact).Node.GetAddressPort()), done)
		queue.Remove(c)
	}

	for pending > 0{
		pending--

		stored := <-done

		if stored {
			logger.Infof("Successfully stored content on contact")
			return &proto.Error{}, nil
		}else{
			return &proto.Error{Message: fmt.Sprintf("Couldn't store content for key %s", err)}, nil
		}
	}

	return nil, nil
}