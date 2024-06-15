package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type BaseUserDetailsRequest struct {
	UserID      string
	Status      string
	Name        string
	PhoneNumber string
	Company     string
	Role        string
}

type UserProfileRequest struct {
	BaseUserDetailsRequest
}

func NewUserProfileRequest(r *http.Request) (*UserProfileRequest, error) {
	userID := r.Header.Get(constants.HeaderXUserID)
	return &UserProfileRequest{
		BaseUserDetailsRequest{
			UserID: userID,
		},
	}, nil
}

func (r UserProfileRequest) Validate() error {
	return nil
}

func (r UserProfileRequest) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserID: r.BaseUserDetailsRequest.UserID,
	}
}

type BulkUserDetailsRequest struct {
	BaseUserDetailsRequest
}

func NewBulkUserDetailsRequest(r *http.Request) (*BulkUserDetailsRequest, error) {
	userID := r.URL.Query().Get(constants.ParamUserID)
	status := r.URL.Query().Get(constants.ParamStatus)
	name := r.URL.Query().Get(constants.ParamName)
	phoneNumber := r.URL.Query().Get(constants.ParamPhoneNumber)
	Company := r.URL.Query().Get(constants.ParamCompany)
	Role := r.URL.Query().Get(constants.ParamRole)

	return &BulkUserDetailsRequest{
		BaseUserDetailsRequest{
			UserID:      userID,
			Status:      status,
			Name:        name,
			PhoneNumber: utils.SanitizePhoneNumber(phoneNumber),
			Company:     Company,
			Role:        Role,
		},
	}, nil
}

func (r BulkUserDetailsRequest) Validate() error {
	return nil
}

func (r BulkUserDetailsRequest) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserID:      r.UserID,
		Status:      r.Status,
		PhoneNumber: r.PhoneNumber,
		Name:        r.Name,
		Company:     r.Company,
		Role:        r.Role,
	}
}
