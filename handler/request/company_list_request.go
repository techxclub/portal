package request

import (
	"net/http"
)

type CompanyListRequest struct{}

func NewCompanyListRequest(_ *http.Request) (*CompanyListRequest, error) {
	return &CompanyListRequest{}, nil
}

func (r CompanyListRequest) Validate() error {
	return nil
}
