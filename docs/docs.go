package docs

import (
	"github.com/techx/portal/handler/request"
)

// swagger:parameters registerUserV1
type RegisterRequestHeader struct {
	// in: header
	RequestTraceID string `json:"X-Request-Trace-Id"`
}

// swagger:parameters registerUserV1
type RegisterUserV1Request struct {
	// in: body
	// required: true
	RegisterUserV1Request request.RegisterUserV1Request
}

// swagger:parameters userProfile
type UserProfileRequestParams struct {
	// in: query
	UserID string `json:"user_id"`
	// in: query
	PhoneNumber string `json:"phone_number"`
}

// swagger:parameters bulkGetUsers
type BulkGetUsersRequestParams struct {
	// in: query
	UserID string `json:"user_id"`
	// in: query
	PhoneNumber string `json:"phone_number"`
	// in: query
	Name string `json:"name"`
	// in: query
	Company string `json:"company"`
	// in: query
	Role string `json:"role"`
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
