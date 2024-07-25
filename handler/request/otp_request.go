package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
)

type OTPRequest struct {
	Channel string  `json:"channel"`
	Value   string  `json:"value"`
	OTP     *string `json:"otp,omitempty"`
}

func NewOTPRequest(r *http.Request) (*OTPRequest, error) {
	var req OTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r OTPRequest) Validate() error {
	return nil
}

func (r OTPRequest) ToAuthRequest() domain.OTPRequest {
	authRequest := domain.OTPRequest{
		Channel: r.Channel,
		Value:   r.Value,
	}

	if r.OTP != nil {
		authRequest.OTP = *r.OTP
	}

	return authRequest
}
