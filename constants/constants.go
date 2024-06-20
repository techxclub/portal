package constants

const (
	DefaultLanguage = "en"
	TableNameUsers  = "users"

	// Actors
	ActorSystem = "SYSTEM"
	ActorUser   = "USER"
	ActorAdmin  = "ADMIN"

	// User status
	StatusPendingApproval = "PENDING_APPROVAL"
	StatusAutoApproved    = "AUTO_APPROVED"
	StatusApproved        = "APPROVED"

	// DB Fetch Params
	ParamID              = "id"
	ParamActor           = "actor"
	ParamUserIDNum       = "user_id_num"
	ParamUserID          = "user_id"
	ParamStatus          = "status"
	ParamName            = "name"
	ParamPhoneNumber     = "phone_number"
	ParamPersonalEmail   = "personal_email"
	ParamWorkEmail       = "work_email"
	ParamCompany         = "company"
	ParamRole            = "role"
	ParamCreatedTime     = "created_time"
	ParamRequesterID     = "requester_user_id"
	ParamProviderID      = "provider_user_id"
	ParamJobLink         = "job_link"
	ParamSmallLogo       = "small_logo"
	ParamBigLogo         = "big_logo"
	ParamOfficialWebsite = "official_website"
	ParamCareersPage     = "careers_page"
	ParamPriority        = "priority"
	ParamVerified        = "verified"
	ParamPopular         = "popular"

	// Third party sms service provider
	ThirdPartyMsg91  = "msg91"
	ThirdPartyGomail = "gomail"

	// User authorization
	AuthActionGenerate = "GENERATE"
	AuthActionVerify   = "VERIFY"

	AuthChannelSMS      = "sms"
	AuthChannelEmail    = "email"
	AuthChannelWhatsapp = "whatsapp"

	AuthStatusGenerated = "generated"
	AuthStatusPending   = "pending"
	AuthStatusResent    = "resent"
	AuthStatusVerified  = "verified"
	AuthStatusFailed    = "failed"

	// Actions
	ActionRetryOTP        = "RETRY_OTP"
	ActionVerifyOTP       = "VERIFY_OTP"
	ActionSignUp          = "SIGN_UP"
	ActionPendingApproval = "PENDING_APPROVAL"
	ActionLogIn           = "LOG_IN"

	//	Referral status
	ReferralStatusPending = "PENDING"
)
