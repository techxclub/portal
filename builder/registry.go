package builder

import (
	"github.com/techx/portal/builder/repository"
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	MessageBuilder      MessageBuilder
	UsersRepository     repository.UsersRepository
	CompaniesRepository repository.CompaniesRepository
	ReferralsRepository repository.ReferralsRepository
	MailBuilder         MailBuilder
}

func NewRegistry(cfg config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		MessageBuilder:      NewMessageBuilder(cfg, clientRegistry.Twilio),
		UsersRepository:     repository.NewUsersRepository(clientRegistry.DB),
		CompaniesRepository: repository.NewCompaniesRepository(clientRegistry.DB),
		ReferralsRepository: repository.NewReferralsRepository(clientRegistry.DB),
		MailBuilder:         NewMailBuilder(cfg, clientRegistry.GMail),
	}
}
