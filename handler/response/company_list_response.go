package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type CompanyListResponse struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewCompanyListResponse(_ context.Context, _ config.Config, companies domain.Companies) (CompanyListResponse, HTTPMetadata) {
	companyList := make([]Company, 0)
	for _, c := range companies {
		companyList = append(companyList, Company{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return CompanyListResponse{
		Companies: companyList,
	}, HTTPMetadata{}
}
