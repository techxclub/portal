package domain

import (
	"time"
)

type Referrals []Referral

type Referral struct {
	ID              int64      `db:"id"`
	CompanyID       int64      `db:"company_id"`
	RequesterUserID string     `db:"requester_user_id"`
	ProviderUserID  string     `db:"provider_user_id"`
	JobLink         string     `db:"job_link"`
	Status          string     `db:"status"`
	CreatedAt       *time.Time `db:"created_time"`
}

type ReferralParams struct {
	ID              int64
	CompanyID       int64
	RequesterUserID string
	ProviderUserID  string
	CompanyName     string
	Message         string
	JobLink         string
	Status          string
	ResumeFilePath  string
	CreatedAt       *time.Time
}
