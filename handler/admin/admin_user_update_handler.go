package admin

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/handler/admin/request"
	"github.com/techx/portal/handler/composers"
	"github.com/techx/portal/service"
)

func UserApproveHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminUserApproveRequest,
		func(ctx context.Context, req request.AdminUserUpdateParams) (*domain.EmptyDomain, error) {
			user := req.ToUserProfile()
			return serviceRegistry.AdminService.ApproveUser(ctx, user)
		},
		func(ctx context.Context, _ domain.EmptyDomain) (composers.SuccessResponse, composers.HTTPMetadata) {
			return composers.NewSuccessResponse(ctx)
		},
	)
}

func UserUpdateHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminUserUpdateRequest,
		func(ctx context.Context, req request.AdminUserUpdateRequest) (*domain.EmptyDomain, error) {
			from, to := req.From.ToUserProfile(), req.To.ToUserProfile()
			return serviceRegistry.AdminService.UpdateUsers(ctx, from, to)
		},
		func(ctx context.Context, _ domain.EmptyDomain) (composers.SuccessResponse, composers.HTTPMetadata) {
			return composers.NewSuccessResponse(ctx)
		},
	)
}
