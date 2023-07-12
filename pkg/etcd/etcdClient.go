package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdClient interface {
	GetWithPrefix(ctx context.Context, endpoints []string, prefix string)
	PutWithKV(ctx context.Context, endpoints []string, key string, val string)
	PutWithInfo(ctx context.Context, endpoints []string, info *ServerInfo)
}

type ServerInfo struct {
	ServerKey string
	Data      map[string]string
}

func GetWithPrefix(ctx context.Context, endpoints []string, prefix string) (*clientv3.GetResponse, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	defer cli.Close()

	if err != nil {
		return nil, err
	}

	kv := clientv3.NewKV(cli)
	getResp, err := kv.Get(context.TODO(), prefix, clientv3.WithPrefix())

	return getResp, nil
}

func PutWithKV(ctx context.Context, endpoints []string, key string, val string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	resp, err := cli.Grant(ctx, 60)
	_, err = cli.Put(ctx, key, val, clientv3.WithLease(resp.ID))

	if err != nil {
		return err
	}

	keepAlive, err := cli.KeepAlive(ctx, resp.ID)
	go func() {
		// eat keepAlive channel to keep related lease alive.
		fmt.Printf("start keepalive lease %x for etcd registry\n", resp.ID)
		for range keepAlive {
			select {
			case <-ctx.Done():
				cli.Close()
				fmt.Printf("stop keepalive lease %x for etcd registry\n", resp.ID)
				return
			default:
			}
		}
	}()

	return nil
}

func PutWithInfo(ctx context.Context, endpoints []string, info *ServerInfo) error {
	val, err := json.Marshal(info.Data)

	if err != nil {
		return err
	}

	return PutWithKV(ctx, endpoints, info.ServerKey, string(val))
}
