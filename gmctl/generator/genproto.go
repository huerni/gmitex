package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"gmctl/parser"
	"path/filepath"
)

//go:embed tpl/server/proto/proto.tpl
var protoTemplate string

func (g *Generator) GenProto(ctx DirContext, proto parser.Proto) error {
	protofileName := ctx.GetServerName()

	fileName := filepath.Join(ctx.GetProtoGo().Filename, fmt.Sprintf("%v.proto", protofileName))

	return util.With("proto").GoFmt(false).Parse(protoTemplate).SaveTo(map[string]any{
		"package": ctx.GetServerName(),
	}, fileName, false)
}
