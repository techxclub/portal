package constants

const (
	DefaultLanguage = "en"
	TableNameUsers  = "users"

	// User status
	StatusPendingApproval = "PENDING_APPROVAL"
	StatusAutoApproved    = "AUTO_APPROVED"
	StatusApproved        = "APPROVED"

	// User Info Params
	ParamUserID      = "user_id"
	ParamStatus      = "status"
	ParamName        = "name"
	ParamPhoneNumber = "phone_number"
	ParamCompany     = "company"
	ParamRole        = "role"

	// Third party sms service provider
	ThirdPartyTwilio = "twilio"
	ThirdPartyMsg91  = "msg91"
	ThirdPartyMocked = "mocked"

	// User authorization
	AuthActionGenerate  = "GENERATE"
	AuthActionVerify    = "VERIFY"
	AuthChannelSMS      = "sms"
	AuthChannelWhatsapp = "whatsapp"
	AuthStatusApproved  = "approved"
	AuthStatusPending   = "pending"

	// Actions
	ActionRetryOTP        = "RETRY_OTP"
	ActionVerifyOTP       = "VERIFY_OTP"
	ActionSignUp          = "SIGN_UP"
	ActionPendingApproval = "PENDING_APPROVAL"
	ActionLogIn           = "LOG_IN"

	//	Referral status
	ReferralStatusPending = "PENDING"
)
