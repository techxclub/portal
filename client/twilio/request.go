package twilio

type CreateVerificationRequest struct {
	To      string `json:"-"`
	Channel string `json:"-"`
}

type CheckVerificationRequest struct {
	From string `json:"-"`
	OTP  string `json:"-"`
}

func NewCreateVerificationRequest(to, channel string) CreateVerificationRequest {
	return CreateVerificationRequest{
		To:      to,
		Channel: channel,
	}
}

func NewCheckVerificationRequest(from, otp string) CheckVerificationRequest {
	return CheckVerificationRequest{
		From: from,
		OTP:  otp,
	}
}
