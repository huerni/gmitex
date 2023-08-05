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

//go:embed tpl/server/cmd/main.tpl
var mainTemplate string

func (g *Generator) GenMain(ctx DirContext, proto parser.Proto) error {
	mainFileName := "main"

	fileName := filepath.Join(ctx.GetCmd().Filename, fmt.Sprintf("%v.go", mainFileName))

	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	appImport := fmt.Sprintf(`"%v"`, ctx.GetApp().Package)
	handlerImport := fmt.Sprintf(`"%s"`, ctx.GetHandler().Package)
	svcImport := fmt.Sprintf(`"%s"`, ctx.GetSvc().Package)
	pbImport := fmt.Sprintf(`%s %s`, ctx.GetServerName(), ctx.GetPb().Package)
	imports.AddStr(configImport, appImport, handlerImport, svcImport, pbImport)

	return util.With("main").GoFmt(true).Parse(mainTemplate).SaveTo(map[string]any{
		"imports":    strings.Join(imports.KeysStr(), pathx.NL),
		"serverName": parser.CamelCase(ctx.GetServerName()),
	}, fileName, false)
}
