package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type AuthService interface {
	GenerateOTP(ctx context.Context, detail domain.OTPGeneration) (*domain.AuthDetails, error)
	VerifyOTP(ctx context.Context, detail domain.OTPVerification) (*domain.AuthDetails, error)
}

type authService struct {
	cfg      config.Config
	registry *builder.Registry
}

func NewAuthService(cfg config.Config, registry *builder.Registry) AuthService {
	return &authService{
		cfg:      cfg,
		registry: registry,
	}
}

func (s authService) GenerateOTP(ctx context.Context, detail domain.OTPGeneration) (*domain.AuthDetails, error) {
	authDetails, err := s.registry.UserAuthBuilder.GenerateOTP(ctx, detail)
	if err != nil {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	return authDetails, nil
}

func (s authService) VerifyOTP(ctx context.Context, detail domain.OTPVerification) (*domain.AuthDetails, error) {
	authDetails, err := s.registry.UserAuthBuilder.VerifyOTP(ctx, detail)
	if err != nil {
		log.Err(err).Msg("Failed to verify OTP")
		return nil, err
	}

	return authDetails, nil
}
