package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type UserListResponse struct {
	Users []composers.UserProfile `json:"users"`
}

func NewUserListResponse(_ context.Context, user []domain.User) (UserListResponse, HTTPMetadata) {
	users := make([]composers.UserProfile, 0, len(user))
	for _, u := range user {
		users = append(users, composers.NewUserProfile(u))
	}

	return UserListResponse{
		Users: users,
	}, HTTPMetadata{}
}
