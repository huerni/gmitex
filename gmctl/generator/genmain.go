package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"gmctl/parser"
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
	imports.AddStr(configImport)

	return util.With("main").GoFmt(true).Parse(mainTemplate).SaveTo(map[string]any{
		"imports": strings.Join(imports.KeysStr(), pathx.NL),
	}, fileName, false)
}
