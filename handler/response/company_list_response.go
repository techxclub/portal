package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type CompanyListResponse struct {
	PopularCompanies []composers.Company `json:"popular_companies"`
	AllCompanies     []composers.Company `json:"all_companies"`
}

func NewCompanyListResponse(_ context.Context, cfg *config.Config, companies domain.Companies) (CompanyListResponse, HTTPMetadata) {
	return CompanyListResponse{
		PopularCompanies: composers.GetPopularCompanies(companies, cfg.PopularCompanyListLimit),
		AllCompanies:     composers.GetAllCompanies(companies, cfg.CompanyListLimit),
	}, HTTPMetadata{}
}
