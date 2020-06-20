package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Store
func (n *FrisbeeDHT) Store(ctx context.Context, reqBody *proto.StoreRequest) (*proto.Error, error) {
	err := n.datastore.Put(string(reqBody.Key), reqBody.Content)
	if err != nil {
		logger.Errorf("Error adding in datastore: %s", err)
	}
	return nil, nil
}

// Store
func (n *FrisbeeDHT) StoreRequest(ctx context.Context, key string, value []byte, addr string) error{

	client,err := n.service.Connect(addr)
	if err != nil{
		return err
	}
	defer client.Close()

	logger.Infof("Store to Key: %s Node with address = %s", key, addr)
	_, err = client.GetClient().Store(ctx, &proto.StoreRequest{Key: key, Content: value})
	if err != nil{
		return err
	}
	logger.Info("Successfully stored value in network")
	return nil


}

func (n *FrisbeeDHT) iterativeStore() {

}
