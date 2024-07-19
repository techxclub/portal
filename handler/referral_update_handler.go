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

func ReferralUpdateHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewReferralUpdateRequest,
		func(ctx context.Context, req request.ReferralUpdateRequest) (*domain.EmptyDomain, error) {
			return serviceRegistry.AdminService.UpdateReferralDetails(ctx, req.ToReferralUpdateParams())
		},
		func(ctx context.Context, _ domain.EmptyDomain) (response.SuccessResponse, response.HTTPMetadata) {
			return response.NewSuccessResponse(ctx)
		},
	)
}
