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
}

func NewRegistry(_ config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		MessageBuilder: NewMessageBuilder(clientRegistry.Twilio),
		UsersRepo:      repository.NewUsersRepo(clientRegistry.DB),
		ReferralsRepo:  repository.NewReferralsRepo(clientRegistry.DB),
	}
}
