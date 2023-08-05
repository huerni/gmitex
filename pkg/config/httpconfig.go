package config

import (
	"fmt"
	"github.com/duke-git/lancet/netutil"
)

type HttpConf struct {
	HttpListenOn string `json:"listenOn,option"`
}

func (c *HttpConf) FigureConfig() error {
	ip := c.HttpListenOn[:len(c.HttpListenOn)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		c.HttpListenOn = netutil.GetInternalIp() + ":" + c.HttpListenOn[len(c.HttpListenOn)-4:]
		fmt.Println(c.HttpListenOn)
	}

	return nil
}

func (c *HttpConf) HasConfig() bool {
	if len(c.HttpListenOn) == 0 {
		return false
	}
	return true
}
