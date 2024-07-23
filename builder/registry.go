package builder

import (
	"github.com/techx/portal/builder/repository"
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	UsersRepository     repository.UsersRepository
	CompaniesRepository repository.CompaniesRepository
	ReferralsRepository repository.ReferralsRepository
	GoogleOAuthBuilder  GoogleOAuthBuilder
	OTPBuilder          OTPBuilder
	ReferralMailBuilder ReferralMailBuilder
}

func NewRegistry(cfg *config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		UsersRepository:     repository.NewUsersRepository(clientRegistry.DB),
		CompaniesRepository: repository.NewCompaniesRepository(clientRegistry.DB),
		ReferralsRepository: repository.NewReferralsRepository(clientRegistry.DB),
		OTPBuilder:          NewOTPBuilder(cfg, clientRegistry.OTPMailClient, clientRegistry.OTPCache),
		ReferralMailBuilder: NewReferralMailBuilder(cfg, clientRegistry.ReferralMailClient),
		GoogleOAuthBuilder:  NewGoogleOAuthBuilder(cfg.GoogleAuth, clientRegistry.GoogleClient),
	}
}
