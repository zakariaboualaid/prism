package config

import (
	"github.com/zakariaboualaid/prism/internal/service"
	"github.com/zakariaboualaid/prism/internal/strategy"
)

type Config struct {
	Services []*service.Service `yaml:"backends"`
	Strategy *strategy.Strategy
}
