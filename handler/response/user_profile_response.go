package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type UserProfileResponse struct {
	composers.UserProfile
}

func NewUserProfileResponse(_ context.Context, user domain.User) (UserProfileResponse, HTTPMetadata) {
	profile := composers.NewUserProfile(user)
	return UserProfileResponse{
		UserProfile: profile,
	}, HTTPMetadata{}
}
