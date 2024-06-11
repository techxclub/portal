package handler

import (
	"context"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/request"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
	"net/http"
)

func CompanyUsersListHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewCompanyUsersListRequest,
		func(ctx context.Context, req request.CompanyUsersListRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToUserProfileParams())
		},
		func(ctx context.Context, domainObj domain.Users) (response.CompanyUsersListResponse, response.HTTPMetadata) {
			return response.NewCompanyUsersListResponse(ctx, cfg, domainObj)
		},
	)
}
