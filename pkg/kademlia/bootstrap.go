package kademlia

// Bootstrap is used to bootstrap a new node into the network
func Bootstrap(host string, port uint32, isBootstrap bool) {

	dht, _ := New(host, port, nil)
	if !isBootstrap {
		dht.Join()
	}
}
