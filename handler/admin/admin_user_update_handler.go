package admin

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/handler/request"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
)

func UserUpdateHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminUserUpdateRequest,
		func(ctx context.Context, req request.AdminUserUpdateRequest) (*domain.EmptyDomain, error) {
			from, to := req.From.ToUserProfile(), req.To.ToUserProfile()
			return serviceRegistry.AdminService.BulkUpdateUsers(ctx, from, to)
		},
		func(ctx context.Context, _ domain.EmptyDomain) (response.SuccessResponse, response.HTTPMetadata) {
			return response.NewSuccessResponse(ctx)
		},
	)
}
