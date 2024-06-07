package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type BaseUserDetailsRequest struct {
	UserID      string
	PhoneNumber string
	FirstName   string
	LastName    string
	Company     string
	Role        string
}

type UserProfileRequest struct {
	BaseUserDetailsRequest
}

func NewUserProfileRequest(r *http.Request) (*UserProfileRequest, error) {
	userID := r.URL.Query().Get(constants.ParamUserID)
	phoneNumber := r.URL.Query().Get(constants.ParamPhoneNumber)

	return &UserProfileRequest{
		BaseUserDetailsRequest{
			UserID:      userID,
			PhoneNumber: phoneNumber,
		},
	}, nil
}

func (r UserProfileRequest) Validate() error {
	return nil
}

func (r UserProfileRequest) ToDomainObject() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserID:      r.BaseUserDetailsRequest.UserID,
		PhoneNumber: r.BaseUserDetailsRequest.PhoneNumber,
	}
}

type BulkUserDetailsRequest struct {
	BaseUserDetailsRequest
}

func NewBulkUserDetailsRequest(r *http.Request) (*BulkUserDetailsRequest, error) {
	userID := r.URL.Query().Get(constants.ParamUserID)
	phoneNumber := r.URL.Query().Get(constants.ParamPhoneNumber)
	firstName := r.URL.Query().Get(constants.ParamFirstName)
	lastName := r.URL.Query().Get(constants.ParamLastName)
	Company := r.URL.Query().Get(constants.ParamCompany)
	Role := r.URL.Query().Get(constants.ParamRole)

	return &BulkUserDetailsRequest{
		BaseUserDetailsRequest{
			UserID:      userID,
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
			LastName:    lastName,
			Company:     Company,
			Role:        Role,
		},
	}, nil
}

func (g BulkUserDetailsRequest) Validate() error {
	return nil
}

func (g BulkUserDetailsRequest) ToDomainObject() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserID:      g.UserID,
		PhoneNumber: g.PhoneNumber,
		FirstName:   g.FirstName,
		LastName:    g.LastName,
		Company:     g.Company,
		Role:        g.Role,
	}
}
