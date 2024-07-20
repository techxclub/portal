package constants

const (
	// API Names
	APINameAdminUserList       = "AdminUserList"
	APINameAdminUserUpdate     = "AdminUserUpdate"
	APINameAdminCompanyList    = "AdminCompanyList"
	APINameAdminCompanyUpdate  = "AdminCompanyUpdate"
	APINameAdminReferralList   = "AdminReferralList"
	APINameAdminReferralUpdate = "AdminReferralUpdate"
	APINameAdminReferralExpire = "AdminReferralExpire"
	APINameGoogleOAuthDebug    = "GoogleOAuthDebug"
	APINameGoogleOAuthLogin    = "GoogleOAuthLogin"
	APINameGoogleOAuthCallback = "GoogleOAuthCallback"
	APINameGenerateOTP         = "GenerateOTP"
	APINameVerifyOTP           = "VerifyOTP"
	APINameResendOTP           = "ResendOTP"
	APINameUserRegister        = "UserRegister"
	APINameMentorRegister      = "MentorRegister"
	APINameUserProfile         = "UserProfile"
	APINameCompanyList         = "CompanyList"
	APINameCompanyUserList     = "CompanyUserList"
	APINameMentorList          = "MentorList"
	APINameReferralRequest     = "ReferralRequest"
	APINameReferralList        = "ReferralList"
	APINameReferralUpdate      = "ReferralUpdate"
)

var (
	AuthRoutes  = []string{APINameGoogleOAuthLogin, APINameGoogleOAuthCallback, APINameGenerateOTP, APINameVerifyOTP, APINameResendOTP}
	AdminRoutes = []string{APINameAdminUserList, APINameAdminUserUpdate, APINameAdminCompanyList, APINameAdminCompanyUpdate}
)
