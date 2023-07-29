package gw

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/gw/etc/cfg.tpl
var gwCfgTemplate string

func (g *GwGenerator) GenGwCfg(ctx DirContext, gctx *GwContext) error {

	fileName := filepath.Join(ctx.GetEtc().Filename, fmt.Sprintf("%v.toml", "cfg"))

	return util.With("gwCfg").GoFmt(false).Parse(gwCfgTemplate).SaveTo(map[string]any{
		"project": ctx.GetProjectName(),
	}, fileName, false)
}
