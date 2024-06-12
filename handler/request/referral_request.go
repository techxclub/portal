package request

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type ReferralRequest struct {
	RequesterUserID string `json:"requester_user_id"`
	ProviderUserID  string `json:"provider_user_id"`
	Company         string `json:"company"`
	JobLink         string `json:"job_link"`
	Message         string `json:"message"`
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
		return errors.ErrRequesterFieldIsEmpty
	}

	if r.ProviderUserID == "" {
		return errors.ErrProviderFieldIsEmpty
	}

	if r.Company == "" {
		return errors.ErrCompanyRequired
	}

	_, err := url.ParseRequestURI(r.JobLink)
	if err != nil {
		return errors.ErrInvalidJobLink
	}
	return nil
}

func (r ReferralRequest) ToReferral() domain.ReferralParams {
	return domain.ReferralParams{
		RequesterUserID: r.RequesterUserID,
		ProviderUserID:  r.ProviderUserID,
		JobLink:         r.JobLink,
		Company:         r.Company,
		Message:         r.Message,
		Status:          constants.ReferralStatusPending,
	}
}
