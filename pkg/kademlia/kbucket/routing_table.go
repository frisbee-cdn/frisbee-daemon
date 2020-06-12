package peer

import (
	model "github/frisbee-cdn/frisbee-daemon/pkg/rpc/proto"
	"time"
)

// RoutingTable used to store a subset of kbucket's of the network.
type RoutingTable struct {
	self ID

	KBuckets   []*KBucket
	bucketSize int

	// Maximum acceptable latency for peers in this table
	maxLatency time.Duration
}

// NewRoutingTable is used to create an new empty routing table.
func NewRoutingTable(bucketsize int, local ID, latency time.Duration) (*RoutingTable, error) {

	rt := &RoutingTable{
		KBuckets:   []*KBucket{NewBucket()},
		bucketSize: bucketsize,

		self: local,

		maxLatency: latency,
	}

	return rt, nil
}

// Update is used to add a new FrisbeeNode inside the Routing Table
func (rt *RoutingTable) Add(node *model.Node, queryPeer bool, isReplaceable bool) (bool, error) {

	bucketId := rt.bucketIdForPeer(node.Id)
	bucket := rt.KBuckets[bucketId]

	now := time.Now()
	var lastUsefulAt time.Time
	if queryPeer {
		lastUsefulAt = now
	}

	// peer already exists in the Routing TAble
	if peer := bucket.find(node.Id); peer != nil {

		if peer.LastUsefulAt.IsZero() && queryPeer {
			peer.LastUsefulAt = lastUsefulAt
		}

		bucket.MoveToBack(node.Id)
		return true, nil
	} else {

		if bucket.Len() < rt.bucketSize {
			bucket.PushBack(&Contact{
				Id:			  node.Id
				LastUsefulAt: lastUsefulAt,
				AddedAt:      now,
				replaceable:  isReplaceable,
			})
			return true, nil
		} else {

		}
	}
	return true, nil
}

// Remove is used to delete a FrisbeeNode inside the Routing Table
func (rt *RoutingTable) Remove(id ID) bool {

	bucketId := rt.bucketIdForPeer(id)
	bucket := rt.KBuckets[bucketId]

	if bucket.remove(id) {

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
func (rt *RoutingTable) FindClosestPeer(targetId ID) Contact {
	return Contact{}
}

// FindClosestPeers
func (rt *RoutingTable) FindClosestPeers(targetId ID,  count int) []ID{

	cpl := CommonPrefixLen(targetId, rt.self)

	if cpl >= len(rt.KBuckets)-1 {
		cpl = len(rt.KBuckets) - 1
	}

	pds := peerDistanceSorter{
		peers: make([]peerDistance, 0, count+rt.bucketSize),
		target: targetId,
	}

	pds.appendPeersFromList(rt.KBuckets[cpl].List)

	if pds.Len() < count {
		for i := cpl + 1; i < len(rt.KBuckets); i++{
			pds.appendPeersFromList(rt.KBuckets[i].List)
		}
	}

	for i := cpl - 1; i >= 0 && pds.Len() < count; i-- {
		pds.appendPeersFromList(rt.buckets[i].list)
	}

	pds.sort()

	if count < pds.Len() {
		pds.peers = prds.peers[:count]
	}

	out := make([]ID, 0, pds.Len())
	for _, p := range pds.peers{
		out = append(out, p.p)
	}

	return out
}

// PrinInfo prints a description about this RoutingTable
func (rt *RoutingTable) PrintInfo() {

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
	newBucket := bucket.Split(len(rt.KBuckets), rt.self)
	rt.KBuckets = append(rt.KBuckets, newBucket)

	if newBucket.Len() >= rt.bucketSize {
		rt.nextBucket()
	}

}

func (rt *RoutingTable) bucketIdForPeer(target ID) int {
	cpl := CommonPrefixLen(target, rt.self)

	bucketId := cpl
	if bucketId >= len(rt.KBuckets) {
		bucketId = len(rt.KBuckets) - 1
	}

	return bucketId
}
