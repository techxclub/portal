package errors

import (
	"fmt"
)

type ServiceError interface {
	error
	GetCode() string
	GetResponseStatus() int
	GetAnnotation() string
	UnWrap() error
}

type serviceError struct {
	err                error
	code               string
	responseStatusCode int
	annotation         string
}

func NewServiceError(code string, responseCode int, err error) ServiceError {
	return &serviceError{
		err:                err,
		code:               code,
		responseStatusCode: responseCode,
	}
}

func (e *serviceError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return fmt.Sprintf("ServiceError: %s", e.code)
}

func (e *serviceError) GetCode() string {
	return e.code
}

func (e *serviceError) GetResponseStatus() int {
	return e.responseStatusCode
}

func (e *serviceError) GetAnnotation() string {
	if e.annotation != "" {
		return e.annotation
	}

	return e.code
}

func (e *serviceError) UnWrap() error {
	return e.err
}
