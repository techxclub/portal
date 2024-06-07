package response

import "net/http"

type HTTPMetadata struct {
	Headers *http.Header
	Cookies *http.Cookie
}
