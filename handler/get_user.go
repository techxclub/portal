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

func GetUserByIDHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewGetUserByIDRequest,
		func(ctx context.Context, req request.GetUserByIDRequest) (*domain.User, error) {
			return serviceRegistry.UserService.GetUserByID(ctx, req.UserID)
		},
		func(ctx context.Context, domainObj domain.User) (response.UserDetailsResponse, http.Header) {
			return response.NewGetUserResponse(ctx, cfg, domainObj)
		},
	)
}
