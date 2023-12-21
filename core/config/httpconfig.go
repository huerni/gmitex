package config

import (
	"github.com/duke-git/lancet/netutil"
)

type HttpConf struct {
	Addr string `json:"Addr,optional"`
	Port int    `json:"Port,optional"`
}

func (c *HttpConf) FigureConfig() error {
	ip := c.Addr
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.Addr = netutil.GetInternalIp()
	}

	return nil
}

func (c *HttpConf) HasConfig() bool {
	if len(c.Addr) == 0 {
		return false
	}
	return true
}
