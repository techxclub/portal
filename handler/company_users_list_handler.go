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

func CompanyUsersListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewCompanyUsersListRequest,
		func(ctx context.Context, req request.CompanyUsersListRequest) (*domain.CompanyUsersService, error) {
			return serviceRegistry.UserService.GetCompanyUsers(ctx, req.ToFetchUserParams())
		},
		response.NewCompanyUsersListResponse,
	)
}
