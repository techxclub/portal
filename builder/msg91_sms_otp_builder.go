package builder

import (
	"context"

	"github.com/techx/portal/domain"
)

func (mb messageBuilder) sendMobileOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}

func (mb messageBuilder) resendMobileOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}

func (mb messageBuilder) verifyMobileOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}
