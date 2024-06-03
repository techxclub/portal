package builder

import (
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct{}

func NewRegistry(_ config.Config, _ *client.Registry) *Registry {
	return &Registry{}
}
