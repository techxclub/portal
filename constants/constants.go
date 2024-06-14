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
	ThirdPartMsg91   = "msg91"

	// User authorization
	AuthActionGenerate  = "GENERATE"
	AuthActionVerify    = "VERIFY"
	AuthChannelSMS      = "sms"
	AuthChannelWhatsapp = "whatsapp"
	AuthStatusApproved  = "approved"
	AuthStatusPending   = "pending"

	// Actions
	ActionRetry  = "RETRY"
	ActionLogIn  = "LOG_IN"
	ActionSignUp = "SIGN_UP"

	//	Referral status
	ReferralStatusPending = "PENDING"
)
