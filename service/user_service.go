package service

import (
	"context"

	"github.com/techx/portal/domain"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.User) (*domain.User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u userService) RegisterUser(_ context.Context, userDetails domain.User) (*domain.User, error) {
	return &userDetails, nil
}
