package request

import (
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type CompanyListRequest struct {
	Verified bool
}

func NewCompanyListRequest(_ *http.Request) (*CompanyListRequest, error) {
	return &CompanyListRequest{Verified: true}, nil
}

func (r CompanyListRequest) Validate() error {
	return nil
}

func (r CompanyListRequest) ToFetchCompanyParams() domain.FetchCompanyParams {
	return domain.FetchCompanyParams{
		Verified: utils.ToPtr(r.Verified),
	}
}
