package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/domain"
)

const (
	insertReferralQuery = `INSERT INTO referrals (requester_user_id, provider_user_id, job_link, status, created_time) VALUES (:requester_user_id, :provider_user_id, :job_link, :status, :created_time) RETURNING id`
)

type ReferralsRepo interface {
	CreateReferral(ctx context.Context, referral domain.Referral) (*domain.Referral, error)
}

type referralsRepo struct {
	dbClient *db.Repository
}

type ReferralsReturning struct {
	ID string `db:"id"`
}

func NewReferralsRepo(dbClient *db.Repository) ReferralsRepo {
	return &referralsRepo{
		dbClient: dbClient,
	}
}

func (r referralsRepo) CreateReferral(ctx context.Context, referral domain.Referral) (*domain.Referral, error) {
	var returning ReferralsReturning
	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertReferralQuery, map[string]interface{}{
			"requester_user_id": referral.RequesterUserID,
			"provider_user_id":  referral.ProviderUserID,
			"job_link":          referral.JobLink,
			"status":            referral.Status,
			"created_time":      now,
		})
	})
	if err != nil {
		return nil, err
	}

	referral.ID = returning.ID
	return &referral, nil
}
