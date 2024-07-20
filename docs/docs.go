package docs

import (
	"github.com/techx/portal/handler/request"
)

// swagger:parameters getUserList
type AdminUserListRequestParams struct {
	// in: query
	Status string `json:"status"`
	// in: query
	UserUUID string `json:"user_uuid"`
	// in: query
	PhoneNumber string `json:"phone_number"`
	// in: query
	Name string `json:"name"`
	// in: query
	Company string `json:"company"`
	// in: query
	Designation string `json:"designation"`
}

// swagger:parameters generateOTP
type GenerateOTPRequest struct {
	// in: body
	// required: true
	OTPRequest request.OTPRequest
}

// swagger:parameters verifyOTP
type VerifyOTPRequest struct {
	// in: body
	// required: true
	OTPRequest request.OTPRequest
}

// swagger:parameters registerUserV1
type RegisterUserV1Request struct {
	// in: body
	// required: true
	RegisterUserV1Request request.RegisterUserV1Request
}

// swagger:parameters userProfile
type UserProfileRequest struct {
	// in: header
	// required: true
	UserUUID string `json:"X-User-Id"`
}

// swagger:parameters companyUsersList
type CompanyUsersListRequest struct {
	// in: header
	// required: true
	UserUUID string `json:"X-User-Id"`
}

// swagger:parameters referralRequest
type ReferralRequest struct {
	// in: header
	// required: true
	UserUUID string `json:"X-User-Id"`
	// in: body
	// required: true
	Body request.ReferralRequest
}
