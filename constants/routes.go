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
	APINameAdminFetchAuthToken = "AdminFetchAuthToken"
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
	AuthRoutes  = []string{APINameGoogleOAuthExchange, APINameGenerateOTP, APINameVerifyOTP, APINameResendOTP}
	AdminRoutes = []string{APINameAdminUserList, APINameAdminUserUpdate, APINameAdminCompanyList, APINameAdminCompanyUpdate}
)
