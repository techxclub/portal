package builder

import (
	"context"

	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/domain"
)

type AuthBuilder interface {
	GenerateOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
	VerifyOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error)
}

type authBuilder struct {
	twilioClient twilio.Client
}

func NewAuthBuilder(twilioClient twilio.Client) AuthBuilder {
	return &authBuilder{
		twilioClient: twilioClient,
	}
}

func (u authBuilder) GenerateOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	createOTPRequest := twilio.NewCreateVerificationRequest(params.Value, params.Channel)
	resp, err := u.twilioClient.SendOTP(ctx, createOTPRequest)
	return domain.AuthInfo{Status: resp.Status}, err
}

func (u authBuilder) VerifyOTP(ctx context.Context, params domain.AuthRequest) (domain.AuthInfo, error) {
	verifyOTPRequest := twilio.NewCheckVerificationRequest(params.Value, params.OTP)
	resp, err := u.twilioClient.VerifyOTP(ctx, verifyOTPRequest)
	return domain.AuthInfo{Status: resp.Status}, err
}
