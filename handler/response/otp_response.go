package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type OTPResponse struct {
	Action  string                 `json:"action,omitempty"`
	Profile *composers.UserProfile `json:"profile,omitempty"`
}

func NewGenerateOTPResponse(_ context.Context, _ config.Config, _ domain.AuthDetails) (OTPResponse, HTTPMetadata) {
	return OTPResponse{
		Action: constants.ActionVerifyOTP,
	}, HTTPMetadata{}
}

func NewVerifyOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (OTPResponse, HTTPMetadata) {
	verifyOTPResponse := OTPResponse{
		Action: constants.ActionSignUp,
	}

	if authDetails.AuthInfo.Status == constants.AuthStatusPending {
		verifyOTPResponse.Action = constants.ActionRetryOTP
	}

	if authDetails.UserInfo != nil {
		profile := composers.NewUserProfile(*authDetails.UserInfo)
		verifyOTPResponse.Profile = &profile
		verifyOTPResponse.Action = constants.ActionLogIn
	}

	return verifyOTPResponse, HTTPMetadata{}
}
