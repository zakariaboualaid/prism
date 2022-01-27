package config

import "github.com/zakariaboualaid/prism/internal/service"

type PersonConfig struct {
	Http struct {
		Services []*service.Service `yaml:"services"`
	} `yaml:"http"`
}
