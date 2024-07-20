package domain

import (
	"mime/multipart"
	"time"
)

type Referrals []Referral

type Referral struct {
	ID                int64      `db:"id"`
	CompanyID         int64      `db:"company_id"`
	RequesterUserUUID string     `db:"requester_user_id"`
	ProviderUserUUID  string     `db:"provider_user_id"`
	JobLink           string     `db:"job_link"`
	Status            string     `db:"status"`
	CreatedAt         *time.Time `db:"created_time"`
}

type ReferralParams struct {
	ID                int64
	CompanyID         int64
	RequesterUserUUID string
	ProviderUserUUID  string
	CompanyName       string
	Message           string
	JobLink           string
	Status            string
	CreatedAt         *time.Time
	ResumeFile        multipart.File
}
