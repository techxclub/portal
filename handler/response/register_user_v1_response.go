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
	UserID  string `json:"user_id"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, user domain.User) (RegisterUserV1Response, http.Header) {
	respBody := RegisterUserV1Response{
		Success: true,
		UserID:  user.UserID,
	}

	return respBody, http.Header{}
}
