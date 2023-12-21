package router

import (
	"context"
	"github.com/huerni/gmitex/core/gw"
)

type {{.serverName}}Router struct {
	Paths []string
	AuthPaths []string
}

func New{{.serverName}}Router() *{{.serverName}}Router {
	paths := make([]string, 0)
    authPaths := make([]string, 0)

    {{.importPaths}}
    {{.importAuthPaths}}

	return &{{.serverName}}Router{Paths: paths, AuthPaths: authPaths}
}

func (ur *{{.serverName}}Router) AddRouters(ctx context.Context, client gw.GatewayClient, info any) error {
	return client.AddRoute(ctx, info)
}
