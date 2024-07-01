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

	// Mentor config
	MentorStatusNotApproved     = "NOT_APPROVED"
	MentorStatusPendingApproval = "PENDING_APPROVAL"
	MentorStatusApproved        = "APPROVED"
	ParamMentorConfig           = "mentor_config"
	ParamMentorConfigStatus     = "status"

	// DB Fetch Params
	ParamID                = "id"
	ParamActor             = "actor"
	ParamUserIDNum         = "user_id_num"
	ParamUserID            = "user_id"
	ParamStatus            = "status"
	ParamName              = "name"
	ParamNormalizedName    = "normalized_name"
	ParamDisplayName       = "display_name"
	ParamPhoneNumber       = "phone_number"
	ParamPersonalEmail     = "personal_email"
	ParamWorkEmail         = "work_email"
	ParamCompanyID         = "company_id"
	ParamCompanyName       = "company_name"
	ParamRole              = "role"
	ParamYearsOfExperience = "years_of_experience"
	ParamLinkedIn          = "linkedin"
	ParamCreatedTime       = "created_time"
	ParamRequesterID       = "requester_user_id"
	ParamProviderID        = "provider_user_id"
	ParamJobLink           = "job_link"
	ParamSmallLogo         = "small_logo"
	ParamBigLogo           = "big_logo"
	ParamOfficialWebsite   = "official_website"
	ParamCareersPage       = "careers_page"
	ParamPriority          = "priority"
	ParamVerified          = "verified"
	ParamPopular           = "popular"
	ParamResumeFile        = "resume_file"
	ParamMessage           = "message"

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

	//	Gomail constants
	GomailHeaderFrom       = "From"
	GomailHeaderTo         = "To"
	GomailHeaderSubject    = "Subject"
	GomailHeaderMessageID  = "Message-ID"
	GomailHeaderInReplyTo  = "In-Reply-To"
	GomailHeaderReferences = "References"
	GomailContentTypeHTML  = "text/html"
)
