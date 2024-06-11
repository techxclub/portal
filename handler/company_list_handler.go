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

func CompanyListHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewCompaniesListRequest,
		func(ctx context.Context, _ request.CompaniesListRequest) (*domain.Companies, error) {
			return serviceRegistry.UserService.GetCompanies(ctx)
		},
		func(ctx context.Context, domainObj domain.Companies) (response.AllCompaniesListResponse, response.HTTPMetadata) {
			return response.NewAllCompaniesListResponse(ctx, cfg, domainObj)
		},
	)
}
