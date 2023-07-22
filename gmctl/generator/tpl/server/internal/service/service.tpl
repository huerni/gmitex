package {{.packageName}}

import (
    "context"

    {{.imports}}
)

type {{.serviceName}}Service struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.serviceName}}Service(ctx context.Context, svcCtx *svc.ServiceContext) *{{.serviceName}}Service {
	return &{{.serviceName}}Service{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

{{.functions}}