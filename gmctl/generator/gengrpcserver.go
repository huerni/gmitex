package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"gmctl/parser"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/internal/app/grpcserver.tpl
var grpcServerTemplate string

func (g *Generator) GenGrpcServer(ctx DirContext, proto parser.Proto) error {
	grpcFileName := "grpcServer"

	fileName := filepath.Join(ctx.GetApp().Filename, fmt.Sprintf("%v.go", grpcFileName))

	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	handlerImport := fmt.Sprintf(`"%v"`, ctx.GetHandler().Package)
	svcImport := fmt.Sprintf(`"%v"`, ctx.GetSvc().Package)
	pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	imports.AddStr(configImport, handlerImport, svcImport, pbImport)

	return util.With("grpcServer").GoFmt(true).Parse(grpcServerTemplate).SaveTo(map[string]any{
		"imports":    strings.Join(imports.KeysStr(), pathx.NL),
		"serverName": stringx.From(ctx.GetServerName()).ToCamel(),
	}, fileName, false)
}
