package service

import (
	"context"
	"time"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type ReferralService interface {
	CreateReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
}

type referralService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewReferralService(cfg *config.Config, registry *builder.Registry) ReferralService {
	return &referralService{
		cfg:      cfg,
		registry: registry,
	}
}

func (r referralService) CreateReferral(ctx context.Context, referralDetails domain.ReferralParams) (*domain.Referral, error) {
	requester, err := r.registry.UsersRepository.GetUserForParams(ctx, domain.UserProfileParams{
		UserID: referralDetails.RequesterUserID,
	})
	if err != nil {
		return nil, errors.ErrRequesterNotFound
	}

	provider, err := r.registry.UsersRepository.GetUserForParams(ctx, domain.UserProfileParams{
		UserID: referralDetails.ProviderUserID,
	})
	if err != nil {
		return nil, errors.ErrProviderNotFound
	}

	if provider.Company != referralDetails.Company {
		return nil, errors.ErrCompanyNotMatch
	}

	referralMaxLookupTime := time.Now().Add(-r.cfg.Referral.ReferralMaxLookupDuration)
	requesterReferrals, err := r.registry.ReferralsRepository.GetReferralsForParams(ctx, domain.ReferralParams{
		RequesterUserID: requester.UserID,
		CreatedAt:       &referralMaxLookupTime,
		Status:          constants.ReferralStatusPending,
	})
	if err != nil {
		return nil, err
	}

	if len(*requesterReferrals) >= r.cfg.Referral.RequesterReferralLimit {
		return nil, errors.ErrReferralLimitReachedForRequester
	}

	providerReferrals, err := r.registry.ReferralsRepository.GetReferralsForParams(ctx, domain.ReferralParams{
		ProviderUserID: provider.UserID,
		CreatedAt:      &referralMaxLookupTime,
		Status:         constants.ReferralStatusPending,
	})
	if err != nil {
		return nil, err
	}

	if len(*providerReferrals) >= r.cfg.Referral.ProviderReferralLimit {
		return nil, errors.ErrReferralLimitReachedForProvider
	}

	if referralExists(*requesterReferrals, provider.UserID) {
		return nil, errors.ErrReferralAlreadyExists
	}

	referralMailParams := builder.ReferralMailParams{
		Requester:      *requester,
		Provider:       *provider,
		JobLink:        referralDetails.JobLink,
		Message:        referralDetails.Message,
		ResumeFilePath: referralDetails.ResumeFilePath,
	}
	err = r.registry.MailBuilder.SendReferralMail(ctx, referralMailParams)
	if err != nil {
		return nil, err
	}

	referral, err := r.registry.ReferralsRepository.CreateReferral(ctx, referralDetails)
	if err != nil {
		return nil, err
	}

	return referral, nil
}

func referralExists(requesterReferrals domain.Referrals, providerUserID string) bool {
	for _, r := range requesterReferrals {
		if r.ProviderUserID == providerUserID {
			return true
		}
	}

	return false
}
