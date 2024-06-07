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

func AdminUserDetailsHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewBulkUserDetailsRequest,
		func(ctx context.Context, req request.BulkUserDetailsRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToDomainObject())
		},
		func(ctx context.Context, domainObj domain.Users) (response.BulkUserDetailsResponse, http.Header) {
			return response.NewBulkUserDetailsResponse(ctx, cfg, domainObj)
		},
	)
}
