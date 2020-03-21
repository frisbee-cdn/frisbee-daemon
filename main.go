package main

import (
	cfg "github/frisbee-cdn/frisbee-daemon/internal"
	"github/frisbee-cdn/frisbee-daemon/pkg/chord/peer"
)

func main() {

	config := cfg.InitConfiguration("development")
	println(config.Server.Port)

	node := peer.BootStrap(config)
	println(node.Node.Id)
	println(node.Node.Addr)

}
