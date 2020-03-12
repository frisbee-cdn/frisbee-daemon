package peer

// KBucket is a list of routing addresses of other nodes in
// the network.
type KBucket struct {
	IP     string
	Port   string
	NodeID string
}

// NewBucket is used to create an empty KBucket
func NewBucket() *KBucket {
	return nil
}
