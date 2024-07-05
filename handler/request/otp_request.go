package request

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

var (
	supportedAuthChannels = []string{constants.AuthChannelSMS, constants.AuthChannelEmail}
	phoneAuthChannels     = []string{constants.AuthChannelSMS}
)

type OTPRequest struct {
	Action  string  `json:"-"`
	Channel string  `json:"channel"`
	Value   string  `json:"value"`
	OTP     *string `json:"otp,omitempty"`
}

func newOTPRequest(r *http.Request) (*OTPRequest, error) {
	var req OTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func NewGenerateOTPRequest(r *http.Request) (*OTPRequest, error) {
	req, err := newOTPRequest(r)
	if err != nil {
		return nil, err
	}

	req.Action = constants.AuthActionGenerate
	return req, nil
}

func NewVerifyOTPRequest(r *http.Request) (*OTPRequest, error) {
	req, err := newOTPRequest(r)
	if err != nil {
		return nil, err
	}

	req.Action = constants.AuthActionVerify
	return req, nil
}

func (r OTPRequest) Validate() error {
	if strings.ToUpper(r.Action) == constants.AuthActionVerify && r.OTP == nil {
		return errors.ErrMissingOTP
	}

	if !slices.Contains(supportedAuthChannels, r.Channel) {
		return errors.ErrInvalidAuthChannel
	}

	if slices.Contains(phoneAuthChannels, r.Channel) && !utils.IsValidPhoneNumber(r.Value) {
		return errors.ErrInvalidPhoneNumber
	}

	return nil
}

func (r OTPRequest) ToAuthRequest() domain.AuthRequest {
	authRequest := domain.AuthRequest{
		Channel: r.Channel,
		Value:   r.Value,
	}

	if slices.Contains(phoneAuthChannels, r.Channel) {
		authRequest.Value = utils.SanitizePhoneNumber(r.Value)
	}

	if strings.ToUpper(r.Action) == constants.AuthActionVerify {
		authRequest.OTP = *r.OTP
	}
	return authRequest
}
