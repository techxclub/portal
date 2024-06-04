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

func UserDetailsHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewUserDetailsRequest,
		func(ctx context.Context, req request.UserDetailsRequest) (*domain.User, error) {
			return serviceRegistry.UserService.UserDetails(ctx, req.ToDomainObject())
		},
		func(ctx context.Context, domainObj domain.User) (response.UserDetailsResponse, http.Header) {
			return response.NewUserDetailsResponse(ctx, cfg, domainObj)
		},
	)
}
