package constants

const (
	DefaultLanguage    = "en"
	GlobalRateLimitKey = "global_rate_limit"

	// Actors
	ActorUser = "USER"

	// Genders
	GenderMale = "MALE"

	// User status
	StatusIncompleteProfile = "INCOMPLETE_PROFILE"
	StatusPendingApproval   = "PENDING_APPROVAL"
	StatusAutoApproved      = "AUTO_APPROVED"
	StatusApproved          = "APPROVED"

	// Mentor config
	MentorStatusNotApproved     = "NOT_APPROVED"
	MentorStatusPendingApproval = "PENDING_APPROVAL"
	MentorStatusApproved        = "APPROVED"

	// User Params
	ParamUserNumber           = "user_number"
	ParamUserUUID             = "user_uuid"
	ParamCreateTime           = "create_time"
	ParamUpdateTime           = "update_time"
	ParamStatus               = "status"
	ParamGoogleOAuth          = "google_auth_details"
	ParamTechnicalInformation = "technical_information"
	ParamMentorConfig         = "mentor_config"

	// personal details
	ParamName            = "name"
	ParamPhoneNumber     = "phone_number"
	ParamRegisteredEmail = "registered_email"
	ParamProfilePicture  = "profile_picture"
	ParamLinkedIn        = "linkedin"
	ParamGender          = "gender"

	// professional details
	ParamCompanyID         = "company_id"
	ParamCompanyName       = "company_name"
	ParamWorkEmail         = "work_email"
	ParamDesignation       = "designation"
	ParamYearsOfExperience = "years_of_experience"

	// Company Params
	ParamID              = "id"
	ParamActor           = "actor"
	ParamNormalizedName  = "normalized_name"
	ParamDisplayName     = "display_name"
	ParamSmallLogo       = "small_logo"
	ParamBigLogo         = "big_logo"
	ParamOfficialWebsite = "official_website"
	ParamCareersPage     = "careers_page"
	ParamPriority        = "priority"
	ParamVerified        = "verified"
	ParamPopular         = "popular"

	// Referral Params
	ParamRequesterID = "requester_user_id"
	ParamProviderID  = "provider_user_id"
	ParamJobLink     = "job_link"
	ParamResumeFile  = "resume_file"
	ParamMessage     = "message"

	OTPStatusGenerated = "generated"
	OTPStatusPending   = "pending"
	OTPStatusVerified  = "verified"
	OTPStatusFailed    = "failed"

	// Actions
	ActionRetryOTP        = "RETRY_OTP"
	ActionVerifyOTP       = "VERIFY_OTP"
	ActionSuccess         = "SIGN_UP"
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
