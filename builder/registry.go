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
	MailBuilder         MailBuilder
	GoogleOAuthBuilder  GoogleOAuthBuilder
	OTPBuilder          OTPBuilder
}

func NewRegistry(cfg *config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		UsersRepository:     repository.NewUsersRepository(clientRegistry.DB),
		CompaniesRepository: repository.NewCompaniesRepository(clientRegistry.DB),
		ReferralsRepository: repository.NewReferralsRepository(clientRegistry.DB),
		MailBuilder:         NewMailBuilder(clientRegistry.ServiceMailClient, clientRegistry.SupportMailClient),
		OTPBuilder:          NewOTPBuilder(cfg, clientRegistry.SupportMailClient, clientRegistry.OTPCache),
		GoogleOAuthBuilder:  NewGoogleOAuthBuilder(cfg.GoogleAuth, clientRegistry.GoogleClient),
	}
}
