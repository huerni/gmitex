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

//go:embed tpl/server/internal/app/httpserver.tpl
var httpServerTemplate string

func (g *Generator) GenHttpServer(ctx DirContext, proto parser.Proto) error {
	httpFileName := "httpServer"

	fileName := filepath.Join(ctx.GetApp().Filename, fmt.Sprintf("%v.go", httpFileName))

	imports := collection.NewSet()
	pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	imports.AddStr(pbImport)

	return util.With("httpServer").GoFmt(true).Parse(httpServerTemplate).SaveTo(map[string]any{
		"imports": strings.Join(imports.KeysStr(), pathx.NL),
	}, fileName, false)
}
