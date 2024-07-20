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

func (s otpService) GenerateOTP(ctx context.Context, otpGenerationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.SendOTP(ctx, otpGenerationDetail)
	if err != nil || authInfo.Status != constants.OTPStatusPending {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}
	return &authDetails, nil
}

func (s otpService) ResendOTP(ctx context.Context, otpGenerationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.ResendOTP(ctx, otpGenerationDetail)
	if err != nil || authInfo.Status != constants.OTPStatusPending {
		log.Err(err).Msg("Failed to generate OTP")
		return nil, errors.ErrOTPGenerateFailed
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}
	return &authDetails, nil
}

func (s otpService) VerifyOTP(ctx context.Context, otpVerificationDetail domain.OTPRequest) (*domain.AuthDetails, error) {
	authInfo, err := s.registry.OTPBuilder.VerifyOTP(ctx, otpVerificationDetail)
	if err != nil {
		log.Err(err).Msg("Failed to verify OTP")
		return nil, err
	}

	authDetails := domain.AuthDetails{
		AuthInfo: authInfo,
	}

	if authInfo.Status != constants.OTPStatusVerified {
		return &authDetails, nil
	}

	if otpVerificationDetail.Channel == constants.OTPChannelEmail {
		authToken, _ := domain.GenerateToken(otpVerificationDetail.Value, s.cfg.Auth)
		authDetails.Token = authToken
		return &authDetails, nil
	}

	phoneNumber := otpVerificationDetail.Value
	userProfileParams := domain.FetchUserParams{PhoneNumber: phoneNumber}
	userInfo, err := s.registry.UsersRepository.FetchUserForParams(ctx, userProfileParams)
	if err != nil {
		log.Err(err).Msg("User is not registered")
		return &authDetails, nil
	}

	authToken, _ := domain.GenerateToken(userInfo.UserUUID, s.cfg.Auth)

	authDetails.UserInfo = userInfo
	authDetails.Token = authToken
	return &authDetails, nil
}
