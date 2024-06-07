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

func GenerateOTPHandler(cfg config.Config, serviceRegistry *service.Registry, _ string) http.HandlerFunc {
	return phoneOTPHandler(cfg, serviceRegistry)
}

func phoneOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewPhoneOTPRequest,
		func(ctx context.Context, req request.PhoneOTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.GenerateOTP(ctx, req.ToOTPGenerationObject())
		},
		func(ctx context.Context, domainObj domain.AuthDetails) (response.GenerateOTPResponse, response.HTTPMetadata) {
			return response.NewGenerateOTPResponse(ctx, cfg, domainObj)
		},
	)
}
