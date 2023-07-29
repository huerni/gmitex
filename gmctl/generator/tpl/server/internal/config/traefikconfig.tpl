package config

type TraefikConf struct {
	Provider string `json:"provider"`
}

func HasTraefik(c *Config) bool {
	if len(c.Traefik.Provider) > 0 {
		return true
	}
	return false
}
