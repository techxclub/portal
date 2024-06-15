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
	qb.AddEqualCondition("id", r.ID)
	qb.AddEqualCondition("requester_user_id", r.RequesterUserID)
	qb.AddEqualCondition("provider_user_id", r.ProviderUserID)
	qb.AddEqualCondition("company", r.Company)
	qb.AddEqualCondition("status", r.Status)
	qb.AddGreaterEqualCondition("created_time", r.CreatedAt)

	return qb.Build()
}
