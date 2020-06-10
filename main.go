package main


import (

	config "github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia"
)
func main() {

	cfg := config.InitConfiguration("development")

	node := kademlia.BootStrap(cfg)
	println(node.Id)
	println(node.Addr)
	println(node.Port)

	//id1, _ := peer.HashKey("192")
	//id2, _ := peer.HashKey("440")
	//
	//fmt.Printf("Haskeys: \n {%v} \n {%v}", id1, id2)
	//
	//dist := peer.XOR(id1, id2)
	//fmt.Printf("\nDistance {%v}", dist)
	//
	//leadingZeros := peer.ZeroPrefixLen(dist)
	//
	//fmt.Printf("Peer: {%v}", leadingZeros)

}
