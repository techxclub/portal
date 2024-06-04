package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type RegisterUserV1Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, user domain.User) (RegisterUserV1Response, http.Header) {
	respBody := RegisterUserV1Response{
		Success: true,
		UserID:  user.UserID,
		Name:    user.FirstName + " " + user.LastName,
	}

	return respBody, http.Header{}
}
