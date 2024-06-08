package builder

import (
	"context"

	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/domain"
)

type UserAuthBuilder interface {
	GenerateOTP(ctx context.Context, params domain.OTPGeneration) (*domain.AuthDetails, error)
	VerifyOTP(ctx context.Context, params domain.OTPVerification) (*domain.AuthDetails, error)
}

type userAuthBuilder struct {
	twilioClient twilio.Client
}

func NewUserAuthBuilder(twilioClient twilio.Client) UserAuthBuilder {
	return &userAuthBuilder{
		twilioClient: twilioClient,
	}
}

func (u userAuthBuilder) GenerateOTP(ctx context.Context, params domain.OTPGeneration) (*domain.AuthDetails, error) {
	createOTPRequest := twilio.NewCreateVerificationRequest(params.Value, params.Type)
	resp, err := u.twilioClient.SendOTP(ctx, createOTPRequest)
	return &domain.AuthDetails{Status: resp.Status}, err
}

func (u userAuthBuilder) VerifyOTP(ctx context.Context, params domain.OTPVerification) (*domain.AuthDetails, error) {
	verifyOTPRequest := twilio.NewCheckVerificationRequest(params.Value, params.Code)
	resp, err := u.twilioClient.VerifyOTP(ctx, verifyOTPRequest)
	return &domain.AuthDetails{Status: resp.Status}, err
}
