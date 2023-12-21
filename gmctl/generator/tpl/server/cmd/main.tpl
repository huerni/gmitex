package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/huerni/gmitex/core/http/handlers"
	"github.com/huerni/gmitex/core/http/response"
	"gmitest/internal/app"
	"gmitest/internal/config"
	"gmitest/internal/handler"
	"gmitest/internal/svc"
	gmitest "gmitest/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	g := app.NewGmServer(config.Cfg, gmitest.RegisterGmitestHandlerFromEndpoint, func(server *grpc.Server) {
		ctx := svc.NewServiceContext(config.Cfg)
		srv := handler.NewGmitestServer(ctx)
		gmitest.RegisterGmitestServer(server, srv)
	})

	g.HttpServer.AddMuxOp(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &response.CustomMarshaler{
				M: &runtime.JSONPb{
					MarshalOptions:   protojson.MarshalOptions{},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}}}),
		runtime.WithErrorHandler(handlers.ErrorHandler),
	)

	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
