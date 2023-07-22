{{.head}}

package handler

import (
	{{if .notStream}}"context"{{end}}

	{{.imports}}
)

type {{.serverName}}Server struct {
	svcCtx *svc.ServiceContext
	{{.unimplementedServer}}
}

func New{{.serverName}}Server(svcCtx *svc.ServiceContext) *{{.serverName}}Server {
	return &{{.serverName}}Server{
		svcCtx: svcCtx,
	}
}

{{.funcs}}
