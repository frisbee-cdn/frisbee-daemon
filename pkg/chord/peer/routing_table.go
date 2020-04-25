package peer

import "github/frisbee-cdn/frisbee-daemon/pkg/util"

// RoutingTable used to store a subset of peer's of the network.
type RoutingTable struct {
	KBuckets [util.IDLENGTH * 8]*KBucket
}

// NewRoutingTable is used to create an new empty routing table.
func NewRoutingTable() *RoutingTable {

	rt := &RoutingTable{
		KBuckets:[util.IDLENGTH * 8]*KBucket{},
	}

	for i := 0; i<util.IDLENGTH * 8; i++{
		rt.KBuckets[i] = NewBucket()
	}

	return rt
}

// Add is used to add a new Node inside the Routing Table
func (r *RoutingTable)Add(node Node){

}

// FindClosest used to find the closes node in the network
func (r *RoutingTable) FindClosest() {

}

