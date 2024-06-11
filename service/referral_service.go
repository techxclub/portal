package service

import (
	"context"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type ReferralService interface {
	CreateReferral(ctx context.Context, referral domain.Referral) (*domain.Referral, error)
}

type referralService struct {
	cfg      config.Config
	registry *builder.Registry
}

func NewReferralService(cfg config.Config, registry *builder.Registry) ReferralService {
	return &referralService{
		cfg:      cfg,
		registry: registry,
	}
}

func (r referralService) CreateReferral(ctx context.Context, referralDetails domain.Referral) (*domain.Referral, error) {
	// Validate referral requester exists
	_, err := r.registry.UsersRepo.GetUserForParams(ctx, domain.UserProfileParams{UserID: referralDetails.RequesterUserID})
	if err != nil {
		return nil, err
	}

	// Validate referral provider exists
	_, err = r.registry.UsersRepo.GetUserForParams(ctx, domain.UserProfileParams{UserID: referralDetails.ProviderUserID})
	if err != nil {
		return nil, err
	}

	// Create referral
	referral, err := r.registry.ReferralsRepo.CreateReferral(ctx, referralDetails)
	if err != nil {
		return nil, err
	}
	return referral, nil
}
