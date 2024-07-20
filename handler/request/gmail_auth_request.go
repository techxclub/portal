package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type GoogleOAuthRequest struct {
	AuthCode string `json:"auth_code"`
}

func NewGoogleOAuthLoginRequest(r *http.Request) (*GoogleOAuthRequest, error) {
	var req GoogleOAuthRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r GoogleOAuthRequest) Validate() error {
	if r.AuthCode == "" {
		return errors.ErrMissingAuthCode
	}

	return nil
}

func (r GoogleOAuthRequest) ToAuthRequest() domain.GoogleOAuthCallbackRequest {
	return domain.GoogleOAuthCallbackRequest{
		Code: r.AuthCode,
	}
}
