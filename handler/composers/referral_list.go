package composers

import (
	"context"
	"time"

	"github.com/techx/portal/domain"
)

type Referral struct {
	ID                int64      `json:"id"`
	CompanyID         int64      `json:"company_id"`
	RequesterUserUUID string     `json:"requester_user_id"`
	ProviderUserUUID  string     `json:"provider_user_id"`
	JobLink           string     `json:"job_link"`
	Status            string     `json:"status"`
	CreatedAt         *time.Time `json:"create_time"`
	UpdatedAt         *time.Time `json:"update_time"`
}

func ReferralListResponse(_ context.Context, referrals domain.Referrals) []Referral {
	referralList := make([]Referral, 0, len(referrals))
	for _, referral := range referrals {
		referralList = append(referralList, Referral{
			ID:                referral.ID,
			CompanyID:         referral.CompanyID,
			RequesterUserUUID: referral.RequesterUserUUID,
			ProviderUserUUID:  referral.ProviderUserUUID,
			JobLink:           referral.JobLink,
			Status:            referral.Status,
			CreatedAt:         referral.CreatedAt,
		})
	}

	return referralList
}
