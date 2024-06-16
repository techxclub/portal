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

func AdminCompanyListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewAdminCompanyListRequest,
		func(ctx context.Context, req request.AdminCompanyListRequest) (*domain.Companies, error) {
			return serviceRegistry.UserService.GetCompanies(ctx, req.ToFetchCompanyParams())
		},
		response.NewAdminCompanyListResponse,
	)
}
