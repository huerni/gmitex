package config

type EtcdConf struct {
	Hosts []string `json:"hosts,optional"`
	Key   string   `json:"key,optional"`
}

func (c *EtcdConf) FigureConfig() error {
	return nil
}

func (c *EtcdConf) HasConfig() bool {
	if len(c.Hosts) > 0 {
		return true
	}

	return false
}
