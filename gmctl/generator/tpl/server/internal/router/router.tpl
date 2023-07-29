package router

import (
	"context"
	"github.com/huerni/gmitex/pkg/gw"
)

type {{.serverName}}Router struct {
	Paths []string
}

func New{{.serverName}}Router() *{{.serverName}}Router {
	paths := make([]string, 0)

    {{.importPaths}}

	return &{{.serverName}}Router{Paths: paths}
}

func (ur *{{.serverName}}Router) AddRouters(ctx context.Context, client gw.GatewayClient, info any) error {
	return client.AddRoute(ctx, info)
}
