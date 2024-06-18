package builder

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
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

	err = mb.optCache.Set(ctx, params.Value, otp, time.Minute*5)
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

	opt, err := mb.optCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get opt from cache")
		return domain.AuthInfo{}, err
	}

	if opt != params.OTP {
		return domain.AuthInfo{
			Status: constants.AuthStatusFailed,
		}, nil
	}

	err = mb.optCache.Del(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete opt from cache")
	}

	return domain.AuthInfo{
		Status: constants.AuthStatusVerified,
	}, nil
}

func (mb messageBuilder) resendEmailOTPViaGomail(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	otp, err := mb.optCache.Get(ctx, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get opt from cache")
		return domain.AuthInfo{}, err
	}

	if otp == "" {
		newOTP := generateOTP()

		err := mb.sendMail(ctx, newOTP, params.Value)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to send opt")
			return domain.AuthInfo{}, err
		}

		err = mb.optCache.Set(ctx, params.Value, newOTP, time.Minute*5)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to set opt in cache")
			return domain.AuthInfo{}, err
		}

		return domain.AuthInfo{
			OTP:    &newOTP,
			Status: constants.AuthStatusResent,
		}, nil

	}

	err = mb.sendMail(ctx, otp, params.Value)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send opt")
		return domain.AuthInfo{}, err
	}

	return domain.AuthInfo{
		OTP:    &otp,
		Status: constants.AuthStatusResent,
	}, nil
}

func (mb messageBuilder) sendMail(ctx context.Context, otp, email string) error {
	i18nValues := map[string]interface{}{
		"OTP": otp,
	}

	subject := i18n.Translate(ctx, i18nKeyEmailOTPMailSubject)
	bodyHTML := i18n.Translate(ctx, i18nKeyEmailOTPMailBody, i18nValues)

	from := mb.cfg.GMail.From
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyHTML)

	err := mb.otpMailClient.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
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
