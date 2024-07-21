package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UserFetchProfileRequest struct {
	BaseUserListRequest
}

func NewUserFetchProfileRequest(r *http.Request) (*UserFetchProfileRequest, error) {
	userID := r.Header.Get(constants.HeaderXUserUUID)
	return &UserFetchProfileRequest{
		BaseUserListRequest{
			UserUUID: userID,
		},
	}, nil
}

func (r UserFetchProfileRequest) Validate() error {
	if r.UserUUID == "" {
		return errors.ErrInvalidUserID
	}

	return nil
}

func (r UserFetchProfileRequest) ToFetchUserParams() domain.FetchUserParams {
	return domain.FetchUserParams{
		UserUUID: r.BaseUserListRequest.UserUUID,
	}
}
