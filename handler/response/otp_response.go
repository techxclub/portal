package response

import (
	"context"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

var authStatusToActionMap = map[string]string{
	constants.OTPStatusGenerated: constants.ActionVerifyOTP,
	constants.OTPStatusPending:   constants.ActionVerifyOTP,
	constants.OTPStatusFailed:    constants.ActionRetryOTP,
	constants.OTPStatusVerified:  constants.ActionSuccess,
}

// swagger:model
type GenerateOTPResponse struct {
	Action string `json:"action"`
}

// swagger:model
type VerifyOTPResponse struct {
	Action string `json:"action"`
}

func NewGenerateOTPResponse(_ context.Context, _ domain.AuthDetails) (GenerateOTPResponse, composers.HTTPMetadata) {
	return GenerateOTPResponse{Action: constants.ActionVerifyOTP}, composers.HTTPMetadata{}
}

func NewVerifyOTPResponse(_ context.Context, authDetails domain.AuthDetails) (VerifyOTPResponse, composers.HTTPMetadata) {
	action, ok := authStatusToActionMap[authDetails.Status]
	if !ok {
		return VerifyOTPResponse{Action: constants.ActionRetryOTP}, composers.HTTPMetadata{}
	}

	return VerifyOTPResponse{Action: action}, composers.HTTPMetadata{}
}
