package composers

import (
	"context"
	"time"

	"github.com/techx/portal/domain"
)

type Referral struct {
	ID              int64      `db:"id"`
	CompanyID       int64      `db:"company_id"`
	RequesterUserID string     `db:"requester_user_id"`
	ProviderUserID  string     `db:"provider_user_id"`
	JobLink         string     `db:"job_link"`
	Status          string     `db:"status"`
	CreatedAt       *time.Time `db:"created_time"`
}

func ReferralListResponse(_ context.Context, referrals domain.Referrals) []Referral {
	referralList := make([]Referral, 0, len(referrals))
	for _, referral := range referrals {
		referralList = append(referralList, Referral{
			ID:              referral.ID,
			CompanyID:       referral.CompanyID,
			RequesterUserID: referral.RequesterUserID,
			ProviderUserID:  referral.ProviderUserID,
			JobLink:         referral.JobLink,
			Status:          referral.Status,
			CreatedAt:       referral.CreatedAt,
		})
	}

	return referralList
}
