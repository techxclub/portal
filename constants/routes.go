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
	APINameGoogleOAuthExchange = "GoogleOAuthExchange"
	APINameGenerateOTP         = "GenerateOTP"
	APINameVerifyOTP           = "VerifyOTP"
	APINameResendOTP           = "ResendOTP"
	APINameUserFetchProfile    = "UserFetchProfile"
	APINameUserUpdateProfile   = "UserUpdateProfile"
	APINameUserRegister        = "UserRegister"
	APINameMentorRegister      = "MentorRegister"
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
