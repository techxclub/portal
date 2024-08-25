package response

import (
	"context"
	"math"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type AdminCompanyListResponse struct {
	Companies []composers.Company `json:"companies"`
}

func NewAdminCompanyListResponse(_ context.Context, companies domain.Companies) (AdminCompanyListResponse, composers.HTTPMetadata) {
	return AdminCompanyListResponse{Companies: composers.GetAllCompanies(companies, math.MaxInt16)}, composers.HTTPMetadata{}
}
