package admin

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/handler/admin/request"
	"github.com/techx/portal/handler/admin/response"
	"github.com/techx/portal/service"
)

func CompanyListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return handler.Handler(
		request.NewAdminCompanyListRequest,
		func(ctx context.Context, req request.AdminCompanyListRequest) (*domain.Companies, error) {
			return serviceRegistry.UserService.GetCompaniesAdmin(ctx, req.ToFetchCompanyParams())
		},
		response.NewAdminCompanyListResponse,
	)
}
