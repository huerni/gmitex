package discovery

import (
	"github.com/huerni/gmitex/core/discovery/balance"
	"github.com/huerni/gmitex/core/discovery/register"
	"github.com/huerni/gmitex/core/discovery/register/etcd"
	"github.com/huerni/gmitex/core/discovery/register/eureka"
	"github.com/huerni/gmitex/core/gateway/config"
	"sync"
)

type Discovery struct {
	instBalance *balance.InstanceBalance
	lock        *sync.RWMutex
}

var clientDiscovery = &Discovery{
	lock: &sync.RWMutex{},
}

func EnableDiscovery(conf *config.RegisterCenter) {
	var er *register.ApplicationRegisterCenter
	if conf.EurekaConfig.ServiceUrls != nil {
		er = eureka.NewEurekaRegister(conf)
	} else if conf.EtcdConfig.Endpoints != nil {
		er = etcd.NewEtcdRegister(conf)
	}

	clientDiscovery.lock.Lock()
	defer clientDiscovery.lock.Unlock()
	clientDiscovery.instBalance = balance.NewInstanceBalance(er)
}

func GetServiceInstance(serviceId string) (*register.AppInstance, error) {
	return clientDiscovery.instBalance.GetService(serviceId)
}
