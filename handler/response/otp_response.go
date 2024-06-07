package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type GenerateOTPResponse struct {
	Success bool `json:"success"`
}

func NewGenerateOTPResponse(_ context.Context, _ config.Config, authDetails domain.AuthDetails) (GenerateOTPResponse, HTTPMetadata) {
	return GenerateOTPResponse{
		Success: authDetails.Success,
	}, HTTPMetadata{}
}
