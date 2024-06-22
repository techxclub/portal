package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

const (
	insertReferralQuery       = `INSERT INTO referrals (requester_user_id, provider_user_id, company_id, job_link, status, created_time) VALUES (:requester_user_id, :provider_user_id, :company_id, :job_link, :status, :created_time) RETURNING id, created_time`
	getReferralSelectorFields = `id, requester_user_id, provider_user_id, company_id, job_link, status, created_time`
	selectReferralBaseQuery   = `SELECT ` + getReferralSelectorFields + ` FROM referrals WHERE `
)

type ReferralsRepository interface {
	CreateReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
	GetReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error)
}

type referralsRepository struct {
	dbClient *db.Repository
}

type ReferralsReturning struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_time"`
}

func NewReferralsRepository(dbClient *db.Repository) ReferralsRepository {
	return &referralsRepository{
		dbClient: dbClient,
	}
}

func (r referralsRepository) CreateReferral(ctx context.Context, params domain.ReferralParams) (*domain.Referral, error) {
	var returning ReferralsReturning
	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertReferralQuery, map[string]interface{}{
			constants.ParamRequesterID: params.RequesterUserID,
			constants.ParamProviderID:  params.ProviderUserID,
			constants.ParamCompanyID:   params.CompanyID,
			constants.ParamJobLink:     params.JobLink,
			constants.ParamStatus:      params.Status,
			constants.ParamCreatedTime: now,
		})
	})
	if err != nil {
		return nil, err
	}

	referral := domain.Referral{
		ID:              returning.ID,
		RequesterUserID: params.RequesterUserID,
		ProviderUserID:  params.ProviderUserID,
		CompanyID:       params.CompanyID,
		JobLink:         params.JobLink,
		Status:          params.Status,
		CreatedAt:       &returning.CreatedAt,
	}

	return &referral, nil
}

func (r referralsRepository) GetReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamRequesterID, params.RequesterUserID)
	qb.AddEqualCondition(constants.ParamProviderID, params.ProviderUserID)
	qb.AddEqualCondition(constants.ParamCompanyID, params.CompanyID)
	qb.AddEqualCondition(constants.ParamStatus, params.Status)
	qb.AddGreaterEqualCondition(constants.ParamCreatedTime, params.CreatedAt)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrInvalidQueryParams
	}

	var referrals []domain.Referral
	getReferralsByParamsQuery := selectReferralBaseQuery + conditions
	err := r.dbClient.DBSelect(ctx, &referrals, getReferralsByParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	result := domain.Referrals(referrals)
	return &result, nil
}
