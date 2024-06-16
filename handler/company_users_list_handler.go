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
		func(ctx context.Context, req request.CompanyUsersListRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToUserProfileParams())
		},
		response.NewCompanyUsersListResponse,
	)
}
