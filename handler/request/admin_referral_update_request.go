package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
)

type AdminReferralUpdateRequest struct {
	ID              int64  `json:"id"`
	CompanyID       int64  `json:"company_id"`
	RequesterUserID string `db:"requester_user_id"`
	ProviderUserID  string `db:"provider_user_id"`
	JobLink         string `db:"job_link"`
	Status          string `db:"status"`
}

func NewAdminReferralUpdateRequest(r *http.Request) (*AdminReferralUpdateRequest, error) {
	var req AdminReferralUpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r AdminReferralUpdateRequest) Validate() error {
	return nil
}

func (r AdminReferralUpdateRequest) ToReferralParams() *domain.Referral {
	return &domain.Referral{
		ID:              r.ID,
		CompanyID:       r.CompanyID,
		RequesterUserID: r.RequesterUserID,
		ProviderUserID:  r.ProviderUserID,
		JobLink:         r.JobLink,
		Status:          r.Status,
	}
}
