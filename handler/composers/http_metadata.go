package composers

import "net/http"

type HTTPMetadata struct {
	Headers *http.Header
	Cookies *http.Cookie
}

func NewHTTPMetadata(header *http.Header, cookies *http.Cookie) HTTPMetadata {
	return HTTPMetadata{
		Headers: header,
		Cookies: cookies,
	}
}
