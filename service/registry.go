package service

import (
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
)

type Registry struct {
	UserService UserService
}

func NewRegistry(_ config.Config, _ *builder.Registry) *Registry {
	registry := &Registry{
		UserService: NewUserService(),
	}
	return registry
}
