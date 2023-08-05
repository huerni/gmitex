package config

type TraefikConf struct {
	Provider string `json:"provider,option"`
}

func (c *TraefikConf) FigureConfig() error {
	return nil
}

func (c *TraefikConf) HasConfig() bool {
	if len(c.Provider) > 0 {
		return true
	}
	return false
}
