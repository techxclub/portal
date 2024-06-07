package builder

import (
	"context"

	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/domain"
)

type UserAuthBuilder interface {
	GenerateOTP(ctx context.Context, params domain.OTPGeneration) error
}

type userAuthBuilder struct {
	twilioClient twilio.Client
}

func NewUserAuthBuilder(twilioClient twilio.Client) UserAuthBuilder {
	return &userAuthBuilder{
		twilioClient: twilioClient,
	}
}

func (u userAuthBuilder) GenerateOTP(ctx context.Context, params domain.OTPGeneration) error {
	return u.twilioClient.SendOTP(ctx, params.Value, params.Type)
}
