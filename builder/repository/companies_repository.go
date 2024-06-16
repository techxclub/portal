package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

const (
	insertCompanyQuery          = `INSERT INTO companies (name, small_logo, big_logo, official_website, careers_page, priority, verified, popular) VALUES (:name, :small_logo, :big_logo, :official_website, :careers_page, :priority, :verified, :popular) RETURNING id`
	namedUpdateCompanyBaseQuery = `UPDATE companies SET %s WHERE id=:id`
	getCompanySelectorFields    = `id, name, small_logo, big_logo, official_website, careers_page, priority, verified, popular`
	selectCompanyBaseQuery      = `SELECT ` + getCompanySelectorFields + ` FROM companies WHERE `
)

type CompaniesRepository interface {
	AddCompany(ctx context.Context, details domain.Company) (*domain.Company, error)
	UpdateCompany(ctx context.Context, details *domain.Company) error
	GetCompanyForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Company, error)
	GetCompaniesForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error)
}

type CompaniesReturning struct {
	ID int64 `db:"id"`
}

type companiesRepository struct {
	dbClient *db.Repository
}

func NewCompaniesRepository(userDB *db.Repository) CompaniesRepository {
	return &companiesRepository{
		dbClient: userDB,
	}
}

func (r companiesRepository) AddCompany(ctx context.Context, details domain.Company) (*domain.Company, error) {
	var returning CompaniesReturning
	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertCompanyQuery, map[string]interface{}{
			constants.ParamName:            details.Name,
			constants.ParamActor:           details.Actor,
			constants.ParamSmallLogo:       details.SmallLogo,
			constants.ParamBigLogo:         details.BigLogo,
			constants.ParamOfficialWebsite: details.OfficialWebsite,
			constants.ParamCareersPage:     details.CareersPage,
			constants.ParamPriority:        details.Priority,
			constants.ParamVerified:        details.Verified,
			constants.ParamPopular:         details.Popular,
		})
	})
	if err != nil {
		return nil, err
	}

	company := domain.Company{
		ID:              returning.ID,
		Name:            details.Name,
		SmallLogo:       details.SmallLogo,
		BigLogo:         details.BigLogo,
		OfficialWebsite: details.OfficialWebsite,
		CareersPage:     details.CareersPage,
		Priority:        details.Priority,
		Verified:        details.Verified,
		Popular:         details.Popular,
	}

	return &company, nil
}

func (r companiesRepository) UpdateCompany(ctx context.Context, details *domain.Company) error {
	nqb := domain.NewNamedQueryBuilder()
	nqb.AddEqualCondition(constants.ParamName, details.Name)
	nqb.AddEqualCondition(constants.ParamSmallLogo, details.SmallLogo)
	nqb.AddEqualCondition(constants.ParamBigLogo, details.BigLogo)
	nqb.AddEqualCondition(constants.ParamOfficialWebsite, details.OfficialWebsite)
	nqb.AddEqualCondition(constants.ParamCareersPage, details.CareersPage)
	if details.Priority != nil {
		nqb.AddEqualCondition(constants.ParamPriority, *details.Priority)
	}
	if details.Verified != nil {
		nqb.AddEqualCondition(constants.ParamVerified, *details.Verified)
	}
	if details.Popular != nil {
		nqb.AddEqualCondition(constants.ParamPopular, *details.Popular)
	}
	namedParams, namedArgs := nqb.BuildNamedConditions()

	updateCompanyQuery := fmt.Sprintf(namedUpdateCompanyBaseQuery, namedParams)
	namedArgs[constants.ParamID] = details.ID

	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r companiesRepository) GetCompanyForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Company, error) {
	qb := domain.NewQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamName, params.Name)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrInvalidQueryParams
	}

	var company domain.Company
	getCompanyForParamsQuery := selectCompanyBaseQuery + conditions
	err := r.dbClient.DBGet(ctx, &company, getCompanyForParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r companiesRepository) GetCompaniesForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	qb := domain.NewQueryBuilder()
	if params.ID != nil {
		qb.AddEqualCondition(constants.ParamID, *params.ID)
	}
	qb.AddEqualCondition(constants.ParamName, params.Name)
	if params.Priority != nil {
		qb.AddEqualCondition(constants.ParamPriority, *params.Priority)
	}
	if params.Verified != nil {
		qb.AddEqualCondition(constants.ParamVerified, *params.Verified)
	}
	if params.Popular != nil {
		qb.AddEqualCondition(constants.ParamPopular, *params.Popular)
	}

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrInvalidQueryParams
	}

	var companies []domain.Company
	getCompanyForParamsQuery := selectCompanyBaseQuery + conditions
	err := r.dbClient.DBSelect(ctx, &companies, getCompanyForParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	if len(companies) == 0 {
		return nil, errors.ErrNoDataFound
	}

	result := domain.Companies(companies)
	return &result, nil
}
