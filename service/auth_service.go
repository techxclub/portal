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
	VerifyUser(ctx context.Context, detail domain.AuthRequest) (*domain.AuthDetails, error)
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

func (s authService) GenerateOTP(ctx context.Context, otpGenerationDetail domain.AuthRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.AuthBuilder.GenerateOTP(ctx, otpGenerationDetail)
	if err != nil || authInfo.Status != constants.AuthStatusPending {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}
	return &authDetails, nil
}

func (s authService) VerifyUser(ctx context.Context, otpVerificationDetail domain.AuthRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.AuthBuilder.VerifyOTP(ctx, otpVerificationDetail)
	if err != nil {
		log.Err(err).Msg("Failed to verify OTP")
		return nil, err
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}

	if otpVerificationDetail.Channel != constants.AuthChannelSMS || authInfo.Status != constants.AuthStatusApproved {
		return &authDetails, nil
	}

	phoneNumber := otpVerificationDetail.Value
	userProfileParams := domain.UserProfileParams{PhoneNumber: phoneNumber}
	userInfo, err := s.registry.UserInfoBuilder.GetUserForParams(ctx, userProfileParams)
	if err != nil {
		log.Err(err).Msg("Failed to get user profile")
		return &authDetails, nil
	}

	authDetails.UserInfo = userInfo
	return &authDetails, nil
}
