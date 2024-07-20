package constants

const (
	HeaderXUserUUID       = "X-User-Id"
	HeaderXForwardedFor   = "X-FORWARDED-FOR"
	HeaderXRequestTraceID = "X-Request-Trace-Id"
	HeaderContentType     = "Content-Type"
	HeaderAuthorization   = "Authorization"
	HeaderClientID        = "Client-ID"
	HeaderPasskey         = "Passkey"
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
	AllowedHeaders = []string{HeaderXUserUUID, HeaderXRequestTraceID, HeaderContentType, HeaderXForwardedFor, HeaderAuthorization, HeaderAuthToken}
	ExposedHeaders = []string{HeaderAuthToken, HeaderContentType}
	AllowedMethods = []string{MethodGet, MethodPost, MethodPut, MethodPatch, MethodHead, MethodOptions}
	AllowedOrigins = []string{"*"}
)
