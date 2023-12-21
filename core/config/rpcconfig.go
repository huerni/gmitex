package config

import (
	"fmt"
	"github.com/duke-git/lancet/netutil"
)

type RpcConf struct {
	Name        string `json:"name"`
	RpcListenOn string `json:"listenOn"`
}

func (c *RpcConf) FigureConfig() error {
	ip := c.RpcListenOn[:len(c.RpcListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.RpcListenOn = netutil.GetInternalIp() + ":" + c.RpcListenOn[len(c.RpcListenOn)-4:]
		fmt.Println(c.RpcListenOn)
	}
	return nil
}

func (c *RpcConf) HasConfig() bool {
	if len(c.Name) == 0 || len(c.RpcListenOn) == 0 {
		return false
	}
	return true
}
