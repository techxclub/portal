package composers

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/i18n"
)

const (
	templateCTAEnabled  = "cta_enabled_gold"
	templateCTADisabled = "cta_disabled_gray"

	i18nKeyGetReferralCTA     = "get_referral_cta"
	i18nKeyPendingReferralCTA = "pending_referral_cta"
)

type CompanyUser struct {
	Name              string      `json:"name"`
	UserUUID          string      `json:"user_uuid"`
	CompanyName       string      `json:"company_name"`
	Designation       string      `json:"designation"`
	YearsOfExperience float32     `json:"years_of_experience"`
	ReferralCTA       ReferralCTA `json:"referral_cta"`
}

type ReferralCTA struct {
	Template string `json:"template"`
	Enabled  bool   `json:"enabled"`
	Text     string `json:"text"`
}

func GetCompanyUsers(ctx context.Context, companyUsers domain.CompanyUsersService) []CompanyUser {
	referralReceivedMap := make(map[string]bool)

	for _, referral := range *companyUsers.Referrals {
		referralReceivedMap[referral.ProviderUserUUID] = true
	}

	companyUsersList := make([]CompanyUser, 0)
	for _, user := range *companyUsers.Users {
		companyUsersList = append(companyUsersList, CompanyUser{
			Name:              user.Name,
			UserUUID:          user.UserUUID,
			CompanyName:       user.CompanyName,
			Designation:       user.Designation,
			YearsOfExperience: user.YearsOfExperience,
			ReferralCTA:       getReferralCTA(ctx, referralReceivedMap[user.UserUUID]),
		})
	}

	return companyUsersList
}

func getReferralCTA(ctx context.Context, referralExist bool) ReferralCTA {
	if referralExist {
		return ReferralCTA{
			Template: templateCTADisabled,
			Enabled:  false,
			Text:     i18n.Title(ctx, i18nKeyPendingReferralCTA),
		}
	}

	return ReferralCTA{
		Template: templateCTAEnabled,
		Enabled:  true,
		Text:     i18n.Title(ctx, i18nKeyGetReferralCTA),
	}
}
