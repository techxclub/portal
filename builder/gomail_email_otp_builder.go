package builder

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

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

func (mb messageBuilder) sendEmailOTPViaGomail(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	otp := generateOTP()

	err := mb.sendMail(ctx, otp, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send opt")
		return domain.AuthInfo{}, err
	}

	ttl := mb.cfg.OTP.TTL
	expiry := time.Now().Add(ttl)
	otpCacheValue := &cache.OTPCache{
		OTP:        otp,
		ExpiryTime: expiry,
		Attempts:   0,
	}
	err = mb.otpCache.Set(ctx, params.Value, otpCacheValue, ttl)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set opt in cache")
		return domain.AuthInfo{}, err
	}

	return domain.AuthInfo{
		OTP:    &otp,
		Status: constants.AuthStatusPending,
	}, nil
}

func (mb messageBuilder) verifyEmailOTPViaGomail(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	if params.OTP == "" {
		return domain.AuthInfo{
			Status: constants.AuthStatusFailed,
		}, errors.ErrOtpNotProvided
	}

	otpCacheValue, err := mb.otpCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get otpCacheValue from cache")
		return domain.AuthInfo{}, err
	}

	if otpCacheValue.OTP != params.OTP {
		go func() {
			mb.updateOTPAttempts(ctx, params.Value, otpCacheValue)
		}()

		return domain.AuthInfo{
			Status: constants.AuthStatusFailed,
		}, nil
	}

	err = mb.otpCache.Del(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete otpCacheValue from cache")
	}

	return domain.AuthInfo{
		Status: constants.AuthStatusVerified,
	}, nil
}

func (mb messageBuilder) resendEmailOTPViaGomail(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	otpCacheValue, err := mb.otpCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Existing OTP not found in cache, sending new otp")
		return mb.sendEmailOTPViaGomail(ctx, params)
	}

	err = mb.sendMail(ctx, otpCacheValue.OTP, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send opt")
		return domain.AuthInfo{}, err
	}

	return domain.AuthInfo{
		OTP:    &otpCacheValue.OTP,
		Status: constants.AuthStatusResent,
	}, nil
}

func (mb messageBuilder) sendMail(ctx context.Context, otp, email string) error {
	i18nValues := map[string]interface{}{
		"OTP": otp,
	}

	subject := i18n.Translate(ctx, i18nKeyEmailOTPMailSubject)
	bodyHTML := i18n.Translate(ctx, i18nKeyEmailOTPMailBody, i18nValues)

	m := gomail.NewMessage()
	m.SetHeader("From", mb.cfg.OTPMail.GetFrom())
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyHTML)

	err := mb.otpMailClient.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

func (mb messageBuilder) updateOTPAttempts(ctx context.Context, email string, otpCacheValue *cache.OTPCache) {
	otpCacheValue.Attempts++
	err := mb.otpCache.Set(ctx, email, otpCacheValue, mb.cfg.OTP.TTL)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to update attempts in cache")
	}
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
