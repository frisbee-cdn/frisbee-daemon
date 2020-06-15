package kbucket

import (
	"container/list"
	"time"

	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
)

// Contact holds all related information for a peer in the K-Bucket
type Contact struct {
	*model.Node

	// Id of the peer
	ownerId peer.ID

	LastUsefulAt time.Time

	// Added At is the time this peer was added to the routing table
	AddedAt time.Time

	// if a bucket is full, this peer can be replaced to make space for a new peer.
	replaceable bool
}

// KBucket is a list of routing addresses of other nodes in
// the network.
type KBucket struct {
	*list.List
}

// NewBucket is used to create an empty KBucket
func NewBucket() *KBucket {
	return &KBucket{
		list.New(),
	}
}

func (kb *KBucket) GetAllPeers() []Contact {

	var peers []Contact

	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		p := elem.Value.(*Contact)
		peers = append(peers, *p)
	}
	return peers
}

// find returns the kbucket with the given Id if it exists
// returns nil if the peerId does not exist in the kBucket
func (kb *KBucket) find(p peer.ID) *Contact {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).ownerId == p {
			return elem.Value.(*Contact)
		}
	}
	return nil
}

// remove deletes the kbucket with the given Id it exists and returns true
// returns false if the kbucket does not exist in the kBucket
func (kb *KBucket) remove(p peer.ID) bool {

	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).ownerId == p {
			kb.List.Remove(elem)
			return true
		}
	}
	return false
}

// Len returns the size of the bucket
func (kb *KBucket) Len() int {
	return kb.List.Len()
}

// MoveToFront
func (kb *KBucket) MoveToFront(p peer.ID) {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).ownerId == p {
			kb.List.MoveToFront(elem)
		}
	}
}

// MoveToBack
func (kb *KBucket) MoveToBack(p peer.ID) {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).ownerId == p {
			kb.List.MoveToBack(elem)
		}
	}
}

// PushFront
func (kb *KBucket) PushFront(p *Contact) {
	kb.List.PushFront(p)
}

// PushBack
func (kb *KBucket) PushBack(p *Contact) {
	kb.List.PushBack(p)
}

// Split splits a bucket peers into two buckets
// TODO: Further examination to understand how it works for Kademlia
func (kb *KBucket) Split(cpl int, target ID) *KBucket {

	newBucket := NewBucket()

	elem := kb.List.Front()

	for elem != nil {

		pId := elem.Value.(*Contact).Id
		peerCpl := CommonPrefixLen(pId, target)
		if peerCpl > cpl {

			cur := elem
			newBucket.List.PushBack(elem.Value)
			elem = elem.Next()
			kb.List.Remove(cur)
			continue
		}

		elem = elem.Next()
	}

	return newBucket
}

func (kb *KBucket) maxCommonPrefix(target ID) uint {

	maxCpl := uint(0)
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		cpl := uint(CommonPrefixLen(elem.Value.(*Contact).Id, target))
		if cpl > maxCpl {
			maxCpl = cpl
		}
	}
	return maxCpl
}
