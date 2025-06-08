package main

import (
	"github.com/manattan/mumbai/internal/config"
	"github.com/manattan/mumbai/internal/gateway/server"
)

func main() {
	cfg := config.Load()
	server.StartEchoServer(cfg)
}