package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

// swagger:model
type RegisterUserV1Response struct {
	Success bool        `json:"success"`
	Profile UserProfile `json:"profile"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, registration domain.Registration) (RegisterUserV1Response, HTTPMetadata) {
	respBody := RegisterUserV1Response{
		Success: true,
		Profile: getUserProfile(*registration.User),
	}

	return respBody, HTTPMetadata{
		Cookies: &http.Cookie{
			Name:     constants.CookieAuthToken,
			Value:    registration.AuthToken,
			SameSite: http.SameSiteStrictMode,
		},
	}
}
