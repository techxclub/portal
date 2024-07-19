package constants

const (
	// API Names
	APINameAdminUserList      = "ADMIN_USER_LIST"
	APINameAdminUserUpdate    = "ADMIN_USER_UPDATE"
	APINameAdminCompanyList   = "ADMIN_COMPANY_LIST"
	APINameAdminCompanyUpdate = "ADMIN_COMPANY_UPDATE"
	APINameGenerateOTP        = "GENERATE_OTP"
	APINameVerifyOTP          = "VERIFY_OTP"
	APINameResendOTP          = "RESEND_OTP"
	APINameUserRegister       = "USER_REGISTER"
	APINameMentorRegister     = "MENTOR_REGISTER"
	APINameUserProfile        = "USER_PROFILE"
	APINameCompanyList        = "COMPANY_LIST"
	APINameCompanyUserList    = "COMPANY_USER_LIST"
	APINameMentorList         = "MENTOR_LIST"
	APINameReferralRequest    = "REFERRAL_REQUEST"
)

var (
	AuthRoutes  = []string{APINameGenerateOTP, APINameVerifyOTP, APINameResendOTP}
	AdminRoutes = []string{APINameAdminUserList, APINameAdminUserUpdate, APINameAdminCompanyList, APINameAdminCompanyUpdate}
)
