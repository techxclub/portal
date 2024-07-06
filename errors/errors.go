package errors

import "errors"

var (
	New    = errors.New
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

var (
	ErrInvalidPhoneNumber               = NewServiceError("invalid_phone_number", 400, nil)
	ErrInvalidYearsOfExperience         = NewServiceError("invalid_years_of_experience", 400, nil)
	ErrInvalidWorkEmail                 = NewServiceError("invalid_work_email", 400, nil)
	ErrInvalidPersonalEmail             = NewServiceError("invalid_personal_email", 400, nil)
	ErrInvalidAuthChannel               = NewServiceError("invalid_auth_channel", 400, nil)
	ErrInvalidSMSProvider               = NewServiceError("invalid_sms_service_provider", 500, nil)
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
	ErrReferralNotFound                 = NewServiceError("referral_not_found", 404, nil)
	ErrSearchParamRequired              = NewServiceError("search_param_required", 400, nil)
	ErrCompanyNotFound                  = NewServiceError("company_not_found", 404, nil)
	ErrNoDataFound                      = NewServiceError("no_data_found", 404, nil)
	ErrSavingResume                     = NewServiceError("error_saving_resume_file", 500, nil)
	ErrOtpNotProvided                   = NewServiceError("otp_not_provided", 400, nil)
	ErrValueCannotBeEmpty               = NewServiceError("value_cannot_be_empty", 400, nil)
	ErrKeyCannotBeEmpty                 = NewServiceError("key_cannot_be_empty", 400, nil)
	ErrKeyNotFound                      = NewServiceError("key_not_found", 404, nil)
	ErrCalendalyLinkRequired            = NewServiceError("calendaly_link_required", 400, nil)
	ErrDescriptionRequired              = NewServiceError("description_required", 400, nil)
	ErrTagsRequired                     = NewServiceError("tags_required", 400, nil)
	ErrTagsLimitExceededByFive          = NewServiceError("tags_limit_exceeded_by_five", 400, nil)
	ErrDomainRequired                   = NewServiceError("domain_required", 400, nil)
	ErrUserNotApproved                  = NewServiceError("user_not_approved", 400, nil)
	ErrUserAlreadyMentor                = NewServiceError("user_already_mentor", 400, nil)
	ErrUserNotFound                     = NewServiceError("user_not_found", 404, nil)
	ErrUnableInsertUser                 = NewServiceError("unable_insert_user", 500, nil)
	ErrUnableInsertCompany              = NewServiceError("unable_insert_company", 500, nil)
	ErrGettingCompany                   = NewServiceError("error_getting_company", 500, nil)
	ErrGettingUserID                    = NewServiceError("error_getting_user_id", 500, nil)
	ErrUsersNotFound                    = NewServiceError("users_not_found", 404, nil)
	ErrGettingUser                      = NewServiceError("error_getting_user", 500, nil)
	ErrGettingUserReferrals             = NewServiceError("error_getting_user_referrals", 500, nil)
	ErrGettingRequesterReferrals        = NewServiceError("error_getting_requester_referrals", 500, nil)
	ErrGettingProviderReferrals         = NewServiceError("error_getting_provider_referrals", 500, nil)
	ErrUpdatingUserDetails              = NewServiceError("error_updating_user_details", 500, nil)
	ErrGettingCompanies                 = NewServiceError("error_getting_companies", 500, nil)
	ErrUpdateReferralDetails            = NewServiceError("error_updating_referral_details", 500, nil)
	ErrUpdateCompanyDetails             = NewServiceError("error_updating_company_details", 500, nil)
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
