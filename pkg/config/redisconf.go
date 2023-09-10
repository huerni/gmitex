package config

type RedisConf struct {
	Addr     string `json:"addr,optional"`
	Password string `json:"password,optional"`
}

func (c *RedisConf) FigureConfig() error {
	return nil
}

func (c *RedisConf) HasConfig() bool {
	if len(c.Addr) > 0 {
		return true
	}

	return false
}
