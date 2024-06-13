package domain

import (
	"time"
)

type Referrals []Referral

type Referral struct {
	ID              string     `db:"id"`
	RequesterUserID string     `db:"requester_user_id"`
	ProviderUserID  string     `db:"provider_user_id"`
	Company         string     `db:"company"`
	JobLink         string     `db:"job_link"`
	Status          string     `db:"status"`
	CreatedAt       *time.Time `db:"created_time"`
}

type ReferralParams struct {
	ID              string
	RequesterUserID string
	ProviderUserID  string
	Company         string
	Message         string
	JobLink         string
	Status          string
	CreatedAt       *time.Time
}

func (r ReferralParams) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                r.ID,
		"requester_user_id": r.RequesterUserID,
		"provider_user_id":  r.ProviderUserID,
		"company":           r.Company,
		"job_link":          r.JobLink,
		"status":            r.Status,
		"created_time":      r.CreatedAt,
	}
}
