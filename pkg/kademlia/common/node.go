package common

// Node is the generic representation of the node
// will e added to the Peer Node structure
type Node struct {
	NetworkAddress AddressInfo
	ID             NodeID
}

// NewNode creates an new Node
func NewNode(hostAddress string, hostPort uint32, id NodeID) *Node {

	return &Node{
		NetworkAddress: AddressInfo{
			Host: hostAddress,
			Port: hostPort,
		},
		ID: id,
	}
}

// GetHostAddress
func (n *Node) GetHostAddress() string {
	return n.NetworkAddress.Host
}

// GetAddressPort
func (n *Node) GetAddressPort() uint32 {
	return n.NetworkAddress.Port
}
