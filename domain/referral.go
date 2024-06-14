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

func (r ReferralParams) GetQueryConditions() (string, []interface{}) {
	qb := NewQueryBuilder()
	qb.AddEqualParam("id", r.ID)
	qb.AddEqualParam("requester_user_id", r.RequesterUserID)
	qb.AddEqualParam("provider_user_id", r.ProviderUserID)
	qb.AddEqualParam("company", r.Company)
	qb.AddEqualParam("status", r.Status)
	qb.AddGreaterEqualParam("created_time", r.CreatedAt)

	return qb.Build()
}
