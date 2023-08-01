package generator

import (
	_ "embed"
	"fmt"
	"github.com/huerni/gmitex/gmctl/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/internal/app/gmserver.tpl
var gmServerTemplate string

func (g *Generator) GenGmServer(ctx DirContext, proto parser.Proto) error {
	gmFileName := "gmServer"

	fileName := filepath.Join(ctx.GetApp().Filename, fmt.Sprintf("%v.go", gmFileName))

	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	dbImport := fmt.Sprintf(`"%v"`, ctx.GetDb().Package)
	routerImport := fmt.Sprintf(`"%v"`, ctx.GetRouter().Package)
	imports.AddStr(configImport, dbImport, routerImport)

	return util.With("gmServer").GoFmt(true).Parse(gmServerTemplate).SaveTo(map[string]any{
		"imports":    strings.Join(imports.KeysStr(), pathx.NL),
		"serverName": parser.CamelCase(ctx.GetServerName()),
	}, fileName, false)
}
