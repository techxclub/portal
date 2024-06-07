package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type PhoneOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func NewPhoneOTPRequest(r *http.Request) (*PhoneOTPRequest, error) {
	var req PhoneOTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r PhoneOTPRequest) Validate() error {
	return IsValidPhoneNumber(r.PhoneNumber)
}

func (r PhoneOTPRequest) ToOTPGenerationObject() domain.OTPGeneration {
	return domain.OTPGeneration{
		Type:  constants.TwilioChannelSMS,
		Value: r.PhoneNumber,
	}
}
