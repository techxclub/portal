package errors

import "errors"

var (
	New    = errors.New
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

var (
	ErrInvalidPhoneNumberFormat = NewServiceError("invalid_phone_number_format", 400, nil)
	ErrInvalidPhoneNumber       = NewServiceError("invalid_phone_number", 400, nil)
	ErrInvalidYearsOfExperience = NewServiceError("invalid_years_of_experience", 400, nil)
	ErrInvalidWorkEmail         = NewServiceError("invalid_work_email", 400, nil)
	ErrInvalidPersonalEmail     = NewServiceError("invalid_personal_email", 400, nil)
	ErrInvalidAuthChannel       = NewServiceError("invalid_auth_channel", 400, nil)
	ErrTwilioCreateVerification = NewServiceError("twilio_create_verification_failed", 500, nil)
	ErrTwilioCheckVerification  = NewServiceError("twilio_check_verification_failed", 500, nil)
	ErrMissingOTP               = NewServiceError("otp_missing_otp", 500, nil)
	ErrOTPGenerateFailed        = NewServiceError("otp_generation_failed", 500, nil)
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
