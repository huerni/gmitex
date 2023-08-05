package config

import (
	"github.com/huerni/gmitex/pkg/config"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Prefix string `json:"prefix"`

	Etcd  config.EtcdConf  `json:"etcd,option"`
	Mysql config.MysqlConf `json:"mysql,option"`
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
	err := c.Etcd.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Mysql.FigureConfig()
	if err != nil {
		return err
	}

	return nil
}
