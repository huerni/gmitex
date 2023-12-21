package config

import (
	"github.com/duke-git/lancet/netutil"
)

type RpcConf struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

func (c *RpcConf) FigureConfig() error {
	ip := c.Addr
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Addr = netutil.GetInternalIp()
	}
	return nil
}

func (c *RpcConf) HasConfig() bool {
	if len(c.Name) == 0 || len(c.Addr) == 0 {
		return false
	}
	return true
}
