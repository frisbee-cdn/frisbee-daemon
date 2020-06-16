package common

// Node is the generic representation of the node
// will e added to the Peer Node structure
type Node struct {
	HostAddress AddressInfo
	SelfID      ID
	SelfKey     []byte
}

// NewNode creates an new Node
func NewNode(hostAddress string, hostPort uint32, selfID ID) *Node {

	return &Node{
		HostAddress: AddressInfo{
			Host: hostAddress,
			Port: hostPort,
		},
		SelfID: selfID,
	}
}

// GetHostAddress
func (n *Node) GetHostAddress() string {
	return n.HostAddress.Host
}

// GetAddressPort
func (n *Node) GetAddressPort() uint32 {
	return n.HostAddress.Port
}