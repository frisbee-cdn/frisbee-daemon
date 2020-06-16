package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Store
func (n *FrisbeeDHT) Store(ctx context.Context, reqBody *proto.StoreRequest) (*proto.Error, error) {

	return nil, nil
}

// Store
func (n *FrisbeeDHT) StoreRequest(ctx context.Context, addr string) {

}

func (n *FrisbeeDHT) iterativeStore() {

}
