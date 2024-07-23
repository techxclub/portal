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
	SendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error)
	ResendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error)
	VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error)
	Check(ctx context.Context, email string) bool
}

type otpBuilder struct {
	cfg           *config.Config
	otpMailClient *gomail.Dialer
	otpCache      cache.Cache[cache.OTPCache]
}

func NewOTPBuilder(cfg *config.Config, otpMailClient *gomail.Dialer, otpCache cache.Cache[cache.OTPCache]) OTPBuilder {
	return &otpBuilder{
		cfg:           cfg,
		otpMailClient: otpMailClient,
		otpCache:      otpCache,
	}
}

func (ob otpBuilder) SendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	if ob.cfg.OTP.MockingEnabled {
		return ob.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return ob.sendEmailOTP(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidAuthChannel
	}
}

func (ob otpBuilder) ResendOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	if ob.cfg.OTP.MockingEnabled {
		return ob.mockSendOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return ob.resendEmailOTP(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidAuthChannel
	}
}

func (ob otpBuilder) VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	if ob.cfg.OTP.MockingEnabled {
		return ob.mockVerifyOTP(ctx, params)
	}

	switch params.Channel {
	case constants.OTPChannelEmail:
		return ob.verifyEmailOTP(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidAuthChannel
	}
}

func (ob otpBuilder) Check(ctx context.Context, email string) bool {
	if ob.cfg.OTP.MockingEnabled {
		return true
	}

	val, err := ob.otpCache.Get(ctx, email)
	if err != nil {
		return false
	}

	return val.Verified
}

func (ob otpBuilder) sendEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	switch ob.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return ob.sendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidEmailProvider
	}
}

func (ob otpBuilder) resendEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	switch ob.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return ob.resendEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidEmailProvider
	}
}

func (ob otpBuilder) verifyEmailOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	switch ob.cfg.OTP.EmailThirdPartyProvider {
	case constants.ThirdPartyGomail:
		return ob.verifyEmailOTPViaGomail(ctx, params)
	default:
		return domain.AuthDetails{}, errors.ErrInvalidEmailProvider
	}
}

func (ob otpBuilder) mockSendOTP(_ context.Context, _ domain.OTPRequest) (domain.AuthDetails, error) {
	return domain.AuthDetails{Status: constants.OTPStatusPending}, nil
}

func (ob otpBuilder) mockVerifyOTP(_ context.Context, req domain.OTPRequest) (domain.AuthDetails, error) {
	if req.OTP == "123456" {
		return domain.AuthDetails{Status: constants.OTPStatusVerified}, nil
	}

	return domain.AuthDetails{Status: constants.OTPStatusPending}, nil
}
