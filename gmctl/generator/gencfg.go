package generator

import (
	_ "embed"
	"fmt"
	"github.com/huerni/gmitex/gmctl/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
	"strings"
)

//go:embed tpl/server/etc/cfg.tpl
var cfgTemplate string

func (g *Generator) GenCfg(ctx DirContext, proto parser.Proto, gctx *GmContext) error {
	FileName := "cfg"

	fileName := filepath.Join(ctx.GetEtc().Filename, fmt.Sprintf("%v.toml", FileName))

	return util.With("etc").Parse(cfgTemplate).SaveTo(map[string]any{
		"serverName":  strings.ToLower(ctx.GetServerName()),
		"projectName": gctx.ProjectName,
	}, fileName, false)
}
