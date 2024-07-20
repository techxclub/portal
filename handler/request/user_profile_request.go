package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type UserProfileRequest struct {
	BaseUserListRequest
}

func NewUserProfileRequest(r *http.Request) (*UserProfileRequest, error) {
	userID := r.Header.Get(constants.HeaderXUserUUID)
	return &UserProfileRequest{
		BaseUserListRequest{
			UserUUID: userID,
		},
	}, nil
}

func (r UserProfileRequest) Validate() error {
	return nil
}

func (r UserProfileRequest) ToFetchUserParams() domain.FetchUserParams {
	return domain.FetchUserParams{
		UserUUID: r.BaseUserListRequest.UserUUID,
	}
}
