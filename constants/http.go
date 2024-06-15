package constants

const (
	HeaderXUserID         = "X-User-Id"
	HeaderContentType     = "Content-Type"
	HeaderXForwardedFor   = "X-FORWARDED-FOR"
	HeaderXRequestTraceID = "X-Request-Trace-Id"
	HeaderAuthorization   = "Authorization"
	HeaderAuthToken       = "Auth-Token"

	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
	MethodDelete  = "DELETE"

	ApplicationJSON = "application/json; charset=utf-8"
)

var (
	AllowedHeaders = []string{HeaderXUserID, HeaderXRequestTraceID, HeaderContentType, HeaderXForwardedFor, HeaderAuthorization}
	AllowedMethods = []string{MethodGet, MethodPost, MethodPut, MethodPatch, MethodHead, MethodOptions}
	AllowedOrigins = []string{"*"}
)
