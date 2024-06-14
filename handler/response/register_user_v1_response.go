package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type RegisterUserV1Response struct {
	Action  string                `json:"action"`
	Profile composers.UserProfile `json:"profile"`
}

func NewRegisterUserV1Response(_ context.Context, _ config.Config, registration domain.Registration) (RegisterUserV1Response, HTTPMetadata) {
	profile := composers.NewUserProfile(*registration.User)
	respBody := RegisterUserV1Response{
		Profile: profile,
		Action:  constants.ActionLogIn,
	}

	if registration.User.Status == constants.StatusPendingApproval {
		respBody.Action = constants.ActionPendingApproval
	}

	return respBody, HTTPMetadata{
		Cookies: &http.Cookie{
			Name:     constants.CookieAuthToken,
			Value:    registration.AuthToken,
			SameSite: http.SameSiteStrictMode,
		},
	}
}
