package domain

type GoogleLogin struct {
	RedirectURI string
}

type GoogleOAuthExchangeRequest struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

type OTPRequest struct {
	Channel string
	Value   string
	OTP     string
}

type AuthDetails struct {
	Status string
}
