package service

import (
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
)

type Registry struct {
	AdminService    AdminService
	OAuthService    OAuthService
	OTPService      OTPService
	UserService     UserService
	ReferralService ReferralService
}

func NewRegistry(cfg *config.Config, builderRegistry *builder.Registry) *Registry {
	registry := &Registry{
		AdminService:    NewAdminService(cfg, builderRegistry),
		OAuthService:    NewOAuthService(cfg, builderRegistry),
		OTPService:      NewOTPService(cfg, builderRegistry),
		UserService:     NewUserService(cfg, builderRegistry),
		ReferralService: NewReferralService(cfg, builderRegistry),
	}
	return registry
}
