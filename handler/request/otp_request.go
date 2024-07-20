package request

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

var (
	supportedAuthChannels = []string{constants.OTPChannelSMS, constants.OTPChannelEmail}
	phoneAuthChannels     = []string{constants.OTPChannelSMS}
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
	if !slices.Contains(supportedAuthChannels, r.Channel) {
		return errors.ErrInvalidAuthChannel
	}

	if slices.Contains(phoneAuthChannels, r.Channel) && !utils.IsValidPhoneNumber(r.Value) {
		return errors.ErrInvalidPhoneNumber
	}

	return nil
}

func (r OTPRequest) ToAuthRequest() domain.OTPRequest {
	authRequest := domain.OTPRequest{
		Channel: r.Channel,
		Value:   r.Value,
	}

	if slices.Contains(phoneAuthChannels, r.Channel) {
		authRequest.Value = utils.SanitizePhoneNumber(r.Value)
	}

	if r.OTP != nil {
		authRequest.OTP = *r.OTP
	}

	return authRequest
}
