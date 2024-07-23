package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
)

type ReferralUpdateRequest struct {
	ID                int64  `json:"id"`
	CompanyID         int64  `json:"company_id"`
	RequesterUserUUID string `db:"requester_user_id"`
	ProviderUserUUID  string `db:"provider_user_id"`
	JobLink           string `db:"job_link"`
	Status            string `db:"status"`
}

func NewReferralUpdateRequest(r *http.Request) (*ReferralUpdateRequest, error) {
	var req ReferralUpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r ReferralUpdateRequest) Validate() error {
	return nil
}

func (r ReferralUpdateRequest) ToReferralUpdateParams() *domain.Referral {
	return &domain.Referral{
		ID:                r.ID,
		CompanyID:         r.CompanyID,
		RequesterUserUUID: r.RequesterUserUUID,
		ProviderUserUUID:  r.ProviderUserUUID,
		JobLink:           r.JobLink,
		Status:            r.Status,
	}
}
