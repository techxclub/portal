package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type AdminUserUpdateRequest struct {
	From AdminUserUpdateParams `json:"from"`
	To   AdminUserUpdateParams `json:"to"`
}

type AdminUserUpdateParams struct {
	UserNumber  int64  `json:"user_number"`
	UserUUID    string `json:"user_uuid"`
	Status      string `json:"status"`
	CompanyName string `json:"company_name"`
	Designation string `json:"designation"`
}

func NewAdminUserUpdateRequest(r *http.Request) (*AdminUserUpdateRequest, error) {
	var req AdminUserUpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r AdminUserUpdateRequest) Validate() error {
	if r.To.UserNumber != 0 && r.To.UserUUID != "" {
		return errors.ErrInvalidUpdateRequest
	}
	return nil
}

func (r AdminUserUpdateParams) ToUserProfile() domain.User {
	return domain.User{
		UserNumber: r.UserNumber,
		UserUUID:   r.UserUUID,
		Status:     r.Status,
		ProfessionalInformation: domain.ProfessionalInformation{
			CompanyName: r.CompanyName,
			Designation: r.Designation,
		},
	}
}
