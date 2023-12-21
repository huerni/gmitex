package app

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/core/etcd"
	"github.com/huerni/gmitex/core/logger"
	"github.com/huerni/gmitex/core/server/gs"
	"github.com/huerni/gmitex/core/server/hs"
	"gmitest/internal/config"
	"gmitest/internal/router"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type GmServer struct {
	RpcServer     *gs.GrpcServer
	HttpServer    *hs.HtServer
	GmitestRouter *router.GmitestRouter
	Cfg           *config.Config
}

func NewGmServer(c *config.Config, registerHandler interface{}, registerFunc gs.RegisterFn) *GmServer {

	return &GmServer{
		RpcServer:     gs.NewGrpcServer(registerFunc),
		HttpServer:    hs.NewHtServer(registerHandler),
		GmitestRouter: router.NewGmitestRouter(),
		Cfg:           c,
	}
}

func (g *GmServer) Start(ctx context.Context) {

	g.RegisterComponents(ctx)

	g.HttpServer.AddDialOps(grpc.WithInsecure())
	go func() {
		err := g.startRpc()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := g.startHttp(ctx)
		if err != nil {
			panic(err)
		}
	}()
	logger.Info("服务启动完成")
}

func (g *GmServer) RegisterComponents(ctx context.Context) {
	if g.Cfg.Etcd.HasConfig() {
		appInstance := etcd.AppInstance{
			InstanceId: "",
			IpAddr:     g.Cfg.Http.Addr,
			Port:       g.Cfg.Http.Port,
			Status:     "UP",
			Secure:     false,
		}
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, g.Cfg.Prefix, g.Cfg.Etcd.Key, &appInstance)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (g *GmServer) startRpc() error {
	err := g.RpcServer.Start(fmt.Sprintf("%s:%d", g.Cfg.Grpc.Addr, g.Cfg.Grpc.Port))
	if err != nil {
		logger.Error("rpc启动失败...")
		return err
	}
	return nil
}

func (g *GmServer) startHttp(ctx context.Context) error {
	err := g.HttpServer.Start(ctx, fmt.Sprintf("%s:%d", g.Cfg.Grpc.Addr, g.Cfg.Grpc.Port), fmt.Sprintf("%s:%d", g.Cfg.Http.Addr, g.Cfg.Http.Port))
	if err != nil && err != http.ErrServerClosed {
		logger.Error("Http启动失败...")
		return err
	}
	return nil
}

func (g *GmServer) stop(ctx context.Context) error {
	err := g.RpcServer.Stop()
	if err != nil {
		return err
	}

	err = g.HttpServer.Stop(ctx)
	if err != nil {
		return err
	}
	fmt.Println("stop server..")
	return nil
}

func (g *GmServer) WaitForShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	err := g.stop(ctx)
	if err != nil {
		panic(err)
	}
}
