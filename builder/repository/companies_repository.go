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
	insertCompanyQuery          = `INSERT INTO companies (normalized_name, display_name, small_logo, big_logo, official_website, careers_page, priority, verified, popular) VALUES (:normalized_name, :display_name, :small_logo, :big_logo, :official_website, :careers_page, :priority, :verified, :popular) RETURNING id`
	namedUpdateCompanyBaseQuery = `UPDATE companies SET %s WHERE id=:id`
	getCompanySelectorFields    = `id, normalized_name, display_name, small_logo, big_logo, official_website, careers_page, priority, verified, popular`
	selectCompanyBaseQuery      = `SELECT ` + getCompanySelectorFields + ` FROM companies WHERE `
)

type CompaniesRepository interface {
	InsertCompany(ctx context.Context, details domain.Company) (*domain.Company, error)
	UpdateCompany(ctx context.Context, details *domain.Company) error
	FetchCompanyForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Company, error)
	FetchCompaniesForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error)
}

type CompaniesReturning struct {
	ID int64 `db:"id"`
}

type companiesRepository struct {
	dbClient db.Client
}

func NewCompaniesRepository(userDB db.Client) CompaniesRepository {
	return &companiesRepository{
		dbClient: userDB,
	}
}

func (r companiesRepository) InsertCompany(ctx context.Context, details domain.Company) (*domain.Company, error) {
	var returning CompaniesReturning
	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertCompanyQuery, map[string]interface{}{
			constants.ParamNormalizedName:  details.NormalizedName,
			constants.ParamDisplayName:     details.DisplayName,
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
		NormalizedName:  details.NormalizedName,
		DisplayName:     details.DisplayName,
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
	nqb := domain.NewSetQueryBuilder()
	nqb.AddEqualCondition(constants.ParamNormalizedName, details.NormalizedName)
	nqb.AddEqualCondition(constants.ParamDisplayName, details.DisplayName)
	nqb.AddEqualCondition(constants.ParamSmallLogo, details.SmallLogo)
	nqb.AddEqualCondition(constants.ParamBigLogo, details.BigLogo)
	nqb.AddEqualCondition(constants.ParamOfficialWebsite, details.OfficialWebsite)
	nqb.AddEqualCondition(constants.ParamCareersPage, details.CareersPage)
	nqb.AddEqualCondition(constants.ParamPriority, details.Priority)
	nqb.AddEqualCondition(constants.ParamVerified, details.Verified)
	nqb.AddEqualCondition(constants.ParamPopular, details.Popular)
	namedParams, namedArgs := nqb.BuildNamed()

	updateCompanyQuery := fmt.Sprintf(namedUpdateCompanyBaseQuery, namedParams)
	namedArgs[constants.ParamID] = details.ID

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r companiesRepository) FetchCompanyForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Company, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamNormalizedName, params.NormalizedName)
	qb.AddEqualCondition(constants.ParamDisplayName, params.DisplayName)

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

func (r companiesRepository) FetchCompaniesForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamID, params.ID)
	qb.AddEqualCondition(constants.ParamNormalizedName, params.NormalizedName)
	qb.AddEqualCondition(constants.ParamDisplayName, params.DisplayName)
	qb.AddEqualCondition(constants.ParamPriority, params.Priority)
	qb.AddEqualCondition(constants.ParamVerified, params.Verified)
	qb.AddEqualCondition(constants.ParamPopular, params.Popular)

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
