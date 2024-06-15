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

func GenerateOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewGenerateOTPRequest,
		func(ctx context.Context, req request.OTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.GenerateOTP(ctx, req.ToAuthRequest())
		},
		func(ctx context.Context, domainObj domain.AuthDetails) (response.GenerateOTPResponse, response.HTTPMetadata) {
			return response.NewGenerateOTPResponse(ctx, cfg, domainObj)
		},
	)
}

func VerifyOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewVerifyOTPRequest,
		func(ctx context.Context, req request.OTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.VerifyUser(ctx, req.ToAuthRequest())
		},
		func(ctx context.Context, domainObj domain.AuthDetails) (response.VerifyOTPResponse, response.HTTPMetadata) {
			return response.NewVerifyOTPResponse(ctx, cfg, domainObj)
		},
	)
}
