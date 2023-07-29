package main

import (
	"context"
	"gateway/internal/app"
	"gateway/internal/config"
)

func main() {
	c, err := config.InitConfig("etc/cfg.toml")
	if err != nil {
		panic(err)
	}

	g := app.NewGmServer(c)
	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
