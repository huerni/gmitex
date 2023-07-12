package traefik

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/etcd"
)

type (
	TraefikClient struct{}

	RouterInfo struct {
		endpoints []string
		Server    string
		Path      string
		Url       string
	}
)

func (tc *TraefikClient) AddRoute(ctx context.Context, info any) error {
	routerInfo, ok := info.(RouterInfo)
	if !ok {
		return fmt.Errorf("转换router info失败")
	}
	err := etcd.PutWithKV(ctx, routerInfo.endpoints, fmt.Sprintf("traefik/http/routers/%s/rule", routerInfo.Server), fmt.Sprintf("PathPrefix(`%s`)", routerInfo.Path))
	if err != nil {
		return err
	}
	err = etcd.PutWithKV(ctx, routerInfo.endpoints, fmt.Sprintf("traefik/http/routers/%s/service", routerInfo.Server), routerInfo.Server)
	if err != nil {
		return err
	}
	err = etcd.PutWithKV(ctx, routerInfo.endpoints, fmt.Sprintf("traefik/http/services/%s/loadbalancer/servers/0/url", routerInfo.Server), "http://"+routerInfo.Url)
	if err != nil {
		return err
	}

	return nil
}
