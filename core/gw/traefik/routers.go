package traefik

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/core/etcd"
)

type TraefikClient struct {
}

type RouterInfo struct {
	Endpoints []string
	Server    string
	Paths     []string
	Url       string
}

func NewTClient() *TraefikClient {
	return &TraefikClient{}
}

func (tc *TraefikClient) AddRoute(ctx context.Context, info any) error {
	routerInfo, ok := info.(*RouterInfo)
	if !ok {
		return fmt.Errorf("转换router info失败")
	}

	for _, path := range routerInfo.Paths {
		err := etcd.PutWithKV(ctx, routerInfo.Endpoints, fmt.Sprintf("traefik/http/routers/%s/rule", routerInfo.Server), fmt.Sprintf("PathPrefix(`%s`)", path))
		if err != nil {
			return err
		}
	}

	err := etcd.PutWithKV(ctx, routerInfo.Endpoints, fmt.Sprintf("traefik/http/routers/%s/service", routerInfo.Server), routerInfo.Server)
	if err != nil {
		return err
	}
	err = etcd.PutWithKV(ctx, routerInfo.Endpoints, fmt.Sprintf("traefik/http/services/%s/loadbalancer/servers/0/url", routerInfo.Server), "http://"+routerInfo.Url)
	if err != nil {
		return err
	}

	return nil
}
