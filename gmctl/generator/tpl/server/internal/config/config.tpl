package config

import (
	"fmt"
	"github.com/duke-git/lancet/netutil"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/huerni/gmitex/pkg/etcd"
	"strings"
)

type Config struct {
	Prefix string             `json:"prefix"`
	Grpc   zrpc.RpcServerConf `json:"grpc"`
	Http   HttpConf           `json:"http"`

	Etcd  discov.EtcdConf `json:"etcd,option"`
	Mysql MysqlConf       `json:"mysql,option"`
}

var (
	Cfg    *Config
	EtcdSd *etcd.ServiceDiscovery
)

func InitConfig(filePath string) (*Config, error) {
	if Cfg == nil {
		Cfg = &Config{}
		conf.MustLoad(filePath, Cfg)
		InitEtcdWatcher(Cfg)
		err := FigureConf(Cfg)
		if err != nil {
			return nil, err
		}
	}

	return Cfg, nil
}

func GetConfig() *Config {
	return Cfg
}

func GetEtcdSd() *etcd.ServiceDiscovery {
	return EtcdSd
}

func FigureConf(c *Config) error {
	err := FigureIP(c)
	if err != nil {
		return err
	}
	err = FigureEtcdConf(c)
	if err != nil {
		return err
	}

	err = FigureMysql(c)
	if err != nil {
		return err
	}

	return nil
}

func FigureEtcdConf(c *Config) error {
	c.Grpc.Etcd.Hosts = c.Etcd.Hosts
	if !strings.Contains(c.Grpc.Etcd.Key, c.Prefix) {
		c.Grpc.Etcd.Key = c.Prefix + c.Etcd.Key
	}

	return nil
}

func InitEtcdWatcher(c *Config) {
	if EtcdSd == nil {
		EtcdSd = etcd.NewServiceDiscovery(c.Etcd.Hosts)
		err := EtcdSd.WatchService(c.Prefix)
		if err != nil {
			panic(err)
		}

	}
}

type HttpConf struct {
	HttpListenOn string `json:"listenOn"`
}

type MysqlConf struct {
	Key string `json:"key,option"`
	DSN string `json:"dsn,option"`
}

func (c Config) HasMysql() bool {
	return len(c.Mysql.DSN) > 10
}

func FigureMysql(c *Config) error {
	if c.Mysql.Key == "" {
		c.Mysql.Key = c.Prefix + "mysql"
	}

	return nil
}

func FigureIP(c *Config) error {
	ip := c.Grpc.ListenOn[:len(c.Grpc.ListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Grpc.ListenOn = netutil.GetInternalIp() + ":" + c.Grpc.ListenOn[len(c.Grpc.ListenOn)-4:]
		fmt.Println(c.Grpc.ListenOn)
	}
	ip = c.Http.HttpListenOn[:len(c.Http.HttpListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Http.HttpListenOn = netutil.GetInternalIp() + ":" + c.Http.HttpListenOn[len(c.Http.HttpListenOn)-4:]
		fmt.Println(c.Http.HttpListenOn)
	}

	return nil
}
