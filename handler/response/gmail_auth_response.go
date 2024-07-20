package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type GoogleOAuthLoginResponse struct {
	Profile composers.UserProfile `json:"profile"`
}

func NewGoogleOAuthLoginResponse(_ context.Context, authDetails domain.AuthDetails) (GoogleOAuthLoginResponse, HTTPMetadata) {
	httpMetadata := HTTPMetadata{
		Headers: &http.Header{
			constants.HeaderAuthToken: []string{authDetails.Token},
		},
	}
	loginResponse := GoogleOAuthLoginResponse{
		Profile: composers.NewUserProfile(*authDetails.UserInfo),
	}

	return loginResponse, httpMetadata
}
