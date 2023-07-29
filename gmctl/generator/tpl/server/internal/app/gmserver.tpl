package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	{{.imports}}
	"github.com/huerni/gmitex/pkg/etcd"
	"github.com/huerni/gmitex/pkg/gw/traefik"
	"syscall"
)

type GmServer struct {
	RpcServer  *GrpcServer
	HttpServer *HtServer
	{{.serverName}}Router *router.{{.serverName}}Router
	Cfg        *config.Config
}

func NewGmServer(c *config.Config) *GmServer {

	return &GmServer{
		RpcServer:  NewGrpcServer(),
		HttpServer: NewHtServer(),
		{{.serverName}}Router: router.New{{.serverName}}Router(),
		Cfg:        c,
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
}

func (g *GmServer) RegisterComponents(ctx context.Context) {
	if config.HasEtcd(g.Cfg) {
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, &etcd.ServerInfo{
			ServerKey: g.Cfg.Etcd.Key,
			Data:      map[string]string{"rpc": g.Cfg.Grpc.RpcListenOn, "http": g.Cfg.Http.HttpListenOn},
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if config.HasTraefik(g.Cfg) && g.Cfg.Traefik.Provider == "etcd" {
		err := g.{{.serverName}}Router.AddRouters(ctx, traefik.NewTClient(), &traefik.RouterInfo{
			Server: g.Cfg.Etcd.Key,
			Paths:  g.{{.serverName}}Router.Paths,
			Url:    g.Cfg.Http.HttpListenOn,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if config.HasMysql(g.Cfg) {
		if config.HasMysql(g.Cfg) {
			db.Init(g.Cfg)
		}
	}

}

func (g *GmServer) startRpc() error {
	err := g.RpcServer.Start(g.Cfg)
	if err != nil {
		return err
	}

	return nil
}

func (g *GmServer) startHttp(ctx context.Context) error {
	err := g.HttpServer.Start(ctx, g.Cfg.Grpc.RpcListenOn, g.Cfg.Http.HttpListenOn)
	if err != nil && err != http.ErrServerClosed {
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
