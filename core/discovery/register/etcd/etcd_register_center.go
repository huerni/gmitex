package etcd

import (
	"context"
	"encoding/json"
	"github.com/huerni/gmitex/core/discovery/register"
	"github.com/huerni/gmitex/core/gateway/config"
	"github.com/huerni/gmitex/core/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

const etcdApp = "/etcd/apps"

type etcdRegister struct {
	register *register.ApplicationRegisterCenter
}

func NewEtcdRegister() *register.ApplicationRegisterCenter {
	reg := register.NewApplicationRegisterCenter()
	r := &etcdRegister{
		register: reg,
	}
	r.enableEtcdClient()
	return reg
}

func (r *etcdRegister) enableEtcdClient() {
	conf := config.GetRegisterCenter()
	etcdRefresh := time.NewTicker(time.Second * conf.RefreshFrequency)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.EtcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logger.Fatal("连接ETCD失败: ", err)
	}
	go func(tick *time.Ticker) {
		defer tick.Stop()
		for {
			logger.Info("同步 Etcd 注册中心")
			resp, err := cli.Get(context.Background(), conf.EtcdConfig.Prefix, clientv3.WithPrefix())
			if err != nil {
				logger.Error("获取ETCD信息失败: ", err)
				continue
			}

			go r.updateRegister(resp)

			<-tick.C
		}
	}(etcdRefresh)
}

func (r *etcdRegister) updateRegister(resp *clientv3.GetResponse) {
	instancesmap := make(map[string][]*register.AppInstance)
	for _, kv := range resp.Kvs {
		parts := strings.Split(string(kv.Key), "/")
		var instance register.AppInstance
		err := json.Unmarshal(kv.Value, &instance)
		if err != nil {
			logger.Error("json解析失败: ", string(kv.Value))
			continue
		}
		if _, ok := instancesmap[parts[1]]; !ok {
			instancesmap[parts[1]] = make([]*register.AppInstance, 0)
		}
		instancesmap[parts[1]] = append(instancesmap[parts[1]], &instance)
	}

	for k, v := range instancesmap {
		appService := &register.AppService{
			ServiceId:  k,
			Instances:  v,
			UpdateTime: time.Now().Unix(),
		}
		r.register.UpdateApplication(appService)
	}
}
