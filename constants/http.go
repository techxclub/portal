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

	ApplicationJSON = "application/json; charset=utf-8"
)

var (
	AllowedHeaders = []string{HeaderXUserID, HeaderXRequestTraceID, HeaderContentType, HeaderXForwardedFor, HeaderAuthorization, HeaderAuthToken}
	ExposedHeaders = []string{HeaderAuthToken, HeaderContentType}
	AllowedMethods = []string{MethodGet, MethodPost, MethodPut, MethodPatch, MethodHead, MethodOptions}
	AllowedOrigins = []string{"*"}
)
