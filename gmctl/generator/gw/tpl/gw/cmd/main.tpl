package main

import (
	"gateway/internal/config"
	"github.com/huerni/gmitex/core/discovery"
	"github.com/huerni/gmitex/core/gateway"
)

func main() {
	discovery.EnableDiscovery(config.GetRegisterCenter())
	server := core.NewAPIGatewayServer(config.GetServerConfig())
	server.Start(config.GetGatewayRouter())
}
