package builder

import (
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	AuthBuilder     AuthBuilder
	UserInfoBuilder UserInfoBuilder
}

func NewRegistry(_ config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		AuthBuilder:     NewAuthBuilder(clientRegistry.Twilio),
		UserInfoBuilder: NewUsersInfoBuilder(clientRegistry.UsersDB),
	}
}
