package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type AuthService interface {
	GenerateOTP(ctx context.Context, detail domain.AuthRequest) (*domain.AuthDetails, error)
	ResendOTP(ctx context.Context, detail domain.AuthRequest) (*domain.AuthDetails, error)
	VerifyOTP(ctx context.Context, detail domain.AuthRequest) (*domain.AuthDetails, error)
}

type authService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewAuthService(cfg *config.Config, registry *builder.Registry) AuthService {
	return &authService{
		cfg:      cfg,
		registry: registry,
	}
}

func (s authService) GenerateOTP(ctx context.Context, otpGenerationDetail domain.AuthRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.SendOTP(ctx, otpGenerationDetail)
	if err != nil || authInfo.Status != constants.AuthStatusPending {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}
	return &authDetails, nil
}

func (s authService) ResendOTP(ctx context.Context, otpGenerationDetail domain.AuthRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.SendOTP(ctx, otpGenerationDetail)
	if err != nil || authInfo.Status != constants.AuthStatusPending {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}
	return &authDetails, nil
}

func (s authService) VerifyOTP(ctx context.Context, otpVerificationDetail domain.AuthRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.VerifyOTP(ctx, otpVerificationDetail)
	if err != nil {
		log.Err(err).Msg("Failed to verify OTP")
		return nil, err
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}

	if otpVerificationDetail.Channel != constants.AuthChannelSMS || authInfo.Status != constants.AuthStatusVerified {
		return &authDetails, nil
	}

	phoneNumber := otpVerificationDetail.Value
	userProfileParams := domain.FetchUserParams{PhoneNumber: phoneNumber}
	userInfo, err := s.registry.UsersRepository.FetchUserForParams(ctx, userProfileParams)
	if err != nil {
		log.Err(err).Msg("User is not registered")
		return &authDetails, nil
	}

	authToken, _ := domain.GenerateToken(userInfo.UserID, s.cfg.Auth)

	authDetails.UserInfo = userInfo
	authDetails.AuthToken = authToken
	return &authDetails, nil
}
