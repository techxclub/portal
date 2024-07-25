package admin

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/handler/admin/request"
	"github.com/techx/portal/handler/admin/response"
	"github.com/techx/portal/service"
)

func UserReferralListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminUserReferralListRequest,
		func(ctx context.Context, req request.AdminUserReferralListRequest) (*domain.Referrals, error) {
			return serviceRegistry.ReferralService.FetchReferrals(ctx, req.ToFetchReferralParams())
		},
		response.NewAdminUserReferralListResponse,
	)
}
