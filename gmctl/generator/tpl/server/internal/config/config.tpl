package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/huerni/gmitex/pkg/config"
)

type Config struct {
	Prefix string          `json:"prefix"`
	Grpc   config.RpcConf  `json:"grpc"`
	Http   config.HttpConf `json:"http"`

	Etcd    config.EtcdConf    `json:"etcd,option"`
	Mysql   config.MysqlConf   `json:"mysql,option"`
	Traefik config.TraefikConf `json:"traefik,option"`
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
	err := c.Grpc.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Http.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Etcd.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Mysql.FigureConfig()
	if err != nil {
		return err
	}
	if c.Mysql.DSN == "" && c.Mysql.HasConfig() {
		err := config.GetFigureFromEtcd(c.Prefix, c.Etcd.Hosts, &c.Mysql)
		if err != nil {
			return err
		}
	}

	err = c.Traefik.FigureConfig()
	if err != nil {
		return err
	}

	return nil
}
