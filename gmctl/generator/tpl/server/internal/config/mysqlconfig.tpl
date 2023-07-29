package config

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/etcd"
)

type MysqlConf struct {
	Key      string `json:"key,option"`
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
		c.Mysql.DSN = fmt.Sprintf("%v:%v@tcp(%v)/%v%v", c.Mysql.Username, c.Mysql.Password, c.Mysql.Address, c.Mysql.Database, c.Mysql.Other)
		return nil
	}

	c.Mysql.Key = fmt.Sprintf("%v-%v", c.Prefix, c.Mysql.Key)
	resp, err := etcd.GetWithPrefix(context.Background(), c.Etcd.Hosts, c.Mysql.Key)
	if err != nil {
		return err
	}
	c.Mysql.DSN = resp["dsn"]

	return nil
}

func HasMysql(c *Config) bool {
	if c.Mysql.DSN != "" || c.Mysql.Key != "" || (c.Mysql.Username != "" && c.Mysql.Password != "" && c.Mysql.Address != "" && c.Mysql.Database != "") {
		return true
	}
	return false
}
