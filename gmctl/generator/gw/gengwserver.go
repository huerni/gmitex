package gw

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/gw/internal/app/gwserver.tpl
var gwServerTemplate string

func (g *GwGenerator) GenGwServer(ctx DirContext, gctx *GwContext) error {

	fileName := filepath.Join(ctx.GetApp().Filename, fmt.Sprintf("%v.go", "gwserver"))

	return util.With("gwserver").GoFmt(true).Parse(gwServerTemplate).SaveTo(map[string]any{}, fileName, false)
}
