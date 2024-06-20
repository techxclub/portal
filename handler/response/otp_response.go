package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

var authStatusToActionMap = map[string]string{
	constants.AuthStatusGenerated: constants.ActionVerifyOTP,
	constants.AuthStatusPending:   constants.ActionVerifyOTP,
	constants.AuthStatusResent:    constants.ActionVerifyOTP,
	constants.AuthStatusFailed:    constants.ActionRetryOTP,
	constants.AuthStatusVerified:  constants.ActionSignUp,
}

// swagger:model
type GenerateOTPResponse struct {
	Action string `json:"action"`
}

// swagger:model
type VerifyOTPResponse struct {
	Action  string                 `json:"action"`
	Profile *composers.UserProfile `json:"profile,omitempty"`
}

func NewGenerateOTPResponse(_ context.Context, _ domain.AuthDetails) (GenerateOTPResponse, HTTPMetadata) {
	return GenerateOTPResponse{Action: constants.ActionVerifyOTP}, HTTPMetadata{}
}

func NewVerifyOTPResponse(_ context.Context, authDetails domain.AuthDetails) (VerifyOTPResponse, HTTPMetadata) {
	action, ok := authStatusToActionMap[authDetails.AuthInfo.Status]
	if !ok {
		return VerifyOTPResponse{Action: constants.ActionRetryOTP}, HTTPMetadata{}
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
		Headers: &http.Header{
			constants.HeaderAuthToken: []string{authDetails.AuthToken},
		},
	}

	return verifyOTPResponse, httpMetadata
}
