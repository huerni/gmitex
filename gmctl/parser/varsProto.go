package parser

import "github.com/emicklei/proto"

type Proto struct {
	Src       string
	Name      string
	Package   Package
	PbPackage string
	GoPackage string
	Import    []Import
	Message   []Message
	Service   []Service
}

type Package struct {
	*proto.Package
}

type Import struct {
	*proto.Import
}

type Message struct {
	*proto.Message
}

type Service struct {
	*proto.Service
	RPC []*RPC
}

type RPC struct {
	*proto.RPC
}
