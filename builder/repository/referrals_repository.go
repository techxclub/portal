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
	insertReferralQuery          = `INSERT INTO referrals (requester_user_id, provider_user_id, company_id, job_link, status, created_time) VALUES (:requester_user_id, :provider_user_id, :company_id, :job_link, :status, :created_time) RETURNING id, created_time`
	fetchReferralSelectorFields  = `id, requester_user_id, provider_user_id, company_id, job_link, status, created_time`
	selectReferralBaseQuery      = `SELECT ` + fetchReferralSelectorFields + ` FROM referrals WHERE `
	namedUpdateReferralBaseQuery = `UPDATE referrals SET %s WHERE id=:id`
)

type ReferralsRepository interface {
	InsertReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
	FetchReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error)
	UpdateReferral(ctx context.Context, referral *domain.Referral) error
	ExpirePendingReferrals(ctx context.Context, referral *domain.Referral) error
}

type referralsRepository struct {
	dbClient db.Client
}

type ReferralsReturning struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_time"`
}

func NewReferralsRepository(dbClient db.Client) ReferralsRepository {
	return &referralsRepository{
		dbClient: dbClient,
	}
}

func (r referralsRepository) InsertReferral(ctx context.Context, params domain.ReferralParams) (*domain.Referral, error) {
	var returning ReferralsReturning
	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertReferralQuery, map[string]interface{}{
			constants.ParamRequesterID: params.RequesterUserUUID,
			constants.ParamProviderID:  params.ProviderUserUUID,
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
		ID:                returning.ID,
		RequesterUserUUID: params.RequesterUserUUID,
		ProviderUserUUID:  params.ProviderUserUUID,
		CompanyID:         params.CompanyID,
		JobLink:           params.JobLink,
		Status:            params.Status,
		CreatedAt:         &returning.CreatedAt,
	}

	return &referral, nil
}

func (r referralsRepository) FetchReferralsForParams(ctx context.Context, params domain.ReferralParams) (*domain.Referrals, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamRequesterID, params.RequesterUserUUID)
	qb.AddEqualCondition(constants.ParamProviderID, params.ProviderUserUUID)
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
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrReferralNotFound
	}

	result := domain.Referrals(referrals)
	return &result, err
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

func (r referralsRepository) ExpirePendingReferrals(ctx context.Context, _ *domain.Referral) error {
	expirePendingReferralsQuery := `UPDATE referrals SET status='EXPIRED' WHERE status='PENDING' AND created_time <= NOW() - INTERVAL '7 days'`
	return r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBExecInTx(ctx, tx, expirePendingReferralsQuery)
	})
}
