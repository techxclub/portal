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
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, _ domain.User) (RegisterUserV1Response, http.Header) {
	respBody := RegisterUserV1Response{
		Success: true,
	}

	return respBody, http.Header{}
}
