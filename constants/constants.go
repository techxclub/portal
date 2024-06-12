package constants

const (
	DefaultLanguage = "en"
	TableNameUsers  = "users"
	RoleViewer      = "VIEWER"

	// User Info Params
	ParamUserID      = "user_id"
	ParamPhoneNumber = "phone_number"
	ParamName        = "name"
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
