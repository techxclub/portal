package response

import (
	"cmp"
	"context"
	"slices"

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
	slices.SortStableFunc(companies, func(i, j domain.Company) int {
		return cmp.Compare(i.GetPriority(), j.GetPriority())
	})

	return CompanyListResponse{
		PopularCompanies: composers.GetPopularCompanies(companies, cfg.PopularCompanyListLimit),
		AllCompanies:     composers.GetAllCompanies(companies, cfg.CompanyListLimit),
	}, HTTPMetadata{}
}
