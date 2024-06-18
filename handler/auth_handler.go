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

func GenerateOTPHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewGenerateOTPRequest,
		func(ctx context.Context, req request.OTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.GenerateOTP(ctx, req.ToAuthRequest())
		},
		response.NewGenerateOTPResponse,
	)
}

func VerifyOTPHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewVerifyOTPRequest,
		func(ctx context.Context, req request.OTPRequest) (*domain.AuthDetails, error) {
			return serviceRegistry.AuthService.VerifyOTP(ctx, req.ToAuthRequest())
		},
		response.NewVerifyOTPResponse,
	)
}
