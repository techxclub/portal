package config

import "github.com/techx/portal/constants"

var defaultRateLimitConfig = RateLimitConfig{
	Enabled:    true,
	Attempts:   10,
	WindowSecs: 3600,
}

type RateLimit struct {
	AdminUserList      RateLimitConfig `yaml:"ADMIN_USER_LIST" env:",prefix=ADMIN_USER_LIST_"`
	AdminUserUpdate    RateLimitConfig `yaml:"ADMIN_USER_UPDATE" env:",prefix=ADMIN_USER_UPDATE_"`
	AdminCompanyList   RateLimitConfig `yaml:"ADMIN_COMPANY_LIST" env:",prefix=ADMIN_COMPANY_LIST_"`
	AdminCompanyUpdate RateLimitConfig `yaml:"ADMIN_COMPANY_UPDATE" env:",prefix=ADMIN_COMPANY_UPDATE_"`
	GenerateOTP        RateLimitConfig `yaml:"GENERATE_OTP" env:",prefix=GENERATE_OTP_"`
	ResendOTP          RateLimitConfig `yaml:"RESEND_OTP" env:",prefix=RESEND_OTP_"`
	VerifyOTP          RateLimitConfig `yaml:"VERIFY_OTP" env:",prefix=VERIFY_OTP_"`
	UserRegister       RateLimitConfig `yaml:"USER_REGISTER" env:",prefix=USER_REGISTER_"`
	UserProfile        RateLimitConfig `yaml:"USER_PROFILE" env:",prefix=USER_PROFILE_"`
	CompanyList        RateLimitConfig `yaml:"COMPANY_LIST" env:",prefix=COMPANY_LIST_"`
	CompanyUserList    RateLimitConfig `yaml:"COMPANY_USER_LIST" env:",prefix=COMPANY_USER_LIST_"`
	ReferralRequest    RateLimitConfig `yaml:"REFERRAL_REQUEST" env:",prefix=REFERRAL_REQUEST_"`
	MentorList         RateLimitConfig `yaml:"MENTOR_LIST" env:",prefix=MENTOR_LIST_"`
	MentorRegister     RateLimitConfig `yaml:"MENTOR_REGISTER" env:",prefix=MENTOR_REGISTER_"`
}

type RateLimitConfig struct {
	Enabled    bool  `yaml:"ENABLED" env:"ENABLED"`
	Attempts   int64 `yaml:"ATTEMPTS" env:"ATTEMPTS"`
	WindowSecs int64 `yaml:"WINDOW_SECS" env:"WINDOW_SECS"`
}

func defaultRateLimit() RateLimit {
	return RateLimit{
		AdminUserList:      RateLimitConfig{Enabled: true, Attempts: 100},
		AdminUserUpdate:    RateLimitConfig{Enabled: true, Attempts: 100},
		AdminCompanyList:   RateLimitConfig{Enabled: true, Attempts: 100},
		AdminCompanyUpdate: RateLimitConfig{Enabled: true, Attempts: 100},
		GenerateOTP:        RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 600},
		ResendOTP:          RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 600},
		VerifyOTP:          RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		UserRegister:       RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 3600},
		UserProfile:        RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		CompanyList:        RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		CompanyUserList:    RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		ReferralRequest:    RateLimitConfig{Enabled: true, Attempts: 10, WindowSecs: 600},
		MentorList:         RateLimitConfig{Enabled: true, Attempts: 100, WindowSecs: 600},
		MentorRegister:     RateLimitConfig{Enabled: true, Attempts: 5, WindowSecs: 3600},
	}
}

func (rl RateLimit) GetAPIRateLimitConfig(apiName string) RateLimitConfig {
	switch apiName {
	case constants.APINameAdminUserList:
		return rl.AdminUserList
	case constants.APINameAdminUserUpdate:
		return rl.AdminUserUpdate
	case constants.APINameAdminCompanyList:
		return rl.AdminCompanyList
	case constants.APINameAdminCompanyUpdate:
		return rl.AdminCompanyUpdate
	case constants.APINameGenerateOTP:
		return rl.GenerateOTP
	case constants.APINameResendOTP:
		return rl.ResendOTP
	case constants.APINameVerifyOTP:
		return rl.VerifyOTP
	case constants.APINameUserRegister:
		return rl.UserRegister
	case constants.APINameUserProfile:
		return rl.UserProfile
	case constants.APINameCompanyList:
		return rl.CompanyList
	case constants.APINameCompanyUserList:
		return rl.CompanyUserList
	case constants.APINameReferralRequest:
		return rl.ReferralRequest
	case constants.APINameMentorRegister:
		return rl.MentorRegister
	case constants.APINameMentorList:
		return rl.MentorList
	default:
		return defaultRateLimitConfig
	}
}
