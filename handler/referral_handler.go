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

func ReferralHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewReferralRequest,
		func(ctx context.Context, req request.ReferralRequest) (*domain.Referral, error) {
			return serviceRegistry.ReferralService.CreateReferral(ctx, req.ToReferral())
		},
		func(ctx context.Context, _ domain.Referral) (response.SuccessResponse, response.HTTPMetadata) {
			return response.NewSuccessResponse(ctx)
		},
	)
}
