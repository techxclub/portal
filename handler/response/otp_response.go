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
type GenerateOTPResponse struct {
	Action string `json:"action"`
}

// swagger:model
type VerifyOTPResponse struct {
	Action  string                 `json:"action"`
	Profile *composers.UserProfile `json:"profile,omitempty"`
}

func NewGenerateOTPResponse(_ context.Context, _ config.Config, _ domain.AuthDetails) (GenerateOTPResponse, HTTPMetadata) {
	return GenerateOTPResponse{Action: constants.ActionVerifyOTP}, HTTPMetadata{}
}

func NewVerifyOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (VerifyOTPResponse, HTTPMetadata) {
	action := constants.ActionSignUp
	if authDetails.AuthInfo.Status == constants.AuthStatusPending {
		action = constants.ActionRetryOTP
	}

	if authDetails.UserInfo == nil {
		return VerifyOTPResponse{Action: action}, HTTPMetadata{}
	}

	profile := composers.NewUserProfile(*authDetails.UserInfo)
	if !authDetails.UserInfo.IsApproved() {
		return VerifyOTPResponse{
			Action:  constants.ActionPendingApproval,
			Profile: &profile,
		}, HTTPMetadata{}
	}

	verifyOTPResponse := VerifyOTPResponse{
		Action:  constants.ActionLogIn,
		Profile: &profile,
	}

	httpMetadata := HTTPMetadata{
		Cookies: &http.Cookie{
			Name:     constants.CookieAuthToken,
			Value:    authDetails.AuthToken,
			SameSite: http.SameSiteStrictMode,
		},
	}

	return verifyOTPResponse, httpMetadata
}
