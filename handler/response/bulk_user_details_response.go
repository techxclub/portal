package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type BulkUserDetailsResponse struct {
	Users []composers.UserProfile `json:"users"`
}

func NewBulkUserDetailsResponse(_ context.Context, _ config.Config, user []domain.UserProfile) (BulkUserDetailsResponse, HTTPMetadata) {
	users := make([]composers.UserProfile, 0, len(user))
	for _, u := range user {
		users = append(users, composers.NewUserProfile(u))
	}

	return BulkUserDetailsResponse{
		Users: users,
	}, HTTPMetadata{}
}
