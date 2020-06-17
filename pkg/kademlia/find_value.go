package kademlia

import (
	"context"
	proto "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// FindValue RPC Abstract Call
func (n *FrisbeeDHT) FindValue(ctx context.Context, reqBody *proto.FindValueRequest) (*proto.FindValueReply, error) {
	return nil, nil
}

// FindValueRequest
func (n *FrisbeeDHT) FindValueRequest(ctx context.Context, addr string) {

}

func (n *FrisbeeDHT) iterativeFindValue() {

}
