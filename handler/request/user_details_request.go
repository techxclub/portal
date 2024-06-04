package request

import (
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UserDetailsRequest struct {
	UserID      string `json:"-"`
	PhoneNumber string `json:"-"`
}

func NewUserDetailsRequest(r *http.Request) (*UserDetailsRequest, error) {
	userID := r.URL.Query().Get("user_id")
	phoneNumber := r.URL.Query().Get("phone_number")

	return &UserDetailsRequest{
		UserID:      userID,
		PhoneNumber: phoneNumber,
	}, nil
}

func (g UserDetailsRequest) Validate() error {
	if g.UserID == "" && g.PhoneNumber == "" {
		return errors.New("invalid user id or phone number")
	}

	return nil
}

func (g UserDetailsRequest) ToDomainObject() domain.UserDetailsRequest {
	return domain.UserDetailsRequest{
		UserID: g.UserID,
		Phone:  g.PhoneNumber,
	}
}
