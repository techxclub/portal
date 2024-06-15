package domain

type Companies []Company

type Company struct {
	ID              int64  `db:"id"`
	Name            string `db:"name"`
	SmallLogo       string `db:"small_logo"`
	BigLogo         string `db:"big_logo"`
	OfficialWebsite string `db:"official_website"`
	CareersPage     string `db:"careers_page"`
	Priority        int64  `db:"priority"`
	Verified        bool   `db:"verified"`
	Popular         bool   `db:"popular"`
}

type FetchCompanyParams struct {
	ID       string
	Name     string
	Verified bool
}

type UpdateCompanyParams struct {
	Name            string
	SmallLogo       string
	BigLogo         string
	OfficialWebsite string
	CareersPage     string
	Priority        int64
	Verified        bool
	Popular         bool
}

func (p UpdateCompanyParams) GetQueryConditions() (string, []interface{}) {
	qb := NewNamedQueryBuilder()
	qb.AddEqualCondition("name", p.Name)
	qb.AddEqualCondition("small_logo", p.SmallLogo)
	qb.AddEqualCondition("big_logo", p.BigLogo)
	qb.AddEqualCondition("official_website", p.OfficialWebsite)
	qb.AddEqualCondition("careers_page", p.CareersPage)
	qb.AddEqualCondition("priority", p.Priority)
	qb.AddEqualCondition("verified", p.Verified)
	qb.AddEqualCondition("popular", p.Popular)

	return qb.Build()
}
