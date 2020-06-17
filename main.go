package main

import (
	"flag"
	"github/frisbee-cdn/frisbee-daemon/pkg/kademlia"

	//"github/frisbee-cdn/frisbee-daemon/pkg/kademlia"
)

func main() {

	// cfg := config.InitConfiguration("development")

	//node, _ := kademlia.New("SanFrancisco-FirstNode", 5001, nil)
	//print(node)
	//

	host := flag.String("host","name", "a string")
	port := flag.Int("port", 8888, "an int")
	isBootstrap := flag.Bool("boot", false, "a bool")

	flag.Parse()

	kademlia.Bootstrap(*host, uint32(*port), *isBootstrap)
	done := make(chan bool)
	_ = <- done
}
