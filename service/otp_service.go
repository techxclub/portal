package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type OTPService interface {
	GenerateOTP(ctx context.Context, otpRequest domain.OTPRequest) (*domain.AuthDetails, error)
	ResendOTP(ctx context.Context, otpRequest domain.OTPRequest) (*domain.AuthDetails, error)
	VerifyOTP(ctx context.Context, otpRequest domain.OTPRequest) (*domain.AuthDetails, error)
}

type otpService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewOTPService(cfg *config.Config, registry *builder.Registry) OTPService {
	return &otpService{
		cfg:      cfg,
		registry: registry,
	}
}

func (os otpService) GenerateOTP(ctx context.Context, otpGenerationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	otp, err := os.registry.OTPBuilder.BuildOTP(ctx, otpGenerationDetail)
	if err != nil {
		return nil, err
	}

	otpMailParams := builder.OTPMailParams{
		Code:  otp,
		Email: otpGenerationDetail.Value,
	}
	refID := utils.GetRandomUUID()

	err = os.registry.MailBuilder.SendOTPMail(ctx, true, refID, otpMailParams)
	if err != nil {
		return nil, errors.ErrOTPSendFailed
	}

	return &domain.AuthDetails{Status: constants.OTPStatusPending}, nil
}

func (os otpService) ResendOTP(ctx context.Context, otpGenerationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	otp, err := os.registry.OTPBuilder.RebuildOTP(ctx, otpGenerationDetail)
	if err != nil {
		return nil, err
	}

	otpMailParams := builder.OTPMailParams{
		Code:  otp,
		Email: otpGenerationDetail.Value,
	}
	refID := utils.GetRandomUUID()

	err = os.registry.MailBuilder.SendOTPMail(ctx, true, refID, otpMailParams)
	if err != nil {
		return nil, errors.ErrOTPResendFailed
	}

	return &domain.AuthDetails{Status: constants.OTPStatusPending}, nil
}

func (os otpService) VerifyOTP(ctx context.Context, otpVerificationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	authDetails, err := os.registry.OTPBuilder.VerifyOTP(ctx, otpVerificationDetail)
	if err != nil {
		log.Err(err).Msg("Failed to verify OTP")
		return nil, err
	}

	return &authDetails, nil
}
