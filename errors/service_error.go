package errors

import (
	"fmt"
)

type ServiceError interface {
	error
	GetCode() string
	GetI18nKey() string
	GetI18nValues() map[string]interface{}
	GetResponseStatus() int
	UnWrap() error
}

type serviceError struct {
	err                error
	code               string
	i18nKey            string
	i18nValues         map[string]interface{}
	responseStatusCode int
}

func NewServiceError(code string, responseCode int, err error) ServiceError {
	return &serviceError{
		err:                err,
		code:               code,
		i18nKey:            code,
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

func (e *serviceError) GetI18nKey() string {
	return "err_" + e.i18nKey
}

func (e *serviceError) GetI18nValues() map[string]interface{} {
	return e.i18nValues
}

func (e *serviceError) GetResponseStatus() int {
	return e.responseStatusCode
}

func (e *serviceError) UnWrap() error {
	return e.err
}
