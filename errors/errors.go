package errors

import "errors"

var (
	New    = errors.New
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

func BadRequestError(err error) ServiceError {
	return &serviceError{
		err:                err,
		code:               "bad_request",
		responseStatusCode: 400,
	}
}

func AsServiceError(err error) ServiceError {
	var e *serviceError
	if errors.As(err, &e) {
		return e
	}

	return &serviceError{
		err:                err,
		code:               "INTERNAL_SERVER_ERROR",
		responseStatusCode: 500,
	}
}
