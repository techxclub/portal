package composers

import (
	"context"
)

// swagger:model
type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewSuccessResponse(_ context.Context) (SuccessResponse, HTTPMetadata) {
	respBody := SuccessResponse{
		Success: true,
	}
	return respBody, HTTPMetadata{}
}
