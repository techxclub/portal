package request

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type GmailAuthRequest struct {
	Action   string  `json:"-"`
	AuthCode *string `json:"auth_code,omitempty"`
}

func newGmailAuthRequest(r *http.Request) (*GmailAuthRequest, error) {
	var req GmailAuthRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func NewVerifyGmailAuthRequest(r *http.Request) (*GmailAuthRequest, error) {
	req, err := newGmailAuthRequest(r)
	if err != nil {
		return nil, err
	}

	req.Action = constants.ActionLogInWithGoogle
	return req, nil
}


func (r GmailAuthRequest) Validate() error {
	if strings.ToUpper(r.Action) == constants.ActionLogInWithGoogle && r.AuthCode == nil {
		return errors.ErrMissingAuthCode
	}

	return nil
}


func (r GmailAuthRequest) ToAuthRequest() domain.GmailAuthRequest {
	authRequest := domain.GmailAuthRequest{}

	if strings.ToUpper(r.Action) == constants.ActionLogInWithGoogle {
		authRequest.AuthCode = *r.AuthCode
	}

	return authRequest
}
