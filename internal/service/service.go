package service

import (
	"github.com/zakariaboualaid/prism/internal/backend"
	"github.com/zakariaboualaid/prism/internal/strategy"
)

type Service struct {
	Name     string
	Backends []*backend.Backend
	Strategy *strategy.Strategy
}
