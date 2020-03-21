package peer

import "github/frisbee-cdn/frisbee-daemon/pkg/util"

// RoutingTable used to store a subset of peer's of the network.
type RoutingTable struct {
	KBuckets [util.IDBYTESLENGTH]*KBucket
}

// NewRoutingTable is used to create an new empty routing table.
func NewRoutingTable() *RoutingTable {
	return nil
}

// FindClosest used to find the closes node in the network
func (r *RoutingTable) FindClosest() {

}
