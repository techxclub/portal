package domain

import (
	"mime/multipart"
	"time"
)

type UserReferrals struct {
	RequestedReferrals *Referrals
	ProvidedReferrals  *Referrals
}

type Referrals []Referral

type Referral struct {
	ID                int64      `db:"id"`
	CompanyID         int64      `db:"company_id"`
	RequesterUserUUID string     `db:"requester_user_id"`
	ProviderUserUUID  string     `db:"provider_user_id"`
	JobLink           string     `db:"job_link"`
	Status            string     `db:"status"`
	CreatedAt         *time.Time `db:"create_time"`
	UpdatedAt         *time.Time `db:"update_time"`
}

type ReferralParams struct {
	Referral
	CompanyName       string
	NoticePeriod      string
	PreferredLocation string
	Message           string
	ResumeFile        multipart.File
}

func (param ReferralParams) ToReferral() Referral {
	return Referral{
		CompanyID:         param.CompanyID,
		RequesterUserUUID: param.RequesterUserUUID,
		ProviderUserUUID:  param.ProviderUserUUID,
		JobLink:           param.JobLink,
		Status:            param.Status,
	}
}
