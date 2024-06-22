package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
)

type AdminUserUpdateRequest struct {
	From AdminUserUpdateParams `json:"from"`
	To   AdminUserUpdateParams `json:"to"`
}

type AdminUserUpdateParams struct {
	UserNumber int64  `json:"user_number"`
	UserID     string `json:"user_id"`
	Status     string `json:"status"`
	Company    string `json:"company"`
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
	return nil
}

func (r AdminUserUpdateParams) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserIDNum:   r.UserNumber,
		UserID:      r.UserID,
		Status:      r.Status,
		CompanyName: r.Company,
	}
}
