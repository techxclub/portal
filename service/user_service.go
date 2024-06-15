package service

import (
	"context"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error)
	GetProfile(ctx context.Context, req domain.UserProfileParams) (*domain.UserProfile, error)
	GetUsers(ctx context.Context, req domain.UserProfileParams) (*domain.Users, error)
	GetCompanies(ctx context.Context) (*domain.Companies, error)
}

type userService struct {
	cfg      config.Config
	registry *builder.Registry
}

func NewUserService(cfg config.Config, registry *builder.Registry) UserService {
	return &userService{
		cfg:      cfg,
		registry: registry,
	}
}

func (u userService) RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error) {
	user, err := u.registry.UsersRepo.CreateUser(ctx, userDetails)
	if err != nil {
		return nil, err
	}

	authToken, err := domain.GenerateToken(user.UserID, u.cfg.Auth)
	if err != nil {
		return nil, err
	}

	return &domain.Registration{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (u userService) GetProfile(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error) {
	users, err := u.registry.UsersRepo.GetUserForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userService) GetUsers(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error) {
	users, err := u.registry.UsersRepo.GetUsersForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userService) GetCompanies(ctx context.Context) (*domain.Companies, error) {
	companies, err := u.registry.UsersRepo.GetCompanies(ctx)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
