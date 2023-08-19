package generator

import (
	"github.com/huerni/gmitex/gmctl/parser"
	"github.com/huerni/gmitex/gmctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"path/filepath"
)

type Generator struct {
}

type GmContext struct {
	Op          string
	Src         string
	Output      string
	GoModule    string
	ProjectName string
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(gctx *GmContext) error {
	abs, err := filepath.Abs(gctx.Output)
	if err != nil {
		return err
	}

	err = util.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	//检查工具是否安装完全
	if gctx.Op == "new" {
		err = g.Prepare()
		if err != nil {
			return err
		}
	}

	p := parser.NewProtoParser()
	proto, err := p.Parse(gctx.Src)
	if err != nil {
		return err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return err
	}

	dirCtx, err := mkdir(projectCtx, proto, gctx)
	// 开始生成文件
	err = g.GenErrno(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenGmServer(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenDb(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenHandler(dirCtx, proto, gctx)
	if err != nil {
		return err
	}

	err = g.GenRouter(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenService(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenSvc(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenConfig(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenCfg(dirCtx, proto, gctx)
	if err != nil {
		return err
	}

	err = g.GenMain(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenPb(dirCtx, gctx)
	if err != nil {
		return err
	}

	return nil
}
