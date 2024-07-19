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

func ReferralListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewReferralListRequest,
		func(ctx context.Context, req request.ReferralListRequest) (*domain.Referrals, error) {
			return serviceRegistry.ReferralService.FetchReferrals(ctx, req.ToReferralListParams())
		},
		response.NewReferralListResponse,
	)
}
