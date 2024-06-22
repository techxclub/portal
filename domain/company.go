package domain

import (
	"math"
)

type Companies []Company

type Company struct {
	ID              int64  `db:"id"`
	Actor           string `db:"actor"`
	NormalizedName  string `db:"normalized_name"`
	DisplayName     string `db:"display_name"`
	SmallLogo       string `db:"small_logo"`
	BigLogo         string `db:"big_logo"`
	OfficialWebsite string `db:"official_website"`
	CareersPage     string `db:"careers_page"`
	Priority        int64  `db:"priority"`
	Verified        *bool  `db:"verified"`
	Popular         *bool  `db:"popular"`
}

type FetchCompanyParams struct {
	ID             int64
	NormalizedName string
	DisplayName    string
	Priority       int64
	Verified       *bool
	Popular        *bool
}

func (c Company) GetPriority() int64 {
	if c.Priority == 0 {
		return math.MaxInt8
	}

	return c.Priority
}
