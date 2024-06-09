package errors

import (
	"fmt"
	"net/http"
)

func IsUpstreamError(err error) bool {
	var upstreamError *upstreamError
	ok := As(err, &upstreamError)
	return ok
}

type upstreamError struct {
	Err         error
	CommandName string
	StatusCode  int
}

type UpstreamError interface {
	error
	GetStatusCode() int
}

func NewUpstreamError(err error, commandName string, statusCode int) UpstreamError {
	return &upstreamError{
		Err:         err,
		CommandName: commandName,
		StatusCode:  statusCode,
	}
}

func (e *upstreamError) GetStatusCode() int {
	return e.StatusCode
}

func (e *upstreamError) Error() string {
	errString := fmt.Sprintf("Client: %s", e.CommandName)
	if e.StatusCode != 0 {
		errString += ", StatusCode: " + http.StatusText(e.StatusCode)
	}
	errString += ", Error: " + e.Err.Error()
	return errString
}
