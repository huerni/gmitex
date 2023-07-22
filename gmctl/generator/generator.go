package generator

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"gmctl/parser"
	"gmctl/util"
	"path/filepath"
)

type Generator struct {
}

type GmContext struct {
	Src      string
	Output   string
	GoModule string
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

	// TODO:检查工具是否安装完全
	// err = g.Prepare()

	p := parser.NewProtoParser()
	proto, err := p.Parse(gctx.Src)
	if err != nil {
		return err
	}

	//err = g.GenGoMod(gctx, abs)
	//if err != nil {
	//	return err
	//}

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

	err = g.GenHandler(dirCtx, proto, gctx)
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

	err = g.GenCfg(dirCtx, proto)
	if err != nil {
		return err
	}

	err = g.GenMain(dirCtx, proto)
	if err != nil {
		return err
	}

	return nil
}
