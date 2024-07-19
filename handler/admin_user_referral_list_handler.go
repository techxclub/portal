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

func AdminUserReferralListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewAdminUserReferralListRequest,
		func(ctx context.Context, req request.AdminUserReferralListRequest) (*domain.Referrals, error) {
			return serviceRegistry.ReferralService.FetchReferrals(ctx, req.ToFetchReferralParams())
		},
		response.NewAdminUserReferralListResponse,
	)
}
