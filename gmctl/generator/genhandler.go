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

//go:embed tpl/server/internal/handler/handler.tpl
var handlerTemplate string

//go:embed tpl/server/internal/handler/handler-func.tpl
var handlerfuncTemplate string

func (g *Generator) GenHandler(ctx DirContext, proto parser.Proto, gctx *GmContext) error {

	return g.genHandlerInCompatibility(ctx, proto, gctx)
}

func (g *Generator) genHandlerInCompatibility(ctx DirContext, proto parser.Proto, gctx *GmContext) error {
	dir := ctx.GetHandler()
	service := proto.Service[0]
	serverFileName := fmt.Sprintf("%sHandler", service.Name)
	serverFile := filepath.Join(dir.Filename, serverFileName+".go")

	imports := collection.NewSet()
	serviceImport := fmt.Sprintf(`"%v"`, ctx.GetService().Package)
	errnoImport := fmt.Sprintf(`"%v"`, ctx.GetErrno().Package)
	pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	svcImport := fmt.Sprintf(`"%v"`, ctx.GetSvc().Package)
	imports.AddStr(serviceImport, errnoImport, pbImport, svcImport)
	funcList, err := g.genFunctions(proto.PbPackage, service)
	if err != nil {
		return err
	}
	notStream := false
	for _, rpc := range service.RPC {
		if !rpc.StreamsRequest && !rpc.StreamsReturns {
			notStream = true
			break
		}
	}
	return util.With("handler").GoFmt(true).Parse(handlerTemplate).SaveTo(map[string]any{
		"head":    "// no edit",
		"imports": strings.Join(imports.KeysStr(), pathx.NL),
		"unimplementedServer": fmt.Sprintf("%s.Unimplemented%sServer", proto.PbPackage,
			stringx.From(service.Name).ToCamel()),
		"serverName": stringx.From(service.Name).ToCamel(),
		"funcs":      strings.Join(funcList, pathx.NL),
		"notStream":  notStream,
	}, serverFile, true)
}

func (g *Generator) genFunctions(goPackage string, service parser.Service) ([]string, error) {
	var (
		functionList []string
		servicePkg   string
	)

	for _, rpc := range service.RPC {
		nameJoin := fmt.Sprintf("%s_service", service.Name)
		servicePkg = strings.ToLower(stringx.From(nameJoin).ToCamel())
		// servicename := fmt.Sprintf("%sService", stringx.From(rpc.Name).ToCamel())

		comment := parser.GetComment(rpc.Doc())
		streamServer := fmt.Sprintf("%s.%s_%s%s", goPackage, parser.CamelCase(service.Name),
			parser.CamelCase(rpc.Name), "Server")
		buffer, err := util.With("func").Parse(handlerfuncTemplate).Execute(map[string]any{
			"serverName":  stringx.From(service.Name).ToCamel(),
			"serviceName": rpc.Name,
			"method":      parser.CamelCase(rpc.Name),
			"path":        rpc.RPC.Options[0].AggregatedConstants[0].Literal.Source,
			"request":     fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.RequestType)),
			"response":    fmt.Sprintf("*%s.%s", goPackage, parser.CamelCase(rpc.ReturnsType)),
			"hasComment":  len(comment) > 0,
			"comment":     comment,
			"hasReq":      !rpc.StreamsRequest,
			"stream":      rpc.StreamsRequest || rpc.StreamsReturns,
			"notStream":   !rpc.StreamsRequest && !rpc.StreamsReturns,
			"streamBody":  streamServer,
			"servicePkg":  servicePkg,
		})
		if err != nil {
			return nil, err
		}

		functionList = append(functionList, buffer.String())
	}

	return functionList, nil
}
