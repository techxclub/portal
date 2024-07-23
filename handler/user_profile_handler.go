package handler

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/request"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
)

func UserFetchProfileHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewUserFetchProfileRequest,
		func(ctx context.Context, req request.UserFetchProfileRequest) (*domain.User, error) {
			return serviceRegistry.UserService.GetUser(ctx, req.ToFetchUserParams())
		},
		response.NewUserProfileResponse,
	)
}

func UserUpdateProfileHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewUserUpdateProfileRequest,
		func(ctx context.Context, req request.UserUpdateProfileRequest) (*domain.User, error) {
			return serviceRegistry.UserService.UpdateUser(ctx, req.ToUserDomainObject())
		},
		response.NewUserProfileResponse,
	)
}
