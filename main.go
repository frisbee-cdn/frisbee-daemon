package main

import (
	"fmt"
	cfg "github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/chord/peer"
	"github/frisbee-cdn/frisbee-daemon/pkg/util"
)

func main() {

	config := cfg.InitConfiguration("development")

	node := peer.BootStrap(config)
	println(node.Node.Id)
	println(node.Node.Addr)
	println(node.Node.Port)


	node1 := peer.BootStrap(config)
	println(node1.Node.Id)

	distance := util.Xor(node.Id, node1.Id)
	fmt.Printf("Distance: %v", distance)


}
