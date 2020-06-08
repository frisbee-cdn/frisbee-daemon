package main

import (
	"fmt"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia/kbucket"
)

func main() {

	//config := cfg.InitConfiguration("development")

	//node := kbucket.BootStrap(config)
	//println(node.Peer.Id)
	//println(node.Peer.Addr)
	//println(node.Peer.Port)

	id1, _ := peer.HashKey("192")
	id2, _ := peer.HashKey("192")

	fmt.Printf("Haskeys: \n {%v} \n {%v}", id1, id2)

	dist := peer.XOR(id1, id2)
	fmt.Printf("\nDistance {%v}", dist)

	leadingZeros := peer.ZeroPrefixLen([20]byte{1})

	fmt.Printf("Peer: {%v}", leadingZeros)


}
