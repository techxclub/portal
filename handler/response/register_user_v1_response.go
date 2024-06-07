package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type RegisterUserV1Response struct {
	Success bool   `json:"success"`
	UserID  string `json:"user_id"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, user domain.UserProfile) (RegisterUserV1Response, HTTPMetadata) {
	respBody := RegisterUserV1Response{
		Success: true,
		UserID:  user.UserID,
	}

	return respBody, HTTPMetadata{}
}
