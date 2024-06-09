package errors

import (
	"fmt"
	"net/http"
)

type HystrixError interface {
	error
	GetStatusCode() int
}

type hystrixError struct {
	httpCode int
	cmdName  string
}

func NewHystrixError(cmdName string, httpCode int) HystrixError {
	return &hystrixError{
		httpCode: httpCode,
		cmdName:  cmdName,
	}
}

func (h hystrixError) GetStatusCode() int {
	return h.httpCode
}

func (h hystrixError) Error() string {
	return fmt.Sprintf("hystrix: %s: %s", h.cmdName, http.StatusText(h.httpCode))
}
