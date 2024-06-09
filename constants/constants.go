package constants

const (
	TableNameUsers = "users"
	RoleViewer     = "VIEWER"

	// User Info Params
	ParamUserID      = "user_id"
	ParamPhoneNumber = "phone_number"
	ParamFirstName   = "first_name"
	ParamLastName    = "last_name"
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
)
