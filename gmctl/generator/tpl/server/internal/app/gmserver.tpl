package app

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/etcd"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	{{.imports}}
	"github.com/huerni/gmitex/pkg/server/gs"
	"github.com/huerni/gmitex/pkg/server/hs"
	"github.com/huerni/gmitex/pkg/gw/traefik"
	"syscall"
)

type GmServer struct {
	RpcServer  *gs.GrpcServer
	HttpServer *hs.HtServer
	{{.serverName}}Router *router.{{.serverName}}Router
	Cfg        *config.Config
}

func NewGmServer(c *config.Config, registerHandler interface{}, registerFunc gs.RegisterFn) *GmServer {

	return &GmServer{
		RpcServer:  gs.NewGrpcServer(registerFunc),
		HttpServer: hs.NewHtServer(registerHandler),
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
	if g.Cfg.Etcd.HasConfig() {
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, &etcd.ServerInfo{
			ServerKey: fmt.Sprintf("%v-%v", g.Cfg.Prefix, g.Cfg.Etcd.Key),
			Data:      map[string]string{"rpc": g.Cfg.Grpc.RpcListenOn, "hs": g.Cfg.Http.HttpListenOn},
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if g.Cfg.Traefik.HasConfig() && g.Cfg.Traefik.Provider == "etcd" {
		err := g.{{.serverName}}Router.AddRouters(ctx, traefik.NewTClient(), &traefik.RouterInfo{
			Endpoints: g.Cfg.Etcd.Hosts,
			Server:    g.Cfg.Etcd.Key,
			Paths:     g.{{.serverName}}Router.Paths,
			Url:       g.Cfg.Http.HttpListenOn,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if g.Cfg.Mysql.HasConfig() {
		err := db.Init(g.Cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (g *GmServer) startRpc() error {
	err := g.RpcServer.Start(g.Cfg.Grpc.RpcListenOn)
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
