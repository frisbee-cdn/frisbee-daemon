package main

import (
	cfg "github/frisbee-cdn/frisbee-daemon/internal"
)

func main() {

	config := cfg.InitConfiguration("development")
	println(config.Server.Port)

}
