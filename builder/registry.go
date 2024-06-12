package builder

import (
	"github.com/techx/portal/builder/repository"
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	MessageBuilder MessageBuilder
	UsersRepo      repository.UsersRepo
	ReferralsRepo  repository.ReferralsRepo
	MailBuilder    MailBuilder
}

func NewRegistry(cfg config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		MessageBuilder: NewMessageBuilder(cfg, clientRegistry.Twilio),
		UsersRepo:      repository.NewUsersRepo(clientRegistry.DB),
		ReferralsRepo:  repository.NewReferralsRepo(clientRegistry.DB),
		MailBuilder:    NewMailBuilder(cfg, clientRegistry.GMail),
	}
}
