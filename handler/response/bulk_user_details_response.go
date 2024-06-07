package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type BulkUserDetailsResponse struct {
	Users []UserProfile `json:"users"`
}

func NewBulkUserDetailsResponse(_ context.Context, _ config.Config, user []domain.UserProfile) (BulkUserDetailsResponse, HTTPMetadata) {
	users := make([]UserProfile, 0, len(user))
	for _, u := range user {
		users = append(users, getUserProfile(u))
	}

	return BulkUserDetailsResponse{
		Users: users,
	}, HTTPMetadata{}
}
