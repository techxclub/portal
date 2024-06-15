package domain

type Companies []Company

type Company struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
