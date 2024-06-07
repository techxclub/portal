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

func RegisterUserV1Handler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewRegisterUserV1Request,
		func(ctx context.Context, req request.RegisterUserV1Request) (*domain.UserProfile, error) {
			return serviceRegistry.UserService.RegisterUser(ctx, req.ToUserDetails())
		},
		func(ctx context.Context, domainObj domain.UserProfile) (response.RegisterUserV1Response, response.HTTPMetadata) {
			return response.NewRegisterUserV1Response(ctx, cfg, domainObj)
		},
	)
}
