package config

import (
    "fmt"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Prefix string `json:"prefix"`

	Etcd  EtcdConf  `json:"etcd,option"`
	Mysql MysqlConf `json:"mysql,option"`
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
	err := FigureMysql(c)
	if err != nil {
		return err
	}

	return nil
}


type EtcdConf struct {
	Hosts []string `json:"hosts"`
	Key   string   `json:"key"`
}

func HasEtcd(c *Config) bool {
	if len(c.Etcd.Hosts) > 0 {
		return true
	}

	return false
}

type MysqlConf struct {
	Username string `json:"username,option"`
	Password string `json:"password,option"`
	Address  string `json:"address,option"`
	Database string `json:"database,option"`
	Other    string `json:"other,option"`
	DSN      string `json:"dsn,option"`
}

func FigureMysql(c *Config) error {
	if c.Mysql.DSN != "" {
		return nil
	}

	if c.Mysql.Username != "" && c.Mysql.Password != "" && c.Mysql.Database != "" && c.Mysql.Address != "" {
		c.Mysql.DSN = fmt.Sprintf("%v:%v@tcp(%v)/%v%v",
			c.Mysql.Username, c.Mysql.Password, c.Mysql.Address, c.Mysql.Database, c.Mysql.Other)
		return nil
	}
	return nil
}

func HasMysql(c *Config) bool {
	if c.Mysql.DSN != "" {
		return true
	}
	return false
}
