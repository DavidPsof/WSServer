package main

import (
	"WSServer/config"
	"WSServer/logger"
)

func main() {
	config.Init()
	logger.Init()
}
