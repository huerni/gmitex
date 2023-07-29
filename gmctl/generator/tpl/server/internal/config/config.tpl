package config

import (
	"fmt"
	"github.com/duke-git/lancet/netutil"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Prefix string   `json:"prefix"`
	Grpc   RpcConf  `json:"grpc"`
	Http   HttpConf `json:"http"`

	Etcd    EtcdConf    `json:"etcd,option"`
	Mysql   MysqlConf   `json:"mysql,option"`
	Traefik TraefikConf `json:"traefik,option"`
}

var (
	Cfg *Config
)

func InitConfig(filePath string) (*Config, error) {
	if Cfg == nil {
		Cfg = &Config{}
		conf.MustLoad(filePath, Cfg)
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

func FigureConf(c *Config) error {
	err := FigureIP(c)
	if err != nil {
		return err
	}

	if HasMysql(c) {
		err = FigureMysql(c)
		if err != nil {
			return err
		}
	}

	return nil
}

type RpcConf struct {
	Name        string `json:"name"`
	RpcListenOn string `json:"listenOn"`
}

type HttpConf struct {
	HttpListenOn string `json:"listenOn"`
}

func FigureIP(c *Config) error {
	ip := c.Grpc.RpcListenOn[:len(c.Grpc.RpcListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Grpc.RpcListenOn = netutil.GetInternalIp() + ":" + c.Grpc.RpcListenOn[len(c.Grpc.RpcListenOn)-4:]
		fmt.Println(c.Grpc.RpcListenOn)
	}
	ip = c.Http.HttpListenOn[:len(c.Http.HttpListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Http.HttpListenOn = netutil.GetInternalIp() + ":" + c.Http.HttpListenOn[len(c.Http.HttpListenOn)-4:]
		fmt.Println(c.Http.HttpListenOn)
	}

	return nil
}
