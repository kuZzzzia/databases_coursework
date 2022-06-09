package main

import (
	"./src/config"
	"./src/server"
)

func main() {
	server.Start(config.GenerateConfig())
}
