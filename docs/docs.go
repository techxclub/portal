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
type RegisterUserV1RequestParams struct {
	// in: body
	RegisterUserV1Request request.RegisterUserV1Request
}

// swagger:parameters getUserDetails
type GetUserByIDRequestParams struct {
	// in: path
	UserID int64 `json:"userID"`
}
