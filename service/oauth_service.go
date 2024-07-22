package service

import (
	"context"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type OAuthService interface {
	GoogleOAuthExchange(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.User, error)
}

type oauthService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewOAuthService(cfg *config.Config, registry *builder.Registry) OAuthService {
	return &oauthService{
		cfg:      cfg,
		registry: registry,
	}
}

func (as oauthService) GoogleOAuthExchange(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.User, error) {
	googleAuthDetails, err := as.registry.GoogleOAuthBuilder.BuildGoogleOAuthDetails(ctx, exchangeReq)
	if err != nil {
		return nil, err
	}

	userProfile, err := as.registry.GoogleOAuthBuilder.BuildUserProfile(ctx, *googleAuthDetails)
	if err != nil {
		return nil, err
	}

	// Check if user exists in the database
	storedProfile, err := as.registry.UsersRepository.FindByRegisteredEmail(ctx, userProfile.RegisteredEmail)
	if err == nil {
		return storedProfile, nil
	}

	return as.registry.UsersRepository.Insert(ctx, *userProfile)
}
