package app

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	{{.imports}}
)

type GrpcServer struct {
	server             *grpc.Server
	options            []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (g *GrpcServer) Start(c *config.Config) error {
	lis, err := net.Listen("tcp", c.Grpc.RpcListenOn)
	if err != nil {
		return err
	}
	unaryInterceptorOption := grpc.ChainUnaryInterceptor(g.unaryInterceptors...)
	streamInterceptorOption := grpc.ChainStreamInterceptor(g.streamInterceptors...)
	options := append(g.options, unaryInterceptorOption, streamInterceptorOption)

	ctx := svc.NewServiceContext(c)
	srv := handler.New{{.serverName}}Server(ctx)
	g.server = grpc.NewServer(options...)

	pb.Register{{.serverName}}Server(g.server, srv)

	return g.server.Serve(lis)
}

func (g *GrpcServer) AddOptions(options ...grpc.ServerOption) {
	g.options = append(g.options, options...)
}

func (g *GrpcServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	g.streamInterceptors = append(g.streamInterceptors, interceptors...)
}

func (g *GrpcServer) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	g.unaryInterceptors = append(g.unaryInterceptors, interceptors...)
}

func (g *GrpcServer) Stop() error {
	g.server.GracefulStop()
	fmt.Println("stop rpc server...")
	return nil
}
