package kbucket

import (
	"container/list"

	id "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
)

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
func (kb *KBucket) find(p id.NodeID) *Contact {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).Node.ID.Equals(p) {
			return elem.Value.(*Contact)
		}
	}
	return nil
}

// remove deletes the kbucket with the given Id it exists and returns true
// returns false if the kbucket does not exist in the kBucket
func (kb *KBucket) remove(p id.NodeID) bool {

	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).Node.ID.Equals(p) {
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
func (kb *KBucket) MoveToFront(p id.NodeID) {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).Node.ID.Equals(p) {
			kb.List.MoveToFront(elem)
		}
	}
}

// MoveToBack
func (kb *KBucket) MoveToBack(p id.NodeID) {
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(*Contact).Node.ID.Equals(p) {
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
func (kb *KBucket) Split(cpl int, target id.NodeID) *KBucket {

	newBucket := NewBucket()

	elem := kb.List.Front()

	for elem != nil {

		pId := elem.Value.(*Contact).Node.ID
		peerCpl := id.CommonPrefixLen(pId, target)
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

func (kb *KBucket) maxCommonPrefix(target id.NodeID) uint {

	maxCpl := uint(0)
	for elem := kb.List.Front(); elem != nil; elem = elem.Next() {
		cpl := uint(id.CommonPrefixLen(elem.Value.(*Contact).Node.ID, target))
		if cpl > maxCpl {
			maxCpl = cpl
		}
	}
	return maxCpl
}
