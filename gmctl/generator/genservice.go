package generator

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"gmctl/parser"
	"path/filepath"
	"strings"
)

//go:embed  tpl/server/internal/service/service.tpl
var serviceTemplate string

//go:embed tpl/server/internal/service/service-func.tpl
var serviceFuncTemplate string

func (g *Generator) GenService(ctx DirContext, proto parser.Proto) error {

	return g.GenServiceInCompatibility(ctx, proto)
}

func (g *Generator) GenServiceInCompatibility(ctx DirContext, proto parser.Proto) error {
	dir := ctx.GetService()
	server := proto.Service[0].Service.Name
	for _, rpc := range proto.Service[0].RPC {
		serviceName := fmt.Sprintf("%s", stringx.From(rpc.Name).ToCamel())

		filename := filepath.Join(dir.Filename, FirstLower(serviceName)+".go")
		functions, err := g.GenServiceFunction(server, proto.PbPackage, serviceName, rpc)
		if err != nil {
			return err
		}
		imports := collection.NewSet()
		imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetSvc().Package))
		imports.AddStr(fmt.Sprintf(`"%v"`, ctx.GetPb().Package))
		err = util.With("service").GoFmt(true).Parse(serviceTemplate).SaveTo(map[string]any{
			"serviceName": fmt.Sprintf("%s", stringx.From(rpc.Name).ToCamel()),
			"functions":   functions,
			"packageName": "service",
			"imports":     strings.Join(imports.KeysStr(), pathx.NL),
		}, filename, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) GenServiceFunction(serverName, goPackage, serviceName string, rpc *parser.RPC) (string,
	error) {
	functions := make([]string, 0)
	comment := parser.GetComment(rpc.Doc())
	streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(serverName),
		parser.CamelCase(rpc.Name), "Server")
	buffer, err := util.With("fun").Parse(serviceFuncTemplate).Execute(map[string]any{
		"serviceName":  serviceName,
		"method":       parser.CamelCase(rpc.Name),
		"hasReq":       !rpc.StreamsRequest,
		"request":      fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
		"hasReply":     !rpc.StreamsRequest && !rpc.StreamsReturns,
		"response":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"responseType": fmt.Sprintf("%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
		"stream":       rpc.StreamsRequest || rpc.StreamsReturns,
		"streamBody":   streamServer,
		"hasComment":   len(comment) > 0,
		"comment":      comment,
	})
	if err != nil {
		return "", err
	}
	functions = append(functions, buffer.String())
	return strings.Join(functions, pathx.NL), nil
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
