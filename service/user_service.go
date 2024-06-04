package service

import (
	"context"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.User) (*domain.User, error)
}

type userService struct {
	cfg      config.Config
	registry *builder.Registry
}

func NewUserService(config config.Config, registry *builder.Registry) UserService {
	return &userService{
		cfg:      config,
		registry: registry,
	}
}

func (u userService) RegisterUser(ctx context.Context, userDetails domain.User) (*domain.User, error) {
	user, err := u.registry.UserInfoBuilder.CreateUser(ctx, userDetails)
	if err != nil {
		return nil, err
	}

	return user, nil
}
