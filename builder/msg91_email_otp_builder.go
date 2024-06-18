package builder

import (
	"context"

	"github.com/techx/portal/domain"
)

func (mb messageBuilder) sendEmailOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}

func (mb messageBuilder) verifyEmailOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}

func (mb messageBuilder) resendEmailOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}
