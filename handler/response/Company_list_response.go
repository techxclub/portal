package response

import (
	"context"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type AllCompaniesListResponse struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	CompanyID int64  `json:"company_id"`
	Name      string `json:"name"`
}

func NewAllCompaniesListResponse(_ context.Context, _ config.Config, companies domain.Companies) (AllCompaniesListResponse, HTTPMetadata) {
	result := make([]Company, 0)
	for _, u := range companies {
		result = append(result, Company{
			CompanyID: u.CompanyID,
			Name:      u.Name,
		})
	}

	return AllCompaniesListResponse{
		Companies: result,
	}, HTTPMetadata{}
}
