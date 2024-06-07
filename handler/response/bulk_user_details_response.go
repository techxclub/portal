package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type BulkUserDetailsResponse struct {
	Users []UserProfile `json:"users"`
}

func NewBulkUserDetailsResponse(_ context.Context, _ config.Config, user []domain.UserProfile) (BulkUserDetailsResponse, http.Header) {
	users := make([]UserProfile, 0, len(user))
	for _, u := range user {
		users = append(users, getUserProfile(u))
	}

	return BulkUserDetailsResponse{
		Users: users,
	}, http.Header{}
}
