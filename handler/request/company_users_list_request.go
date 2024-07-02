package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type CompanyUsersListRequest struct {
	CompanyID string
}

func NewCompanyUsersListRequest(r *http.Request) (*CompanyUsersListRequest, error) {
	companyID := r.URL.Query().Get(constants.ParamCompanyID)
	return &CompanyUsersListRequest{
		CompanyID: companyID,
	}, nil
}

func (r CompanyUsersListRequest) Validate() error {
	if r.CompanyID == "" {
		return errors.ErrCompanyRequired
	}

	if utils.ParseInt64WithDefault(r.CompanyID, 0) == 0 {
		return errors.ErrInvalidCompanyID
	}

	return nil
}

func (r CompanyUsersListRequest) ToFetchUserParams() domain.FetchUserParams {
	return domain.FetchUserParams{
		CompanyID: utils.ParseInt64WithDefault(r.CompanyID, 0),
		Status:    constants.StatusApproved,
	}
}
