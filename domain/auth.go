package domain

type AuthRequest struct {
	Channel string
	Value   string
	OTP     string
}

type AuthDetails struct {
	UserInfo *UserProfile
	AuthInfo AuthInfo
}

type AuthInfo struct {
	Status string
}
