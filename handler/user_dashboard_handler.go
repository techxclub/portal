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

func UserDashboardHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewUserDashboardRequest,
		func(ctx context.Context, req request.UserDashboardRequest) (*domain.UserReferrals, error) {
			return serviceRegistry.ReferralService.FetchReferralsForUser(ctx, req.UserUUID)
		},
		response.NewUserDashboardResponse,
	)
}
