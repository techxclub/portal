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

func AdminUserListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewAdminUserListRequest,
		func(ctx context.Context, req request.AdminUserListRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToUserProfileParams())
		},
		func(ctx context.Context, domainObj domain.Users) (response.UserListResponse, response.HTTPMetadata) {
			return response.NewUserListResponse(ctx, domainObj)
		},
	)
}
