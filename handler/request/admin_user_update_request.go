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
	UserID      string `json:"user_id"`
	Status      string `json:"status"`
	CompanyName string `json:"company_name"`
	Role        string `json:"role"`
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
	if r.To.UserNumber != 0 && r.To.UserID != "" {
		return errors.ErrInvalidUpdateRequest
	}
	return nil
}

func (r AdminUserUpdateParams) ToUserProfile() domain.UserProfile {
	return domain.UserProfile{
		UserIDNum:   r.UserNumber,
		UserID:      r.UserID,
		Status:      r.Status,
		CompanyName: r.CompanyName,
		Role:        r.Role,
	}
}
