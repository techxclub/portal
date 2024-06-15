package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/domain"
)

const (
	insertCompanyQuery = `INSERT INTO companies (id, name, small_logo, big_logo, official_website, careers_page, priority, verified, popular) VALUES (:id, :name, :small_logo, :big_logo, :official_website, :careers_page, :priority, :verified, :popular) RETURNING id`
)

type CompaniesRepository interface {
	AddCompany(ctx context.Context, details domain.Company) (*domain.Company, error)
	UpdateCompany(ctx context.Context, details domain.Company) (*domain.Company, error)
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
			"name":             details.Name,
			"small_logo":       details.SmallLogo,
			"big_logo":         details.BigLogo,
			"official_website": details.OfficialWebsite,
			"careers_page":     details.CareersPage,
			"priority":         details.Priority,
			"verified":         details.Verified,
			"popular":          details.Popular,
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

func (r companiesRepository) UpdateCompany(ctx context.Context, details domain.Company) (*domain.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (r companiesRepository) GetCompanyForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (r companiesRepository) GetCompaniesForParams(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	//TODO implement me
	panic("implement me")
}
