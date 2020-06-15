package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Store
func (n *FrisbeeNode) Store(ctx context.Context, reqBody *proto.StoreRequest) (*proto.Error, error) {

	return nil, nil
}

// Store
func (n *FrisbeeNode) StoreRequest(ctx context.Context, addr string) {

}
