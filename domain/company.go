package domain

type Companies []Company

type Company struct {
	CompanyID int64  `db:"company_id"`
	Name      string `db:"name"`
}
