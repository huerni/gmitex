package main

import (
	"context"
	"fmt"
    {{.imports}}
)

func main() {
	Init()
	Run()
}

func Init() {
	_, err := config.InitConfig("etc/cfg.toml")
	if err != nil {
		fmt.Println("初始化config失败:", err)
	}
	//db.Init()
	c := config.GetConfig()
	var tc traefik.TraefikClient
	err = tc.AddRoute(context.Background(), &traefik.RouterInfo{
		Server: c.Etcd.Key,
		Path:   "/api/v1/user",
		Url:    c.Http.HttpListenOn,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func Run() (err error) {

	go app.GrpcServer()

	app.HttpServer()
	return nil
}
