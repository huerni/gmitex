package gw

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/gw/cmd/main.tpl
var gwMainTemplate string

func (g *GwGenerator) GenGwMain(ctx DirContext, gctx *GwContext) error {

	fileName := filepath.Join(ctx.GetCmd().Filename, fmt.Sprintf("%v.go", "main"))

	return util.With("gwmain").GoFmt(true).Parse(gwMainTemplate).SaveTo(map[string]any{}, fileName, false)
}
