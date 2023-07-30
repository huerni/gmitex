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

//go:embed tpl/server/internal/errno/errno.tpl
var errnoTemplate string

func (g *Generator) GenErrno(ctx DirContext, proto parser.Proto) error {
	errnoFileName := "errno"
	fileName := filepath.Join(ctx.GetErrno().Filename, fmt.Sprintf("%v.go", errnoFileName))

	imports := collection.NewSet()
	pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	imports.AddStr(pbImport)

	return util.With("errno").GoFmt(true).Parse(errnoTemplate).SaveTo(map[string]any{
		"imports":        strings.Join(imports.KeysStr(), pathx.NL),
		"servicePackage": proto.PbPackage,
	}, fileName, true)
}
