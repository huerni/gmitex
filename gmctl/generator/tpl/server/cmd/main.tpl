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

	g := app.NewGmServer(c, {{.mserverName}}.Register{{.serverName}}HandlerFromEndpoint, func(server *grpc.Server) {
    		ctx := svc.NewServiceContext(c)
    		srv := handler.New{{.serverName}}Server(ctx)
    		{{.mserverName}}.Register{{.serverName}}Server(server, srv)
    })
	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
