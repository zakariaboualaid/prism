package config

import "github.com/zakariaboualaid/prism/internal/service"

type Config struct {
	Http struct {
		Services []service.Service
	} `yaml:"http"`
}
