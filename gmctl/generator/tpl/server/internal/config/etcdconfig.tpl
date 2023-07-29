package config

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
