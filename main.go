package main

import (
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia"
)

func main() {

	// cfg := config.InitConfiguration("development")

	node, _ := kademlia.New("SanFrancisco", 5001, nil)
	print(node)

	// time.Sleep(2 * time.Second)
	// node1, _ := kademlia.New("Romania", 5002, nil)

	// node1.Join(node.Addr, node.Port)

	// id1, _ := peer.HashKey("192")
	// id2, _ := peer.HashKey("440")

	// fmt.Printf("Haskeys: \n {%v} \n {%v}", id1, id2)

	// dist := peer.XOR(id1, id2)
	// fmt.Printf("\nDistance {%v}", dist)

	// leadingZeros := peer.ZeroPrefixLen(dist)

	// fmt.Printf("Peer: {%v}", leadingZeros)

}
