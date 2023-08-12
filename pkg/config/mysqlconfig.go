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

func (c *MysqlConf) FigureConfig() error {
	if c.DSN != "" {
		return nil
	}

	if c.Username != "" && c.Password != "" && c.Database != "" && c.Address != "" {
		c.DSN = fmt.Sprintf("%v:%v@tcp(%v)/%v%v", c.Username, c.Password, c.Address, c.Database, c.Other)
		return nil
	}

	return nil
}

func (c *MysqlConf) HasConfig() bool {
	if c.DSN != "" || c.Key != "" || (c.Username != "" && c.Password != "" && c.Address != "" && c.Database != "") {
		return true
	}
	return false
}

func GetFigureFromEtcd(prefix string, endpoints []string, c *MysqlConf) error {
	c.Key = fmt.Sprintf("%v-%v", prefix, c.Key)
	resp, err := etcd.GetWithPrefix(context.Background(), endpoints, c.Key)
	if err != nil {
		return err
	}
	c.DSN = resp["dsn"]
	return nil
}
