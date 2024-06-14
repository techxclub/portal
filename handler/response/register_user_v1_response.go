package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type RegisterUserV1Response struct {
	Success bool        `json:"success"`
	Profile UserProfile `json:"profile"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, user domain.UserProfile) (RegisterUserV1Response, HTTPMetadata) {
	respBody := RegisterUserV1Response{
		Success: true,
		Profile: getUserProfile(user),
	}

	return respBody, HTTPMetadata{}
}
