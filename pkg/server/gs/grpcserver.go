package gs

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type (
	RegisterFn func(*grpc.Server)

	GrpcServer struct {
		server             *grpc.Server
		register           RegisterFn
		options            []grpc.ServerOption
		streamInterceptors []grpc.StreamServerInterceptor
		unaryInterceptors  []grpc.UnaryServerInterceptor
	}
)

func NewGrpcServer(register RegisterFn) *GrpcServer {
	return &GrpcServer{register: register}
}

func (g *GrpcServer) Start(RpcListenOn string) error {
	lis, err := net.Listen("tcp", RpcListenOn)
	if err != nil {
		return err
	}
	unaryInterceptorOption := grpc.ChainUnaryInterceptor(g.unaryInterceptors...)
	streamInterceptorOption := grpc.ChainStreamInterceptor(g.streamInterceptors...)
	options := append(g.options, unaryInterceptorOption, streamInterceptorOption)

	g.server = grpc.NewServer(options...)
	g.register(g.server)

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
