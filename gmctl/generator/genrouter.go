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

//go:embed tpl/server/internal/router/router.tpl
var RouterTemplate string

func (g *Generator) GenRouter(ctx DirContext, proto parser.Proto) error {
	dir := ctx.GetRouter()
	service := proto.Service[0]
	serverFileName := fmt.Sprintf("%sRouter", service.Name)
	serverFile := filepath.Join(dir.Filename, serverFileName+".go")
	importPaths := collection.NewSet()

	for _, rpc := range service.RPC {
		path := rpc.RPC.Options[0].AggregatedConstants[0].Literal.Source
		importPath := fmt.Sprintf(`paths = append(paths, "%s")`, path)
		importPaths.AddStr(importPath)
	}
	return util.With("router").GoFmt(true).Parse(RouterTemplate).SaveTo(map[string]any{
		"serverName":  stringx.From(service.Name).ToCamel(),
		"importPaths": strings.Join(importPaths.KeysStr(), pathx.NL),
	}, serverFile, true)
}
