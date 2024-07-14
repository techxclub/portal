package response

import (
	"context"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)


// swagger:model
type GmailAuthResponse struct {
	Action string `json:"action"`
	Profile *composers.UserProfile `json:"profile,omitempty"`
}

func NewGmailAuthResponse(_ context.Context, authDetails domain.AuthDetails) (GmailAuthResponse, HTTPMetadata) {
	return GmailAuthResponse{Action: constants.ActionLogInWithGoogle}, HTTPMetadata{}
}
