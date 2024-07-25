package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type ReferralService interface {
	CreateReferral(ctx context.Context, referral domain.ReferralParams) (*domain.Referral, error)
	FetchReferrals(ctx context.Context, referral domain.ReferralParams) (*domain.Referrals, error)
	ExpireReferrals(ctx context.Context, referral *domain.Referral) (*domain.EmptyDomain, error)
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
	requester, err := r.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{
		UserUUID: referralDetails.RequesterUserUUID,
	})
	if err != nil {
		return nil, errors.ErrRequesterNotFound
	}

	provider, err := r.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{
		UserUUID: referralDetails.ProviderUserUUID,
	})
	if err != nil {
		return nil, errors.ErrProviderNotFound
	}

	if provider.CompanyID != referralDetails.CompanyID {
		return nil, errors.ErrCompanyNotMatch
	}

	referralMaxLookupTime := time.Now().Add(-r.cfg.Referral.ReferralMaxLookupDuration)
	requesterReferrals, err := r.registry.ReferralsRepository.FetchReferralsForParams(ctx, domain.ReferralParams{
		Referral: domain.Referral{
			RequesterUserUUID: requester.UserUUID,
			CreatedAt:         &referralMaxLookupTime,
			Status:            constants.ReferralStatusPending,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(*requesterReferrals) >= r.cfg.Referral.RequesterReferralLimit {
		return nil, errors.ErrReferralLimitReachedForRequester
	}

	providerReferrals, err := r.registry.ReferralsRepository.FetchReferralsForParams(ctx, domain.ReferralParams{
		Referral: domain.Referral{
			ProviderUserUUID: provider.UserUUID,
			CreatedAt:        &referralMaxLookupTime,
			Status:           constants.ReferralStatusPending,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(*providerReferrals) >= r.cfg.Referral.ProviderReferralLimit {
		return nil, errors.ErrReferralLimitReachedForProvider
	}

	if referralExists(*requesterReferrals, provider.UserUUID) {
		return nil, errors.ErrReferralAlreadyExists
	}

	storeResumeFilePath, err := storeResumeFile(referralDetails.ResumeFile, r.cfg.ResumeDirectory, requester.UserNumber)
	if err != nil {
		return nil, err
	}

	referral, err := r.registry.ReferralsRepository.InsertReferral(ctx, referralDetails.ToReferral())
	if err != nil {
		return nil, err
	}

	referralMailParams := builder.ReferralMailParams{
		Requester:      *requester,
		Provider:       *provider,
		JobLink:        referralDetails.JobLink,
		Message:        referralDetails.Message,
		ResumeFilePath: storeResumeFilePath,
		AttachmentName: getResumeFileName(requester.Name),
	}

	err = r.registry.MailBuilder.SendReferralMail(ctx, true, referral.Status, referralMailParams)
	return referral, err
}

func (r referralService) FetchReferrals(ctx context.Context, referral domain.ReferralParams) (*domain.Referrals, error) {
	return r.registry.ReferralsRepository.FetchReferralsForParams(ctx, referral)
}

func (r referralService) ExpireReferrals(ctx context.Context, referral *domain.Referral) (*domain.EmptyDomain, error) {
	err := r.registry.ReferralsRepository.ExpirePendingReferrals(ctx, referral)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func referralExists(requesterReferrals domain.Referrals, providerUserUUID string) bool {
	for _, r := range requesterReferrals {
		if r.ProviderUserUUID == providerUserUUID {
			return true
		}
	}

	return false
}

// ToDo: Upload resume to S3
func storeResumeFile(file multipart.File, resumeDirectory string, userNumber int64) (string, error) {
	if err := utils.CreateDirectoryIfNotExist(resumeDirectory); err != nil {
		return "", err
	}

	resumeFileName := fmt.Sprintf("resume_user_number_%d_%d.pdf", userNumber, time.Now().Unix())
	return utils.StoreMultipartFile(file, resumeDirectory, resumeFileName)
}

func getResumeFileName(name string) string {
	temp := strings.Split(name, " ")
	sanitizedName := strings.Join(temp, "_")
	return fmt.Sprintf("Resume_%s.pdf", sanitizedName)
}
