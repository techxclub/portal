package config

import "github.com/techx/portal/constants"

type RateLimit struct {
	DefaultConfig       RateLimitConfig `yaml:"DEFAULT_CONFIG" env:"DEFAULT_CONFIG"`
	AdminUserList       RateLimitConfig `yaml:"ADMIN_USER_LIST" env:",prefix=ADMIN_USER_LIST_"`
	AdminUserApprove    RateLimitConfig `yaml:"ADMIN_USER_APPROVE" env:",prefix=ADMIN_USER_APPROVE_"`
	AdminUserUpdate     RateLimitConfig `yaml:"ADMIN_USER_UPDATE" env:",prefix=ADMIN_USER_UPDATE_"`
	AdminCompanyList    RateLimitConfig `yaml:"ADMIN_COMPANY_LIST" env:",prefix=ADMIN_COMPANY_LIST_"`
	AdminCompanyUpdate  RateLimitConfig `yaml:"ADMIN_COMPANY_UPDATE" env:",prefix=ADMIN_COMPANY_UPDATE_"`
	AdminReferralList   RateLimitConfig `yaml:"ADMIN_REFERRAL_LIST" env:",prefix=ADMIN_REFERRAL_LIST_"`
	AdminReferralUpdate RateLimitConfig `yaml:"ADMIN_REFERRAL_UPDATE" env:",prefix=ADMIN_REFERRAL_UPDATE_"`
	AdminReferralExpire RateLimitConfig `yaml:"ADMIN_REFERRAL_EXPIRE" env:",prefix=ADMIN_REFERRAL_EXPIRE_"`
	AdminFetchAuthToken RateLimitConfig `yaml:"ADMIN_FETCH_AUTH_TOKEN" env:",prefix=ADMIN_FETCH_AUTH_TOKEN_"`
	GoogleSignIn        RateLimitConfig `yaml:"GOOGLE_SIGN_IN" env:",prefix=GOOGLE_SIGN_IN_"`
	GenerateOTP         RateLimitConfig `yaml:"GENERATE_OTP" env:",prefix=GENERATE_OTP_"`
	ResendOTP           RateLimitConfig `yaml:"RESEND_OTP" env:",prefix=RESEND_OTP_"`
	VerifyOTP           RateLimitConfig `yaml:"VERIFY_OTP" env:",prefix=VERIFY_OTP_"`
	UserFetchProfile    RateLimitConfig `yaml:"USER_FETCH_PROFILE" env:",prefix=USER_FETCH_PROFILE_"`
	UserUpdateProfile   RateLimitConfig `yaml:"USER_UPDATE_PROFILE" env:",prefix=USER_UPDATE_PROFILE_"`
	UserRegister        RateLimitConfig `yaml:"USER_REGISTER" env:",prefix=USER_REGISTER_"`
	UserDashboard       RateLimitConfig `yaml:"USER_DASHBOARD" env:",prefix=USER_DASHBOARD_"`
	CompanyList         RateLimitConfig `yaml:"COMPANY_LIST" env:",prefix=COMPANY_LIST_"`
	CompanyUserList     RateLimitConfig `yaml:"COMPANY_USER_LIST" env:",prefix=COMPANY_USER_LIST_"`
	ReferralRequest     RateLimitConfig `yaml:"REFERRAL_REQUEST" env:",prefix=REFERRAL_REQUEST_"`
	ReferralList        RateLimitConfig `yaml:"REFERRAL_LIST" env:",prefix=REFERRAL_LIST_"`
	ReferralUpdate      RateLimitConfig `yaml:"REFERRAL_UPDATE" env:",prefix=REFERRAL_UPDATE_"`
	MentorList          RateLimitConfig `yaml:"MENTOR_LIST" env:",prefix=MENTOR_LIST_"`
	MentorRegister      RateLimitConfig `yaml:"MENTOR_REGISTER" env:",prefix=MENTOR_REGISTER_"`
}

type RateLimitConfig struct {
	Enabled    bool  `yaml:"ENABLED" env:"ENABLED"`
	Attempts   int64 `yaml:"ATTEMPTS" env:"ATTEMPTS"`
	WindowSecs int64 `yaml:"WINDOW_SECS" env:"WINDOW_SECS"`
}

func defaultRateLimit() RateLimit {
	return RateLimit{
		DefaultConfig:       RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 3600},
		AdminUserList:       RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminUserApprove:    RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminUserUpdate:     RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminCompanyList:    RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminCompanyUpdate:  RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminReferralList:   RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminReferralUpdate: RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		AdminReferralExpire: RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		GoogleSignIn:        RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 600},
		GenerateOTP:         RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 600},
		ResendOTP:           RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 600},
		VerifyOTP:           RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		UserFetchProfile:    RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		UserUpdateProfile:   RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		UserRegister:        RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 3600},
		UserDashboard:       RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 3600},
		CompanyList:         RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		CompanyUserList:     RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		ReferralRequest:     RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		ReferralList:        RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		ReferralUpdate:      RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		MentorList:          RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		MentorRegister:      RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 3600},
	}
}

func (rl RateLimit) GetAPIRateLimitConfig(apiName string) RateLimitConfig {
	switch apiName {
	case constants.APINameAdminUserList:
		return rl.AdminUserList
	case constants.APINameAdminUserApprove:
		return rl.AdminUserApprove
	case constants.APINameAdminUserUpdate:
		return rl.AdminUserUpdate
	case constants.APINameAdminCompanyList:
		return rl.AdminCompanyList
	case constants.APINameAdminCompanyUpdate:
		return rl.AdminCompanyUpdate
	case constants.APINameAdminReferralList:
		return rl.AdminReferralList
	case constants.APINameAdminReferralUpdate:
		return rl.AdminReferralUpdate
	case constants.APINameAdminReferralExpire:
		return rl.AdminReferralExpire
	case constants.APINameAdminFetchAuthToken:
		return rl.AdminFetchAuthToken
	case constants.APINameGoogleSignIn:
		return rl.GoogleSignIn
	case constants.APINameGenerateOTP:
		return rl.GenerateOTP
	case constants.APINameResendOTP:
		return rl.ResendOTP
	case constants.APINameVerifyOTP:
		return rl.VerifyOTP
	case constants.APINameUserFetchProfile:
		return rl.UserFetchProfile
	case constants.APINameUserUpdateProfile:
		return rl.UserUpdateProfile
	case constants.APINameUserRegister:
		return rl.UserRegister
	case constants.APINameUserDashboard:
		return rl.UserDashboard
	case constants.APINameCompanyList:
		return rl.CompanyList
	case constants.APINameCompanyUserList:
		return rl.CompanyUserList
	case constants.APINameReferralRequest:
		return rl.ReferralRequest
	case constants.APINameReferralList:
		return rl.ReferralList
	case constants.APINameReferralUpdate:
		return rl.ReferralUpdate
	case constants.APINameMentorRegister:
		return rl.MentorRegister
	case constants.APINameMentorList:
		return rl.MentorList
	default:
		return rl.DefaultConfig
	}
}
