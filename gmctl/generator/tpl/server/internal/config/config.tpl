package config

import (
	"github.com/huerni/gmitex/core/config"
	"github.com/huerni/gmitex/core/logger"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Prefix string          `json:"prefix"`
	Grpc   config.RpcConf  `json:"grpc"`
	Http   config.HttpConf `json:"http"`

	Etcd  config.EtcdConf  `json:"etcd,option"`
	Mysql config.MysqlConf `json:"mysql,option"`
}

var (
	Cfg = &Config{}
)

func init() {
	filePath := "etc/cfg.toml"
	err := conf.Load(filePath, Cfg)
	if err != nil {
		logger.Fatal("配置信息读取失败: ", err)
	}
	err = FigureConf(Cfg)
	if err != nil {
		logger.Fatal("配置信息初始化失败: ", err)
	}
	logger.Info("配置信息初始化完成")
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

	return nil
}
