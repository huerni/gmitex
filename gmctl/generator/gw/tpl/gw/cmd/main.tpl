package main

import (
	"context"
	"gateway/internal/app"
	"gateway/internal/config"
)

func main() {
	discovery.EnableDiscovery()
    server := core.NewAPIGatewayServer()
    server.Start()
}
