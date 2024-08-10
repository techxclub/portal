package constants

const (
	// API Names
	APINameAdminUserList          = "AdminUserList"
	APINameAdminUserApprove       = "AdminUserApprove"
	APINameAdminUserUpdate        = "AdminUserUpdate"
	APINameAdminCompanyList       = "AdminCompanyList"
	APINameAdminCompanyUpdate     = "AdminCompanyUpdate"
	APINameAdminReferralList      = "AdminReferralList"
	APINameAdminReferralUpdate    = "AdminReferralUpdate"
	APINameAdminReferralExpire    = "AdminReferralExpire"
	APINameAdminFetchAuthToken    = "AdminFetchAuthToken"
	APINameAdminFetchCompanyLogo  = "AdminFetchCompanyLogo"
	APINameAdminUploadCompanyLogo = "AdminUploadCompanyLogo"
	APINameGoogleSignIn           = "GoogleSignIn"
	APINameGenerateOTP            = "GenerateOTP"
	APINameVerifyOTP              = "VerifyOTP"
	APINameResendOTP              = "ResendOTP"
	APINameUserFetchProfile       = "UserFetchProfile"
	APINameUserUpdateProfile      = "UserUpdateProfile"
	APINameUserRegister           = "UserRegister"
	APINameUserDashboard          = "UserDashboard"
	APINameMentorRegister         = "MentorRegister"
	APINameCompanyList            = "CompanyList"
	APINameCompanyUserList        = "CompanyUserList"
	APINameMentorList             = "MentorList"
	APINameReferralRequest        = "ReferralRequest"
	APINameReferralList           = "ReferralList"
	APINameReferralUpdate         = "ReferralUpdate"
)

var (
	AuthRoutes  = []string{APINameGoogleSignIn, APINameGenerateOTP, APINameVerifyOTP, APINameResendOTP}
	AdminRoutes = []string{APINameAdminUserList, APINameAdminUserUpdate, APINameAdminCompanyList, APINameAdminCompanyUpdate}
)
