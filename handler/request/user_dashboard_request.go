package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/errors"
)

type UserDashboardRequest struct {
	UserUUID string
}

func NewUserDashboardRequest(r *http.Request) (*UserDashboardRequest, error) {
	requesterUserUUID := r.Header.Get(constants.HeaderXUserUUID)
	return &UserDashboardRequest{
		UserUUID: requesterUserUUID,
	}, nil
}

func (r UserDashboardRequest) Validate() error {
	if r.UserUUID == "" {
		return errors.ErrRequesterIDRequired
	}
	return nil
}
