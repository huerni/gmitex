package config

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/core/etcd"
)

type MysqlConf struct {
	Key      string `json:"key,optional"`
	Username string `json:"username,optional"`
	Password string `json:"password,optional"`
	Address  string `json:"address,optional"`
	Database string `json:"database,optional"`
	Other    string `json:"other,optional"`
	DSN      string `json:"dsn,optional"`
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
