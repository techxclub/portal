package docs

import (
	"github.com/techx/portal/handler/request"
)

// swagger:parameters registerUserV1
type RegisterRequestHeader struct {
	// in: header
	RequestTraceID string `json:"X-Request-Trace-Id"`
	// in: header
	UserType string `json:"X-User-Type"`
}

// swagger:parameters registerUserV1
type RegisterUserV1Request struct {
	// in: body
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
	FirstName string `json:"first_name"`
	// in: query
	LastName string `json:"last_name"`
	// in: query
	Company string `json:"company"`
	// in: query
	Role string `json:"role"`
}
