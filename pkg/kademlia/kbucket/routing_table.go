package peer

import (
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia"
	"time"
)



// RoutingTable used to store a subset of kbucket's of the network.
type RoutingTable struct {

	self ID


	KBuckets []*KBucket
	bucketSize int


	// Maximum acceptable latency for peers in this table
	maxLatency time.Duration
}

// NewRoutingTable is used to create an new empty routing table.
func NewRoutingTable(bucketsize int, local ID, latency time.Duration) (*RoutingTable, error) {

	rt := &RoutingTable{
		KBuckets:[]*KBucket{NewBucket()},
		bucketSize: bucketsize,

		self: local,

		maxLatency: latency,


	}

	return rt, nil
}

// Update is used to add a new Peer inside the Routing Table
func (rt *RoutingTable)Update(id ID, queryPeer bool, isReplaceable bool) (bool, error){


	bucketId := rt.bucketIdForPeer(id)

	bucket := rt.KBuckets[bucketId]

	now := time.Now()
	var lastUsefulAt time.Time
	if queryPeer {
		lastUsefulAt = now
	}

	if peer := bucket.find(id); peer != nil{

		if peer.LastUsefulAt.IsZero() && queryPeer{
			peer.LastUsefulAt = lastUsefulAt
		}
		return false, nil
	}
	bucket.Update(id)

	// TODO: Continue Algorithm To implement
}

// Remove is used to delete a Peer inside the Routing Table
func (rt *RoutingTable)Remove(node kademlia.Peer){

}

// FindClosestPeer used to find the closes node in the network
func (rt *RoutingTable) FindClosestPeer(targetId []byte) kademlia.Peer {
}

// FindClosestPeers
func (rt *RoutingTable) FindClosestPeers(targetId []byte) []kademlia.Peer {
	return nil
}

// PrinInfo prints a description about this RoutingTable
func (rt *RoutingTable) PrintInfo(){

}

// Size returns the total number of peers in the routing table
func (rt *RoutingTable) Size() int {

	var totalPeers int

	for _, buck := range rt.KBuckets {
		totalPeers += buck.Len()
	}
	return totalPeers


}


func (rt *RoutingTable) bucketIdForPeer(target ID) int{
	cpl := CommonPrefixLen(target, rt.self)

	bucketId := cpl
	if bucketId >= len(rt.KBuckets){
		bucketId = len(rt.KBuckets) - 1
	}

	return bucketId
}