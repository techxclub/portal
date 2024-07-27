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

func ExpireReferralHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminExpireReferralRequest,
		func(ctx context.Context, req request.AdminExpireReferralRequest) (*domain.EmptyDomain, error) {
			return serviceRegistry.ReferralService.ExpireReferrals(ctx, req.ToExpireReferralParams())
		},
		func(ctx context.Context, _ domain.EmptyDomain) (composers.SuccessResponse, composers.HTTPMetadata) {
			return composers.NewSuccessResponse(ctx)
		},
	)
}
