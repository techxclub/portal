package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

const (
	insertReferralQuery          = `INSERT INTO referrals (create_time, requester_user_id, provider_user_id, company_id, job_link, status) VALUES (:create_time, :requester_user_id, :provider_user_id, :company_id, :job_link, :status) RETURNING id`
	fetchReferralSelectorFields  = `id, requester_user_id, provider_user_id, company_id, job_link, status, create_time`
	selectReferralBaseQuery      = `SELECT ` + fetchReferralSelectorFields + ` FROM referrals WHERE `
	namedUpdateReferralBaseQuery = `UPDATE referrals SET %s WHERE id=:id`
)

type ReferralsRepository interface {
	InsertReferral(ctx context.Context, referral domain.Referral) (*domain.Referral, error)
	UpdateReferral(ctx context.Context, referral *domain.Referral) error
	FetchReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error)
	ExpirePendingReferrals(ctx context.Context, referral *domain.Referral) error
}

type referralsRepository struct {
	dbClient db.Client
}

type ReferralsReturning struct {
	ID int64 `db:"id"`
}

func NewReferralsRepository(dbClient db.Client) ReferralsRepository {
	return &referralsRepository{
		dbClient: dbClient,
	}
}

func (r referralsRepository) InsertReferral(ctx context.Context, referral domain.Referral) (*domain.Referral, error) {
	var returning ReferralsReturning
	now := time.Now()
	referral.CreatedAt = &now
	referral.UpdatedAt = &now

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertReferralQuery, map[string]interface{}{
			constants.ParamRequesterID: referral.RequesterUserUUID,
			constants.ParamProviderID:  referral.ProviderUserUUID,
			constants.ParamCompanyID:   referral.CompanyID,
			constants.ParamJobLink:     referral.JobLink,
			constants.ParamStatus:      referral.Status,
			constants.ParamCreateTime:  referral.CreatedAt,
			constants.ParamUpdateTime:  referral.UpdatedAt,
		})
	})
	if err != nil {
		return nil, err
	}

	return &referral, nil
}

func (r referralsRepository) UpdateReferral(ctx context.Context, referral *domain.Referral) error {
	nqb := domain.NewSetQueryBuilder()
	nqb.AddEqualCondition(constants.ParamID, referral.ID)
	nqb.AddEqualCondition(constants.ParamRequesterID, referral.RequesterUserUUID)
	nqb.AddEqualCondition(constants.ParamProviderID, referral.ProviderUserUUID)
	nqb.AddEqualCondition(constants.ParamCompanyID, referral.CompanyID)
	nqb.AddEqualCondition(constants.ParamJobLink, referral.JobLink)
	nqb.AddEqualCondition(constants.ParamStatus, referral.Status)
	namedParams, namedArgs := nqb.BuildNamed()

	updateReferralQuery := fmt.Sprintf(namedUpdateReferralBaseQuery, namedParams)
	namedArgs[constants.ParamID] = referral.ID

	return r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateReferralQuery, namedArgs)
	})
}

func (r referralsRepository) FetchReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamRequesterID, params.RequesterUserUUID)
	qb.AddEqualCondition(constants.ParamProviderID, params.ProviderUserUUID)
	qb.AddEqualCondition(constants.ParamCompanyID, params.CompanyID)
	qb.AddEqualCondition(constants.ParamStatus, params.Status)
	qb.AddGreaterEqualCondition(constants.ParamCreateTime, params.CreatedAt)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrInvalidQueryParams
	}

	var referrals []domain.Referral
	getReferralsByParamsQuery := selectReferralBaseQuery + conditions
	err := r.dbClient.DBSelect(ctx, &referrals, getReferralsByParamsQuery, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrReferralNotFound
	}

	result := domain.Referrals(referrals)
	return &result, err
}

func (r referralsRepository) ExpirePendingReferrals(ctx context.Context, _ *domain.Referral) error {
	expirePendingReferralsQuery := `UPDATE referrals SET status='EXPIRED' WHERE status='PENDING' AND create_time <= NOW() - INTERVAL '7 days'`
	return r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBExecInTx(ctx, tx, expirePendingReferralsQuery)
	})
}
