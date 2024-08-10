package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type OAuthService interface {
	GoogleSignIn(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.User, error)
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

func (as oauthService) GoogleSignIn(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.User, error) {
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

	user, err := as.registry.UsersRepository.Insert(ctx, *userProfile)
	if err != nil {
		return nil, err
	}

	if exchangeReq.InviteCode != "" {
		invite := domain.Invite{
			Code:          exchangeReq.InviteCode,
			InvitedUserID: user.UserUUID,
		}

		_, err = as.registry.InvitesRepository.Insert(ctx, invite)
		if err != nil {
			log.Err(err).Msg("Failed to insert invite")
		}
	}

	return user, nil
}
