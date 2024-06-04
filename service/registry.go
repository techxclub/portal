package service

import (
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
)

type Registry struct {
	UserService UserService
}

func NewRegistry(cfg config.Config, builderRegistry *builder.Registry) *Registry {
	registry := &Registry{
		UserService: NewUserService(cfg, builderRegistry),
	}
	return registry
}
