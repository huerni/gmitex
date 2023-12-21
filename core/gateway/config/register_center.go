package config

import "time"

type RegisterCenter struct {
	RefreshFrequency time.Duration `yaml:"refreshFrequency"`
	EurekaConfig     struct {
		ServiceUrls []string `yaml:"serviceUrls,omitempty"`
	}
	EtcdConfig struct {
		Prefix    string   `yaml:"prefix"`
		Endpoints []string `yaml:"endpoints"`
	} `yaml:"etcd"`
}
