package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huerni/gmitex/core/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdClient interface {
	GetWithPrefix(ctx context.Context, endpoints []string, prefix string)
	PutWithKV(ctx context.Context, endpoints []string, key string, val string)
	PutWithInfo(ctx context.Context, endpoints []string, info *AppInstance)
}

type AppInstance struct {
	// 实例ID
	InstanceId string
	// 服务 IP 地址
	IpAddr string
	// 服务端口号
	Port int
	// 服务状态
	Status string
	// 是否启用 SSL
	Secure bool
}

func GetWithPrefix(ctx context.Context, endpoints []string, prefix string) (map[string]string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	defer cli.Close()

	if err != nil {
		return nil, err
	}

	kv := clientv3.NewKV(cli)
	getResp, err := kv.Get(ctx, prefix, clientv3.WithPrefix())

	if err != nil {
		return nil, err
	}
	result, err := stringToMap(string(getResp.Kvs[0].Value))
	if err != nil {
		return nil, err
	}
	return result, nil

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
		fmt.Printf("start keepalive key %s for etcd registry\n", key)
		for range keepAlive {
			select {
			case <-ctx.Done():
				cli.Close()
				fmt.Printf("stop keepalive key %s for etcd registry\n", key)
				return
			default:
			}
		}
	}()

	return nil
}

func PutWithInfo(ctx context.Context, endpoints []string, prefix string, serverName string, info *AppInstance) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	resp, err := cli.Grant(ctx, 60)
	key := fmt.Sprintf("%s/%s/%d", prefix, serverName, resp.ID)
	info.InstanceId = fmt.Sprintf("%d", resp.ID)
	val, err := json.Marshal(info)
	_, err = cli.Put(ctx, key, string(val), clientv3.WithLease(resp.ID))

	if err != nil {
		return err
	}

	keepAlive, err := cli.KeepAlive(ctx, resp.ID)
	go func() {
		// eat keepAlive channel to keep related lease alive.
		logger.Info("注册【", key, "】服务成功，保持keepalive")
		for range keepAlive {
			select {
			case <-ctx.Done():
				cli.Close()
				logger.Info(key, "服务掉线")
				return
			default:
			}
		}
	}()

	return nil
}

// stringToMap 将map格式的字符串解析为map类型
func stringToMap(s string) (map[string]string, error) {
	var m map[string]string
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
