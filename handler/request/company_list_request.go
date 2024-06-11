package request

import (
	"net/http"
)

type CompaniesListRequest struct{}

func NewCompaniesListRequest(_ *http.Request) (*CompaniesListRequest, error) {
	return &CompaniesListRequest{}, nil
}

func (r CompaniesListRequest) Validate() error {
	return nil
}
