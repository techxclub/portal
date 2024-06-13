package service

import (
	"context"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type ReferralService interface {
	CreateReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
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

func (r referralService) CreateReferral(ctx context.Context, referralDetails domain.ReferralParams) (*domain.Referral, error) {
	// Validate referral requester exists
	requester, err := r.registry.UsersRepo.GetUserForParams(ctx, domain.UserProfileParams{
		UserID: referralDetails.RequesterUserID,
	})
	if err != nil {
		return nil, errors.ErrRequesterNotFound
	}

	// Validate referral provider exists
	provider, err := r.registry.UsersRepo.GetUserForParams(ctx, domain.UserProfileParams{
		UserID: referralDetails.ProviderUserID,
	})
	if err != nil {
		return nil, errors.ErrProviderNotFound
	}

	// Validate company
	if provider.Company != referralDetails.Company {
		return nil, errors.ErrCompanyNotMatch
	}

	// Check if referral already exists
	referralExist, err := r.registry.ReferralsRepo.GetReferralsForParams(ctx, domain.ReferralParams{
		RequesterUserID: requester.UserID,
		ProviderUserID:  provider.UserID,
		CreatedAt:       &r.cfg.Referral.ReferralMaxTime,
	})
	if err != nil {
		return nil, err
	}

	if len(*referralExist) > 0 {
		return nil, errors.ErrReferralAlreadyExists
	}

	// All referrals for requester
	requesterReferrals, err := r.registry.ReferralsRepo.GetReferralsForParams(ctx, domain.ReferralParams{
		RequesterUserID: requester.UserID,
	})
	if err != nil {
		return nil, err
	}

	// Validate requesterReferral limit
	if len(*requesterReferrals) >= r.cfg.Referral.RequesterReferralLimit {
		return nil, errors.ErrReferralLimitReachedForRequester
	}

	// All referrals for provider
	providerReferrals, err := r.registry.ReferralsRepo.GetReferralsForParams(ctx, domain.ReferralParams{
		ProviderUserID: provider.UserID,
	})
	if err != nil {
		return nil, err
	}

	// Validate providerReferral limit
	if len(*providerReferrals) >= r.cfg.Referral.ProviderReferralLimit {
		return nil, errors.ErrReferralLimitReachedForProvider
	}

	// Send email to provider
	referralMailParams := builder.ReferralMailParams{
		Requester: *requester,
		Provider:  *provider,
		JobLink:   referralDetails.JobLink,
		Message:   referralDetails.Message,
	}
	err = r.registry.MailBuilder.SendReferralMail(ctx, referralMailParams)
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
