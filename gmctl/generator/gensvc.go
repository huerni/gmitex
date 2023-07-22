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

//go:embed tpl/server/internal/svc/servicecontext.tpl
var serviceContextTemplate string

func (g *Generator) GenSvc(ctx DirContext, proto parser.Proto) error {
	svcfileName := "servicecontext"
	fileName := filepath.Join(ctx.GetSvc().Filename, fmt.Sprintf("%v.go", svcfileName))

	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	imports.AddStr(configImport)

	return util.With("svc").GoFmt(true).Parse(serviceContextTemplate).SaveTo(map[string]any{
		"imports": strings.Join(imports.KeysStr(), pathx.NL),
	}, fileName, false)

}
