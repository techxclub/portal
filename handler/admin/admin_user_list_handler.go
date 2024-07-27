package admin

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/handler/admin/request"
	"github.com/techx/portal/handler/admin/response"
	"github.com/techx/portal/handler/composers"
	"github.com/techx/portal/service"
)

func UserListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminUserListRequest,
		func(ctx context.Context, req request.AdminUserListRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToFetchUserParams())
		},
		func(ctx context.Context, domainObj domain.Users) (response.UserListResponse, composers.HTTPMetadata) {
			return response.NewUserListResponse(ctx, domainObj)
		},
	)
}
