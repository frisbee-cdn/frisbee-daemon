package kbucket

import (
	peer "github/frisbee-cdn/frisbee-daemon/pkg/kademlia/common"
	"fmt"
	"time"
)

// RoutingTable used to store a subset of kbucket's of the network.
type RoutingTable struct {
	local peer.NodeID

	KBuckets   []*KBucket
	bucketSize int

	// Maximum acceptable latency for peers in this table
	maxLatency time.Duration
}

// NewRoutingTable is used to create an new empty routing table.
func NewRoutingTable(bucketsize int, localID peer.NodeID, latency time.Duration) (*RoutingTable, error) {

	rt := &RoutingTable{
		KBuckets:   []*KBucket{NewBucket()},
		bucketSize: bucketsize,
		local:      localID,

		maxLatency: latency,
	}

	return rt, nil
}

// Update is used to add a new FrisbeeNode inside the Routing Table
func (rt *RoutingTable) Add(node *peer.Node, queryPeer bool, isReplaceable bool) (bool, error) {

	bucketID := rt.bucketIdForPeer(node.ID)
	bucket := rt.KBuckets[bucketID]

	now := time.Now()
	var lastUsefulAt time.Time
	if queryPeer {
		lastUsefulAt = now
	}

	// peer already exists in the Routing TAble
	if c := bucket.find(node.ID); c != nil {

		if c.LastUsefulAt.IsZero() && queryPeer {
			c.LastUsefulAt = lastUsefulAt
		}

		bucket.MoveToBack(node.ID)
		return true, nil
	} else {
		contact := NewContact(node, lastUsefulAt, now, isReplaceable)
		if bucket.Len() < rt.bucketSize {
			bucket.PushBack(contact)
			return true, nil
		} else {

		}
	}
	return true, nil
}

// Remove is used to delete a FrisbeeNode inside the Routing Table
func (rt *RoutingTable) Remove(p peer.NodeID) bool {

	bucketID := rt.bucketIdForPeer(p)
	bucket := rt.KBuckets[bucketID]

	if bucket.remove(p) {

		for {
			lastBucketIndex := len(rt.KBuckets) - 1
			if len(rt.KBuckets) > 1 && rt.KBuckets[lastBucketIndex].Len() == 0 {
				rt.KBuckets[lastBucketIndex] = nil
				rt.KBuckets = rt.KBuckets[:lastBucketIndex]
			} else if len(rt.KBuckets) >= 2 && rt.KBuckets[lastBucketIndex-1].Len() == 0 {
				rt.KBuckets[lastBucketIndex-1] = rt.KBuckets[lastBucketIndex]
				rt.KBuckets[lastBucketIndex] = nil
				rt.KBuckets = rt.KBuckets[:lastBucketIndex]
			} else {
				break
			}
		}
		return true
	}
	return false
}

// FindClosestPeer used to find the closes node in the network
func (rt *RoutingTable) FindClosestPeer(targetID peer.NodeID) Contact {
	return Contact{}
}

// FindClosestPeers
func (rt *RoutingTable) FindClosestPeers(id peer.NodeID, count int) []*Contact {

	cpl := peer.CommonPrefixLen(id, rt.local)

	if cpl >= len(rt.KBuckets)-1 {
		cpl = len(rt.KBuckets) - 1
	}

	pds := peerDistanceSorter{
		peers:  make([]peerDistance, 0, count+rt.bucketSize),
		target: id,
	}

	pds.appendPeersFromList(rt.KBuckets[cpl].List)

	if pds.Len() < count {
		for i := cpl + 1; i < len(rt.KBuckets); i++ {
			pds.appendPeersFromList(rt.KBuckets[i].List)
		}
	}

	for i := cpl - 1; i >= 0 && pds.Len() < count; i-- {
		pds.appendPeersFromList(rt.KBuckets[i].List)
	}

	pds.sort()

	if count < pds.Len() {
		pds.peers = pds.peers[:count]
	}

	out := make([]*Contact, 0, pds.Len())
	for _, p := range pds.peers {
		out = append(out, p.c)
	}

	return out
}

// PrinInfo prints a description about this RoutingTable
func (rt *RoutingTable) PrintInfo() {

	fmt.Printf("Routing Table, bucket size = %d", rt.bucketSize)

	for i, b := range rt.KBuckets {
		fmt.Printf("\tBucket: %d \n", i)

		for elem := b.List.Front(); elem != nil; elem = elem.Next(){
			p := elem.Value.(*Contact).Node.ID
			fmt.Printf("\t\t - %x", p)
		}
	}
}

// Size returns the total number of peers in the routing table
func (rt *RoutingTable) Size() int {

	var totalPeers int

	for _, buck := range rt.KBuckets {
		totalPeers += buck.Len()
	}
	return totalPeers

}

func (rt *RoutingTable) nextBucket() {

	bucket := rt.KBuckets[len(rt.KBuckets)-1]
	newBucket := bucket.Split(len(rt.KBuckets), rt.local)
	rt.KBuckets = append(rt.KBuckets, newBucket)

	if newBucket.Len() >= rt.bucketSize {
		rt.nextBucket()
	}

}

func (rt *RoutingTable) bucketIdForPeer(p peer.NodeID) int {

	cpl := peer.CommonPrefixLen(p, rt.local)

	bucketId := cpl
	if bucketId >= len(rt.KBuckets) {
		bucketId = len(rt.KBuckets) - 1
	}

	return bucketId
}
