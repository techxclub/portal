package builder

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/i18n"
	"gopkg.in/gomail.v2"
)

const (
	defaultOTP = "972635"

	i18nKeyEmailOTPMailSubject = "email_otp_mail_subject"
	i18nKeyEmailOTPMailBody    = "email_otp_mail_body"
)

func (ob otpBuilder) sendEmailOTPViaGomail(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	otp := generateOTP()

	err := ob.sendMail(ctx, otp, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send opt")
		return domain.AuthDetails{}, err
	}

	ttl := ob.cfg.OTP.TTL
	otpCacheValue := &cache.OTPCache{
		OTP:      otp,
		Attempts: 0,
	}
	err = ob.otpCache.Set(ctx, params.Value, otpCacheValue, ttl)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set opt in cache")
		return domain.AuthDetails{}, err
	}

	return domain.AuthDetails{
		Status: constants.OTPStatusPending,
	}, nil
}

func (ob otpBuilder) verifyEmailOTPViaGomail(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
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

func (ob otpBuilder) resendEmailOTPViaGomail(ctx context.Context, params domain.OTPRequest) (domain.AuthDetails, error) {
	otpCacheValue, err := ob.otpCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Existing OTP not found in cache, sending new otp")
		return ob.sendEmailOTPViaGomail(ctx, params)
	}

	err = ob.sendMail(ctx, otpCacheValue.OTP, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send opt")
		return domain.AuthDetails{}, err
	}

	return domain.AuthDetails{
		Status: constants.OTPStatusPending,
	}, nil
}

func (ob otpBuilder) sendMail(ctx context.Context, otp, email string) error {
	i18nValues := map[string]interface{}{
		"OTP": otp,
	}

	subject := i18n.Translate(ctx, i18nKeyEmailOTPMailSubject)
	bodyHTML := i18n.Translate(ctx, i18nKeyEmailOTPMailBody, i18nValues)
	mailCfg := ob.cfg.OTPMail
	messageID := mailCfg.GetMessageID()

	m := gomail.NewMessage()
	m.SetHeader(constants.GomailHeaderFrom, mailCfg.GetFrom())
	m.SetHeader(constants.GomailHeaderTo, email)
	m.SetHeader(constants.GomailHeaderSubject, subject)
	m.SetHeader(constants.GomailHeaderMessageID, messageID)
	m.SetHeader(constants.GomailHeaderInReplyTo, messageID)
	m.SetHeader(constants.GomailHeaderReferences, messageID)
	m.SetBody(constants.GomailContentTypeHTML, bodyHTML)

	err := ob.otpMailClient.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
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
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate random number")
		return defaultOTP
	}

	otp := randomNumber.Int64() + 100000
	return strconv.FormatInt(otp, 10)
}
