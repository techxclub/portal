package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

// swagger:model
type OTPResponse struct {
	Status string `json:"status"`
	Action string `json:"action,omitempty"`
	UserID string `json:"user_id,omitempty"`
}

func NewGenerateOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (OTPResponse, HTTPMetadata) {
	return OTPResponse{
		Status: authDetails.AuthInfo.Status,
	}, HTTPMetadata{}
}

func NewVerifyOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (OTPResponse, HTTPMetadata) {
	verifyOTPResponse := OTPResponse{
		Status: authDetails.AuthInfo.Status,
		Action: constants.ActionSignUp,
	}

	if authDetails.AuthInfo.Status != constants.AuthStatusApproved {
		verifyOTPResponse.Action = constants.ActionRetry
	}

	if authDetails.UserInfo != nil {
		verifyOTPResponse.UserID = authDetails.UserInfo.UserID
		verifyOTPResponse.Action = constants.ActionLogIn
	}

	return verifyOTPResponse, HTTPMetadata{}
}
