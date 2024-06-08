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
	return generatePhoneOTPHandler(cfg, serviceRegistry)
}

func generatePhoneOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewGeneratePhoneOTPRequest,
		func(ctx context.Context, req request.GeneratePhoneOTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.GenerateOTP(ctx, req.ToOTPGenerationObject())
		},
		func(ctx context.Context, domainObj domain.AuthDetails) (response.GenerateOTPResponse, response.HTTPMetadata) {
			return response.NewGenerateOTPResponse(ctx, cfg, domainObj)
		},
	)
}

func VerifyOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return verifyPhoneOTPHandler(cfg, serviceRegistry)
}

func verifyPhoneOTPHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewPhoneOTPRequest,
		func(ctx context.Context, req request.VerifyOTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.VerifyOTP(ctx, req.ToOTPVerificationObject())
		},
		func(ctx context.Context, domainObj domain.AuthDetails) (response.VerifyOTPResponse, response.HTTPMetadata) {
			return response.NewVerifyOTPResponse(ctx, cfg, domainObj)
		},
	)
}
