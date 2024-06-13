package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

const (
	insertReferralQuery = `INSERT INTO referrals (requester_user_id, provider_user_id, company, job_link, status, created_time) VALUES (:requester_user_id, :provider_user_id, :company, :job_link, :status, :created_time) RETURNING id, created_time`

	getReferralSelectorFields = `id, requester_user_id, provider_user_id, company, job_link, status, created_time`
)

type ReferralsRepo interface {
	CreateReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
	GetReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error)
}

type referralsRepo struct {
	dbClient *db.Repository
}

type ReferralsReturning struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_time"`
}

func NewReferralsRepo(dbClient *db.Repository) ReferralsRepo {
	return &referralsRepo{
		dbClient: dbClient,
	}
}

func (r referralsRepo) CreateReferral(ctx context.Context, params domain.ReferralParams) (*domain.Referral, error) {
	var returning ReferralsReturning
	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertReferralQuery, map[string]interface{}{
			"requester_user_id": params.RequesterUserID,
			"provider_user_id":  params.ProviderUserID,
			"company":           params.Company,
			"job_link":          params.JobLink,
			"status":            params.Status,
			"created_time":      now,
		})
	})
	if err != nil {
		return nil, err
	}

	referral := domain.Referral{
		ID:              returning.ID,
		RequesterUserID: params.RequesterUserID,
		ProviderUserID:  params.ProviderUserID,
		Company:         params.Company,
		JobLink:         params.JobLink,
		Status:          params.Status,
		CreatedAt:       &returning.CreatedAt,
	}

	return &referral, nil
}

func (r referralsRepo) GetReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error) {
	getReferralsByParamsQuery, args, err := getReferralQueryForParams(params)
	if err != nil {
		return nil, err
	}

	var referrals []domain.Referral
	err = r.dbClient.DBSelect(ctx, &referrals, getReferralsByParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	result := domain.Referrals(referrals)
	return &result, nil
}

func getReferralQueryForParams(params domain.ReferralParams) (string, []interface{}, error) {
	counter := 1
	args := make([]interface{}, 0)
	conditions := make([]string, 0)
	for key, value := range params.ToMap() {
		if value == "" {
			continue
		}

		condition := fmt.Sprintf("%s = $%d", key, counter)
		conditions = append(conditions, condition)
		args = append(args, value)
		counter++
	}

	// If no params are provided, return an error
	if len(conditions) == 0 {
		return "", nil, errors.ErrSearchParamRequired
	}

	baseQuery := `SELECT ` + getReferralSelectorFields + ` FROM referrals WHERE `
	query := baseQuery + strings.Join(conditions, " AND ")
	return query, args, nil
}
