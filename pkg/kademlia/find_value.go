package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindValue RPC Abstract Call
func (n *FrisbeeNode) FindValue(ctx context.Context, reqBody *proto.ID) (*proto.StorageResponse, error) {
	return nil, nil
}

// FindValueRequest
func (n *FrisbeeNode) FindValueRequest(ctx context.Context, addr string) {

}
