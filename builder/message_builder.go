package builder

import (
	"context"

	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type MessageBuilder interface {
	SendMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
	VerifyMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
}

type messageBuilder struct {
	cfg          config.Config
	twilioClient twilio.Client
}

func NewMessageBuilder(twilioClient twilio.Client) MessageBuilder {
	return &messageBuilder{
		cfg:          config.Config{},
		twilioClient: twilioClient,
	}
}

func (mb messageBuilder) SendMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.ThirdPartySmsProvider {
	case constants.ThirdPartyTwilio:
		return mb.sendMobileOTPViaTwilio(ctx, params)
	case constants.ThirdPartMsg91:
		return mb.sendMobileOTPViaMsg91(ctx, params)
	default:
		return mb.sendMobileOTPViaTwilio(ctx, params)
	}
}

func (mb messageBuilder) VerifyMobileOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	switch mb.cfg.ThirdPartySmsProvider {
	case constants.ThirdPartyTwilio:
		return mb.verifyMobileOTPViaTwilio(ctx, params)
	case constants.ThirdPartMsg91:
		return mb.verifyMobileOTPViaMsg91(ctx, params)
	default:
		return mb.verifyMobileOTPViaMsg91(ctx, params)
	}
}

func (mb messageBuilder) sendMobileOTPViaTwilio(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	createOTPRequest := twilio.NewCreateVerificationRequest(params.Value, params.Channel)
	resp, err := mb.twilioClient.SendOTP(ctx, createOTPRequest)
	return domain.AuthInfo{Status: resp.Status}, err
}

func (mb messageBuilder) verifyMobileOTPViaTwilio(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	verifyOTPRequest := twilio.NewCheckVerificationRequest(params.Value, params.OTP)
	resp, err := mb.twilioClient.VerifyOTP(ctx, verifyOTPRequest)
	return domain.AuthInfo{Status: resp.Status}, err
}

func (mb messageBuilder) sendMobileOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}

func (mb messageBuilder) verifyMobileOTPViaMsg91(_ context.Context, _ domain.AuthRequest) (domain.AuthInfo, error) {
	panic("implement me")
}
