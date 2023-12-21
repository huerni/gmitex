package gw

import (
	"github.com/huerni/gmitex/gmctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"path/filepath"
)

type GwGenerator struct {
}

type GwContext struct {
	ProjectName string
	Output      string
}

func NewGenerator() *GwGenerator {
	return &GwGenerator{}
}

func (g *GwGenerator) Generate(gctx *GwContext) error {
	abs, err := filepath.Abs(gctx.Output)
	if err != nil {
		return err
	}

	err = util.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return err
	}

	dirCtx, err := mkdir(projectCtx, gctx)
	// 开始生成文件

	err = g.GenDocker(dirCtx, gctx)
	if err != nil {
		return err
	}

	err = g.GenGwMain(dirCtx, gctx)
	if err != nil {
		return err
	}

	err = g.GenGwCfg(dirCtx, gctx)
	if err != nil {
		return err
	}

	err = g.GenGwConfig(dirCtx, gctx)
	if err != nil {
		return err
	}

	return nil
}
