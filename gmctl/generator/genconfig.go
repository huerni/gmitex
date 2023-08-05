package generator

import (
	_ "embed"
	"fmt"
	"github.com/huerni/gmitex/gmctl/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"path/filepath"
)

//go:embed tpl/server/internal/config/config.tpl
var configTemplate string

func (g *Generator) GenConfig(ctx DirContext, proto parser.Proto) error {
	configfileName := "config"
	fileName := filepath.Join(ctx.GetConfig().Filename, fmt.Sprintf("%v.go", configfileName))

	return util.With("config").GoFmt(true).Parse(configTemplate).SaveTo(map[string]any{}, fileName, false)

}
