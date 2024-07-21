package errors

import "errors"

var (
	New    = errors.New
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

var (
	ErrZeroRowsAffected = errors.New("no rows affected")

	ErrGeneratingAuthToken    = errors.New("error generating auth token")
	ErrInvalidUserID          = NewServiceError("invalid_user_id", 400, nil)
	ErrInvalidUserUpdate      = NewServiceError("invalid_user_update", 400, nil)
	ErrWorkEmailNotVerified   = NewServiceError("work_email_not_verified", 400, nil)
	ErrEmptyName              = NewServiceError("empty_name", 400, nil)
	ErrEmptyPhoneNumber       = NewServiceError("empty_phone_number", 400, nil)
	ErrEmptyLinkedIn          = NewServiceError("empty_linkedin", 400, nil)
	ErrEmptyCompanyName       = NewServiceError("empty_company_name", 400, nil)
	ErrEmptyWorkEmail         = NewServiceError("empty_work_email", 400, nil)
	ErrEmptyDesignation       = NewServiceError("empty_designation", 400, nil)
	ErrEmptyYearsOfExperience = NewServiceError("empty_years_of_experience", 400, nil)

	ErrInvalidPhoneNumber               = NewServiceError("invalid_phone_number", 400, nil)
	ErrInvalidYearsOfExperience         = NewServiceError("invalid_years_of_experience", 400, nil)
	ErrInvalidWorkEmail                 = NewServiceError("invalid_work_email", 400, nil)
	ErrInvalidRegisteredEmail           = NewServiceError("invalid_registered_email", 400, nil)
	ErrInvalidAuthChannel               = NewServiceError("invalid_auth_channel", 400, nil)
	ErrInvalidEmailProvider             = NewServiceError("invalid_email_service_provider", 500, nil)
	ErrInvalidUpdateRequest             = NewServiceError("invalid_update_request", 400, nil)
	ErrMissingOTP                       = NewServiceError("otp_missing", 500, nil)
	ErrOTPGenerateFailed                = NewServiceError("otp_generation_failed", 500, nil)
	ErrCompanyNotMatch                  = NewServiceError("company_not_matched", 400, nil)
	ErrRequesterNotFound                = NewServiceError("requester_not_found", 404, nil)
	ErrProviderNotFound                 = NewServiceError("provider_not_found", 404, nil)
	ErrInvalidJobLink                   = NewServiceError("invalid_job_link", 400, nil)
	ErrInvalidQueryParams               = NewServiceError("invalid_query_params", 400, nil)
	ErrNameRequired                     = NewServiceError("name_required", 400, nil)
	ErrCompanyRequired                  = NewServiceError("company_required", 400, nil)
	ErrInvalidCompanyID                 = NewServiceError("invalid_company_id", 400, nil)
	ErrRequesterFieldIsEmpty            = NewServiceError("requester_field_is_empty", 400, nil)
	ErrProviderFieldIsEmpty             = NewServiceError("provider_field_is_empty", 400, nil)
	ErrReferralLimitReachedForRequester = NewServiceError("referral_limit_reached_for_requester", 400, nil)
	ErrReferralLimitReachedForProvider  = NewServiceError("referral_limit_reached_for_provider", 400, nil)
	ErrReferralAlreadyExists            = NewServiceError("referral_already_exists", 400, nil)
	ErrUserNotFound                     = NewServiceError("user_not_found", 404, nil)
	ErrReferralNotFound                 = NewServiceError("referral_not_found", 404, nil)
	ErrSearchParamRequired              = NewServiceError("search_param_required", 400, nil)
	ErrNoDataFound                      = NewServiceError("no_data_found", 404, nil)
	ErrCalendalyLinkRequired            = NewServiceError("calendaly_link_required", 400, nil)
	ErrTagsRequired                     = NewServiceError("tags_required", 400, nil)
	ErrTagsLimitExceededByFive          = NewServiceError("tags_limit_exceeded_by_five", 400, nil)
	ErrDomainRequired                   = NewServiceError("domain_required", 400, nil)
	ErrUserNotApproved                  = NewServiceError("user_not_approved", 400, nil)
	ErrUserAlreadyMentor                = NewServiceError("user_already_mentor", 400, nil)
	ErrRequesterIDRequired              = NewServiceError("requester_id_required", 400, nil)
	ErrMissingAuthCode                  = NewServiceError("missing_auth_code", 400, nil)
	ErrInvalidAuthState                 = NewServiceError("invalid_auth_state", 400, nil)
	ErrUnverifiedEmail                  = NewServiceError("unverified_email", 400, nil)
)

func BadRequestError(err error) ServiceError {
	return &serviceError{
		err:                err,
		code:               "bad_request",
		i18nKey:            "bad_request",
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
		i18nKey:            "unhandled_error",
		responseStatusCode: 500,
	}
}
