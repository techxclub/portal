package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type CompanyListResponse struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	CompanyID int64  `json:"company_id"`
	Name      string `json:"name"`
}

func NewCompanyListResponse(_ context.Context, _ config.Config, companies domain.Companies) (CompanyListResponse, HTTPMetadata) {
	companyList := make([]Company, 0)
	for _, u := range companies {
		companyList = append(companyList, Company{
			CompanyID: u.CompanyID,
			Name:      u.Name,
		})
	}

	return CompanyListResponse{
		Companies: companyList,
	}, HTTPMetadata{}
}
