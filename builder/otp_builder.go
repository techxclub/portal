package builder

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/client/email"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

const (
	otpLength  = 6
	defaultOTP = "972635"
)

type OTPBuilder interface {
	BuildOTP(ctx context.Context, params domain.OTPRequest) (string, error)
	RebuildOTP(ctx context.Context, params domain.OTPRequest) (string, error)
	VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error)
	IsOTPVerified(ctx context.Context, email string) bool
}

type otpBuilder struct {
	cfg           *config.Config
	otpMailClient email.Client
	otpCache      cache.Cache[cache.OTPCache]
}

func NewOTPBuilder(cfg *config.Config, otpMailClient email.Client, otpCache cache.Cache[cache.OTPCache]) OTPBuilder {
	return &otpBuilder{
		cfg:           cfg,
		otpMailClient: otpMailClient,
		otpCache:      otpCache,
	}
}

func (ob otpBuilder) BuildOTP(ctx context.Context, params domain.OTPRequest) (string, error) {
	otp := generateOTP()

	ttl := ob.cfg.OTP.TTL
	otpCacheValue := &cache.OTPCache{
		OTP:      otp,
		Attempts: 0,
	}

	err := ob.otpCache.Set(ctx, params.Value, otpCacheValue, ttl)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set opt in cache")
		return "", err
	}

	return otp, nil
}

func (ob otpBuilder) RebuildOTP(ctx context.Context, params domain.OTPRequest) (string, error) {
	otpCacheValue, err := ob.otpCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Existing OTP not found in cache, generating new otp")
		return ob.BuildOTP(ctx, params)
	}

	return otpCacheValue.OTP, nil
}

func (ob otpBuilder) VerifyOTP(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	if ob.cfg.OTP.MockingEnabled {
		return ob.mockVerifyOTP(ctx, params)
	}

	if params.OTP == "" {
		return domain.AuthDetails{}, errors.ErrMissingOTP
	}

	otpCacheValue, err := ob.otpCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get otpCacheValue from cache")
		return domain.AuthDetails{}, err
	}

	if otpCacheValue.OTP != params.OTP {
		ob.updateOTPAttempts(ctx, params.Value, otpCacheValue)
		return domain.AuthDetails{Status: constants.OTPStatusFailed}, nil
	}

	if err := ob.updateVerifiedOTP(ctx, params.Value, otpCacheValue); err != nil {
		return domain.AuthDetails{}, err
	}

	return domain.AuthDetails{Status: constants.OTPStatusVerified}, nil
}

func (ob otpBuilder) IsOTPVerified(ctx context.Context, email string) bool {
	if ob.cfg.OTP.MockingEnabled {
		return true
	}

	val, err := ob.otpCache.Get(ctx, email)
	if err != nil {
		return false
	}

	return val.Verified
}

func (ob otpBuilder) updateOTPAttempts(ctx context.Context, email string, otpCacheValue *cache.OTPCache) {
	otpCacheValue.Attempts++
	err := ob.otpCache.Set(ctx, email, otpCacheValue, redis.KeepTTL)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to update attempts in cache")
	}
}

func (ob otpBuilder) updateVerifiedOTP(ctx context.Context, email string, otpCacheValue *cache.OTPCache) error {
	otpCacheValue.Verified = true
	err := ob.otpCache.Set(ctx, email, otpCacheValue, ob.cfg.OTP.TTL)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to store verified status in otp cache")
	}

	return err
}

func generateOTP() string {
	otp, err := utils.GenerateRandomNumber(otpLength)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate OTP")
		return defaultOTP
	}

	return otp
}

func (ob otpBuilder) mockVerifyOTP(_ context.Context, req domain.OTPRequest) (domain.AuthDetails, error) {
	if req.OTP == "123456" {
		return domain.AuthDetails{Status: constants.OTPStatusVerified}, nil
	}

	return domain.AuthDetails{Status: constants.OTPStatusPending}, nil
}
