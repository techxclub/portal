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

func CompanyListHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewCompanyListRequest,
		func(ctx context.Context, _ request.CompanyListRequest) (*domain.Companies, error) {
			return serviceRegistry.UserService.GetCompanies(ctx)
		},
		func(ctx context.Context, domainObj domain.Companies) (response.CompanyListResponse, response.HTTPMetadata) {
			return response.NewCompanyListResponse(ctx, cfg, domainObj)
		},
	)
}
