package domain

type GoogleLogin struct {
	RedirectURI string
}

type GoogleOAuthExchangeRequest struct {
	InviteCode string `json:"invite_code"`
	Code       string `json:"code"`
	OAuthCode  string `json:"oauth_code"`
}

type OTPRequest struct {
	Channel string
	Value   string
	OTP     string
}

type AuthDetails struct {
	Status string
}
