package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type GenerateOTPResponse struct {
	Status string `json:"status"`
}

func NewGenerateOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (GenerateOTPResponse, HTTPMetadata) {
	return GenerateOTPResponse{
		Status: authDetails.Status,
	}, HTTPMetadata{}
}

// swagger:model
type VerifyOTPResponse struct {
	Status string `json:"status"`
}

func NewVerifyOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (VerifyOTPResponse, HTTPMetadata) {
	return VerifyOTPResponse{
		Status: authDetails.Status,
	}, HTTPMetadata{}
}
