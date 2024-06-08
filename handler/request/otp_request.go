package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type GeneratePhoneOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func NewGeneratePhoneOTPRequest(r *http.Request) (*GeneratePhoneOTPRequest, error) {
	var req GeneratePhoneOTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r GeneratePhoneOTPRequest) Validate() error {
	return IsValidPhoneNumber(r.PhoneNumber)
}

func (r GeneratePhoneOTPRequest) ToOTPGenerationObject() domain.OTPGeneration {
	return domain.OTPGeneration{
		Type:  constants.AuthTypeSMS,
		Value: r.PhoneNumber,
	}
}

type VerifyOTPRequest struct {
	ID  string `json:"id"`
	OTP string `json:"otp"`
}

func NewPhoneOTPRequest(r *http.Request) (*VerifyOTPRequest, error) {
	var req VerifyOTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r VerifyOTPRequest) Validate() error {
	return nil
}

func (r VerifyOTPRequest) ToOTPVerificationObject() domain.OTPVerification {
	return domain.OTPVerification{
		Value: r.ID,
		Code:  r.OTP,
	}
}
