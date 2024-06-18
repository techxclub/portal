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
	SendOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
	ResendOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
	VerifyOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
}

type messageBuilder struct {
	cfg           *config.Config
	otpMailClient *gomail.Dialer
	optCache      cache.Cache[string]
}

func NewOTPBuilder(cfg *config.Config, otpMailClient *gomail.Dialer, otpCache cache.Cache[string]) OTPBuilder {
	return &messageBuilder{
		cfg:           cfg,
		otpMailClient: otpMailClient,
		optCache:      otpCache,
	}
}

func (mb messageBuilder) SendOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.AuthChannelSMS:
		return mb.sendMobileOTP(ctx, params)
	case constants.AuthChannelEmail:
		return mb.sendEmailOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) ResendOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.AuthChannelSMS:
		return mb.resendMobileOTP(ctx, params)
	case constants.AuthChannelEmail:
		return mb.resendEmailOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) VerifyOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	if mb.cfg.OTP.MockingEnabled {
		return mb.mockVerifyOTP(ctx, params)
	}

	switch params.Channel {
	case constants.AuthChannelEmail:
		return mb.verifyEmailOTP(ctx, params)
	case constants.AuthChannelSMS:
		return mb.verifyMobileOTP(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidAuthChannel
	}
}

func (mb messageBuilder) sendMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.SMSThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.sendMobileOTPViaMsg91(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidSMSProvider
	}
}

func (mb messageBuilder) resendMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.SMSThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.resendMobileOTPViaMsg91(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidSMSProvider
	}
}

func (mb messageBuilder) verifyMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.SMSThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.verifyMobileOTPViaMsg91(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidSMSProvider
	}
}

func (mb messageBuilder) sendEmailOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.sendEmailOTPViaMsg91(ctx, params)
	case constants.ThirdPartyGomail:
		return mb.sendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) resendEmailOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.resendEmailOTPViaMsg91(ctx, params)
	case constants.ThirdPartyGomail:
		return mb.resendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) verifyEmailOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyMsg91:
		return mb.verifyEmailOTPViaMsg91(ctx, params)
	case constants.ThirdPartyGomail:
		return mb.verifyEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthInfo{}, errors.ErrInvalidEmailProvider
	}
}

func (mb messageBuilder) mockSendOTP(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	return domain.AuthInfo{Status: constants.AuthStatusPending}, nil
}

func (mb messageBuilder) mockVerifyOTP(_ context.Context, req domain.AuthRequest) (domain.AuthInfo, error) {
	if req.OTP == "123456" {
		return domain.AuthInfo{Status: constants.AuthStatusVerified}, nil
	}

	return domain.AuthInfo{Status: constants.AuthStatusPending}, nil
}
