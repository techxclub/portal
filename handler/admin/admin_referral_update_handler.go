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

func ReferralUpdateHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminReferralUpdateRequest,
		func(ctx context.Context, req request.AdminReferralUpdateRequest) (*domain.EmptyDomain, error) {
			return serviceRegistry.AdminService.UpdateReferralDetails(ctx, req.ToReferralParams())
		},
		func(ctx context.Context, _ domain.EmptyDomain) (composers.SuccessResponse, composers.HTTPMetadata) {
			return composers.NewSuccessResponse(ctx)
		},
	)
}
