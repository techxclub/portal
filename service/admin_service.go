package service

import (
	"context"
	"database/sql"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type AdminService interface {
	ApproveUser(ctx context.Context, user domain.User) (*domain.EmptyDomain, error)
	UpdateUsers(ctx context.Context, from, to domain.User) (*domain.EmptyDomain, error)
	UpdateCompanyDetails(ctx context.Context, company domain.Company) (*domain.EmptyDomain, error)
	UpdateReferralDetails(ctx context.Context, params *domain.Referral) (*domain.EmptyDomain, error)
}

type adminService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewAdminService(cfg *config.Config, registry *builder.Registry) AdminService {
	return &adminService{
		cfg:      cfg,
		registry: registry,
	}
}

func (as adminService) ApproveUser(ctx context.Context, params domain.User) (*domain.EmptyDomain, error) {
	user, err := as.getUserDetails(ctx, params)
	if err != nil {
		return nil, err
	}

	if user.Status != constants.StatusPendingApproval {
		return nil, errors.ErrInvalidUpdateRequest
	}

	user.Status = constants.StatusApproved
	user, err = as.handleCompanyUpdate(ctx, user)
	if err != nil {
		return nil, err
	}

	err = as.registry.UsersRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	approvalMailParams := builder.ApprovalMailParams{User: *user}
	refID := utils.GetRandomUUID()
	err = as.registry.MailBuilder.SendUserApprovalMail(ctx, true, refID, approvalMailParams)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (as adminService) UpdateUsers(ctx context.Context, from, to domain.User) (*domain.EmptyDomain, error) {
	if to.CompanyName != "" {
		normalizedCompanyName := strcase.ToScreamingSnake(strings.ToUpper(to.CompanyName))
		companyDetails, err := as.registry.CompaniesRepository.FetchCompanyForParams(ctx, domain.FetchCompanyParams{NormalizedName: normalizedCompanyName})
		if err != nil {
			return nil, err
		}

		to.CompanyID = companyDetails.ID
	}

	err := as.registry.UsersRepository.BulkUpdate(ctx, from, to)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (as adminService) UpdateCompanyDetails(ctx context.Context, company domain.Company) (*domain.EmptyDomain, error) {
	if company.DisplayName != "" {
		company.NormalizedName = strcase.ToScreamingSnake(strings.ToUpper(company.DisplayName))
	}
	err := as.registry.CompaniesRepository.UpdateCompany(ctx, &company)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (as adminService) UpdateReferralDetails(ctx context.Context, params *domain.Referral) (*domain.EmptyDomain, error) {
	err := as.registry.ReferralsRepository.UpdateReferral(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (as adminService) getUserDetails(ctx context.Context, params domain.User) (*domain.User, error) {
	var user *domain.User
	var err error

	switch {
	case params.UserNumber != 0:
		user, err = as.registry.UsersRepository.FindByUserNumber(ctx, params.UserNumber)
	case params.RegisteredEmail != "":
		user, err = as.registry.UsersRepository.FindByRegisteredEmail(ctx, params.RegisteredEmail)
	default:
		user, err = as.registry.UsersRepository.FindByUserUUID(ctx, params.UserUUID)
	}

	return user, err
}

func (as adminService) handleCompanyUpdate(ctx context.Context, user *domain.User) (*domain.User, error) {
	normalizedCompanyName := strcase.ToScreamingSnake(strings.ToUpper(user.CompanyName))
	companyDetails, err := as.registry.CompaniesRepository.FetchCompanyForParams(ctx, domain.FetchCompanyParams{NormalizedName: normalizedCompanyName})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		var insertErr error
		companyDetails, insertErr = as.registry.CompaniesRepository.InsertCompany(ctx, domain.Company{
			NormalizedName: normalizedCompanyName,
			DisplayName:    user.CompanyName,
			Verified:       utils.ToPtr(false),
			Popular:        utils.ToPtr(false),
			Actor:          constants.ActorAdmin,
		})

		if insertErr != nil {
			log.Info().Err(err).Msg("Failed to add new company")
			return user, err
		}
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	user.CompanyID = companyDetails.ID
	user.CompanyName = companyDetails.DisplayName
	return user, nil
}
