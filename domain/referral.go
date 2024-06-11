package domain

import "time"

type Referrals []Referral

type Referral struct {
	ID              string     `db:"id"`
	RequesterUserID string     `db:"requester_user_id"`
	ProviderUserID  string     `db:"referral_user_id"`
	JobLink         string     `db:"job_link"`
	Status          string     `db:"status"`
	CreatedAt       *time.Time `db:"created_time"`
}

type ReferralParams struct {
	ID              string
	RequesterUserID string
	ProviderUserID  string
	Status          string
}
