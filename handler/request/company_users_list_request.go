package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type CompanyUsersListRequest struct {
	Company string `json:"company"`
}

func NewCompanyUsersListRequest(r *http.Request) (*CompanyUsersListRequest, error) {
	company := r.URL.Query().Get(constants.ParamCompany)
	return &CompanyUsersListRequest{Company: company}, nil
}

func (r CompanyUsersListRequest) Validate() error {
	if r.Company == "" {
		return errors.ErrCompanyRequired
	}

	return nil
}

func (r CompanyUsersListRequest) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		Company: r.Company,
	}
}
