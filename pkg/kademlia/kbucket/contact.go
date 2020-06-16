package kbucket

import (
	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	"time"
)

// Contact holds all related information for a peer in the K-Bucket
type Contact struct {
	Node *peer.Node

	LastUsefulAt time.Time

	// Added At is the time this peer was added to the routing table
	AddedAt time.Time

	// if a bucket is full, this peer can be replaced to make space for a new peer.
	replaceable bool
}
