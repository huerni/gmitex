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

//go:embed tpl/server/internal/db/init.tpl
var dbTemplate string

func (g *Generator) GenDb(ctx DirContext, proto parser.Proto) error {
	dbFileName := "init"

	fileName := filepath.Join(ctx.GetDb().Filename, fmt.Sprintf("%v.go", dbFileName))

	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	imports.AddStr(configImport)

	return util.With("dbinit").GoFmt(true).Parse(dbTemplate).SaveTo(map[string]any{
		"imports": strings.Join(imports.KeysStr(), pathx.NL),
	}, fileName, false)
}
