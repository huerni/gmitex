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

//go:embed tpl/server/internal/config/etcdconfig.tpl
var etcdTemplate string

//go:embed tpl/server/internal/config/mysqlconfig.tpl
var mysqlTemplate string

//go:embed tpl/server/internal/config/traefikconfig.tpl
var traefikTemplate string

func (g *Generator) GenConfig(ctx DirContext, proto parser.Proto) error {
	configfileName := "config"
	fileName := filepath.Join(ctx.GetConfig().Filename, fmt.Sprintf("%v.go", configfileName))

	err := util.With("config").GoFmt(true).Parse(etcdTemplate).SaveTo(map[string]any{}, filepath.Join(ctx.GetConfig().Filename, "etcdconfig.go"), false)
	if err != nil {
		return err
	}

	err = util.With("config").GoFmt(true).Parse(mysqlTemplate).SaveTo(map[string]any{}, filepath.Join(ctx.GetConfig().Filename, "mysqlconfig.go"), false)
	if err != nil {
		return err
	}

	err = util.With("config").GoFmt(true).Parse(traefikTemplate).SaveTo(map[string]any{}, filepath.Join(ctx.GetConfig().Filename, "traefikconfig.go"), false)
	if err != nil {
		return err
	}

	return util.With("config").GoFmt(true).Parse(configTemplate).SaveTo(map[string]any{}, fileName, false)

}
