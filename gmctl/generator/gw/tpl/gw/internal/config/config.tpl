package config

import (
	"github.com/huerni/gmitex/core/gateway/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/huerni/gmitex/core/logger"
)

var appConfig = &AppConfig{}

type AppConfig struct {
	// eureka 配置信息
	RegisterCenter *config.RegisterCenter `yaml:"registerCenter"`

	// 路由配置信息
	GatewayRouter *config.Routers `yaml:"gateway"`

	// web 服务配置信息
	Server *config.ServerConfig `yaml:"server"`
}

func GetServerConfig() *config.ServerConfig {
	return appConfig.Server
}

func GetGatewayRouter() *config.Routers {
	return appConfig.GatewayRouter
}

func GetRegisterCenter() *config.RegisterCenter {
	return appConfig.RegisterCenter
}

func init() {
	logger.Info("加载路由配置信息")
	data, err := ioutil.ReadFile("etc/properties.yml")
	if err != nil {
		logger.Error("从 conf/properties.yml 配置文件中加载配置信息失败")
	}
	yaml.Unmarshal(data, appConfig)
	logger.Info("路由配置信息解析完成")
}
