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

func RegisterUserV1Handler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewRegisterUserV1Request,
		func(ctx context.Context, req request.RegisterUserV1Request) (*domain.Registration, error) {
			return serviceRegistry.UserService.RegisterUser(ctx, req.ToUserDetails())
		},
		response.NewRegisterUserV1Response,
	)
}
