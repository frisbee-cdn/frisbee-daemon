package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Store
func (n *FrisbeeDHT) Store(ctx context.Context, reqBody *proto.StoreRequest) (*proto.Error, error) {
	err := n.datastore.Put(string(reqBody.Key), reqBody.Content)
	if err != nil {
		return
	}
}

// Store
func (n *FrisbeeDHT) StoreRequest(ctx context.Context, key string, addr string) {

}

func (n *FrisbeeDHT) iterativeStore() {

}
