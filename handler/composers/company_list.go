package composers

import (
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type Company struct {
	ID              int64  `json:"id"`
	NormalizedName  string `json:"normalized_name"`
	DisplayName     string `json:"display_name"`
	SmallLogo       string `json:"small_logo"`
	BigLogo         string `json:"big_logo"`
	OfficialWebsite string `json:"official_website,omitempty"`
	CareersPage     string `json:"careers_page,omitempty"`
	Priority        int64  `json:"priority"`
	Verified        bool   `json:"verified"`
	Popular         bool   `json:"popular"`
}

func GetPopularCompanies(companies domain.Companies, limit int) []Company {
	popularCompanies := make([]Company, 0)
	for _, c := range companies {
		if !utils.FromPtr(c.Popular) {
			continue
		}

		if len(popularCompanies) >= limit {
			break
		}

		popularCompanies = append(popularCompanies, getCompany(c))
	}
	return popularCompanies
}

func GetAllCompanies(companies domain.Companies, limit int) []Company {
	allCompanies := make([]Company, 0)
	for _, c := range companies {
		if len(allCompanies) >= limit {
			break
		}

		allCompanies = append(allCompanies, getCompany(c))
	}
	return allCompanies
}

func getCompany(c domain.Company) Company {
	return Company{
		ID:              c.ID,
		NormalizedName:  c.NormalizedName,
		DisplayName:     c.DisplayName,
		SmallLogo:       c.SmallLogo,
		BigLogo:         c.BigLogo,
		OfficialWebsite: c.OfficialWebsite,
		CareersPage:     c.CareersPage,
		Priority:        c.GetPriority(),
		Verified:        utils.FromPtr(c.Verified),
		Popular:         utils.FromPtr(c.Popular),
	}
}
