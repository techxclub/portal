package request

import (
	"encoding/json"
	"github.com/techx/portal/constants"
	"net/http"
	"net/url"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type ReferralRequest struct {
	RequesterUserID string `json:"requester_user_id"`
	ProviderUserID  string `json:"provider_user_id"`
	JobLink         string `json:"job_link"`
}

func NewReferralRequest(r *http.Request) (*ReferralRequest, error) {
	var req ReferralRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r ReferralRequest) Validate() error {
	if r.RequesterUserID == "" {
		return errors.New("Requester user id is required")
	}

	if r.ProviderUserID == "" {
		return errors.New("Provider user id is required")
	}

	_, err := url.ParseRequestURI(r.JobLink)
	if err != nil {
		return errors.New("Invalid job link")
	}
	return nil
}

func (r ReferralRequest) ToReferral() domain.Referral {
	return domain.Referral{
		RequesterUserID: r.RequesterUserID,
		ProviderUserID:  r.ProviderUserID,
		JobLink:         r.JobLink,
		Status:          constants.ReferralStatusPending,
	}
}
