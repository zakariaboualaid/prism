package service

import (
	"github.com/zakariaboualaid/prism/internal/backend"
	"github.com/zakariaboualaid/prism/internal/strategy"
)

// Service holds data about desired services
type Service struct {
	Name     string             `yaml:"name"`
	Backends []*backend.Backend `yaml:"backends"`
	Strategy *strategy.Strategy `yaml:"strategy"`
}
