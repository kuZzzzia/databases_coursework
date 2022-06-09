package main

import (
	"film-network/src/config"
	"film-network/src/server"
)

func main() {
	server.Start(config.GenerateConfig())
}
