package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gmctl/parser"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/etc/cfg.tpl
var cfgTemplate string

func (g *Generator) GenCfg(ctx DirContext, proto parser.Proto) error {
	FileName := "cfg"

	fileName := filepath.Join(ctx.GetEtc().Filename, fmt.Sprintf("%v.toml", FileName))

	return util.With("etc").Parse(cfgTemplate).SaveTo(map[string]any{
		"serverName": strings.ToLower(ctx.GetServerName()),
	}, fileName, false)
}
