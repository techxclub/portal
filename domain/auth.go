package domain

type GoogleLogin struct {
	RedirectURI string
}

type GoogleOAuthCallbackRequest struct {
	State string
	Code  string
}

type OTPRequest struct {
	Channel string
	Value   string
	OTP     string
}

type AuthDetails struct {
	Token    string
	UserInfo *User
	AuthInfo AuthInfo
}

type AuthInfo struct {
	Status string
}
