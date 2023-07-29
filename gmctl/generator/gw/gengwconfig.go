package gw

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/gw/internal/config/config.tpl
var gwConfigTemplate string

func (g *GwGenerator) GenGwConfig(ctx DirContext, gctx *GwContext) error {

	fileName := filepath.Join(ctx.GetConfig().Filename, fmt.Sprintf("%v.go", "config"))

	return util.With("gwconfig").GoFmt(true).Parse(gwConfigTemplate).SaveTo(map[string]any{}, fileName, false)
}
