package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type CompanyUsersListRequest struct {
	CompanyName string `json:"company"`
}

func NewCompanyUsersListRequest(r *http.Request) (*CompanyUsersListRequest, error) {
	var req CompanyUsersListRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r CompanyUsersListRequest) Validate() error {
	if r.CompanyName == "" {
		return errors.New("Company name is required")
	}

	return nil
}

func (r CompanyUsersListRequest) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		Company: r.CompanyName,
	}
}
