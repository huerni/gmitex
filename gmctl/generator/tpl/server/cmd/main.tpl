package main

import (
	"context"
    {{.imports}}
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
