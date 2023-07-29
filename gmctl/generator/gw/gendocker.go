package gw

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/gw/docker-compose.tpl
var dockerTemplate string

func (g *GwGenerator) GenDocker(ctx DirContext, gctx *GwContext) error {

	fileName := filepath.Join(gctx.Output, fmt.Sprintf("%v.yaml", "docker-compose"))

	return util.With("docker").GoFmt(false).Parse(dockerTemplate).SaveTo(map[string]any{
		"project": ctx.GetProjectName(),
	}, fileName, false)
}
