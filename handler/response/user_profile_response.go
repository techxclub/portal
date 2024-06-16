package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type UserProfileResponse struct {
	Profile composers.UserProfile `json:"profile"`
}

func NewUserProfileResponse(_ context.Context, user domain.UserProfile) (UserProfileResponse, HTTPMetadata) {
	profile := composers.NewUserProfile(user)
	return UserProfileResponse{
		Profile: profile,
	}, HTTPMetadata{}
}
