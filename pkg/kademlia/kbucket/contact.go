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
	Replaceable bool
}

// NewContact creates a new Contact
func NewContact(node *peer.Node, lastUsefulAt time.Time, addedAt time.Time, replaceable bool) *Contact {
	return &Contact{
		Node:         node,
		LastUsefulAt: lastUsefulAt,
		AddedAt:      addedAt,
		Replaceable:  replaceable,
	}
}

// Contacts a list of contacts
type Contacts []*Contact

//func (c Contacts) Len() int { return len(c) }
//func (c Contacts) Less(i, j int) bool { return c[i].Node.ID.Less(c[j].Node.ID)}
//func (c Contacts) Swap(i, j int) { c[i], c[j] = c[j], c[i]}
//
//func (c *Contacts) Push(x interface{}) {
//	*c = append(*c, x.(*Contact))
//}
//
//func (c *Contacts) Pop() interface{} {
//	oldHeap := *c
//	oldLen := len(oldHeap)
//	elem := oldHeap[oldLen - 1]
//	*c = oldHeap[0 : oldLen - 1]
//	return elem
//}


