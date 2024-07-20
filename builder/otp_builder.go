package builder

import (
	"context"

	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"gopkg.in/gomail.v2"
)

type OTPBuilder interface {
	SendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error)
	ResendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error)
	VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error)
}

type messageBuilder struct {
	cfg           *config.Config
	otpMailClient *gomail.Dialer
	otpCache      cache.Cache[cache.OTPCache]
}

func NewOTPBuilder(cfg *config.Config, otpMailClient *gomail.Dialer, otpCache cache.Cache[cache.OTPCache]) OTPBuilder {
	return &messageBuilder{
		cfg:           cfg,
		otpMailClient: otpMailClient,
		otpCache:      otpCache,
	}
}

func (mb messageBuilder) SendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return mb.sendEmailOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) ResendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return mb.resendEmailOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockVerifyOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return mb.verifyEmailOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) sendEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return mb.sendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) resendEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return mb.resendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) verifyEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return mb.verifyEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) mockSendOTP(_ context.Context, _ domain.OTPRequest) (domain.AuthInfo, error) {
	return domain.AuthInfo{Status: constants.OTPStatusPending}, nil
}

func (mb messageBuilder) mockVerifyOTP(_ context.Context, req domain.OTPRequest) (domain.AuthInfo, error) {
	if req.OTP == "123456" {
		return domain.AuthInfo{Status: constants.OTPStatusVerified}, nil
	}

	return domain.AuthInfo{Status: constants.OTPStatusPending}, nil
}
