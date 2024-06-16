package service

import (
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
)

type Registry struct {
	AdminService    AdminService
	UserService     UserService
	AuthService     AuthService
	ReferralService ReferralService
}

func NewRegistry(cfg *config.Config, builderRegistry *builder.Registry) *Registry {
	registry := &Registry{
		AdminService:    NewAdminService(cfg, builderRegistry),
		UserService:     NewUserService(cfg, builderRegistry),
		AuthService:     NewAuthService(cfg, builderRegistry),
		ReferralService: NewReferralService(cfg, builderRegistry),
	}
	return registry
}
