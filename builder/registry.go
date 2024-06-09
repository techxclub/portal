package builder

import (
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	MessageBuilder  MessageBuilder
	UserInfoBuilder UserInfoBuilder
}

func NewRegistry(_ config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		MessageBuilder:  NewMessageBuilder(clientRegistry.Twilio),
		UserInfoBuilder: NewUsersInfoBuilder(clientRegistry.UsersDB),
	}
}
