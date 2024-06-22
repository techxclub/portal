package service

import (
	"cmp"
	"context"
	"slices"
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

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error)
	UpdateUserDetails(ctx context.Context, userDetails domain.UserProfile) (*domain.EmptyDomain, error)
	GetProfile(ctx context.Context, params domain.FetchUserParams) (*domain.UserProfile, error)
	GetUsers(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error)
	GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error)
}

type userService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewUserService(cfg *config.Config, registry *builder.Registry) UserService {
	return &userService{
		cfg:      cfg,
		registry: registry,
	}
}

func (u userService) RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error) {
	var user *domain.UserProfile
	var err error

	normalizedCompanyName := strcase.ToScreamingSnake(strings.ToUpper(userDetails.CompanyName))
	companyDetails, fetchCompanyErr := u.registry.CompaniesRepository.FetchCompanyForParams(ctx, domain.FetchCompanyParams{NormalizedName: normalizedCompanyName})
	if fetchCompanyErr != nil {
		companyDetails, err = u.registry.CompaniesRepository.InsertCompany(ctx, domain.Company{
			NormalizedName: normalizedCompanyName,
			DisplayName:    userDetails.CompanyName,
			Verified:       utils.ToPtr(false),
			Popular:        utils.ToPtr(false),
			Actor:          constants.ActorUser,
		})
		if err != nil {
			log.Info().Err(err).Msg("Failed to add company")
			return nil, err
		}
	}

	if companyDetails == nil || companyDetails.ID <= 0 {
		return nil, errors.ErrCompanyNotFound
	}

	userDetails.CompanyID = companyDetails.ID
	userDetails.CompanyName = companyDetails.DisplayName
	user, err = u.registry.UsersRepository.Insert(ctx, userDetails)
	if err != nil {
		return nil, err
	}

	authToken, err := domain.GenerateToken(user.UserID, u.cfg.Auth)
	if err != nil {
		return nil, err
	}

	return &domain.Registration{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (u userService) UpdateUserDetails(ctx context.Context, params domain.UserProfile) (*domain.EmptyDomain, error) {
	err := u.registry.UsersRepository.Update(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (u userService) GetProfile(ctx context.Context, params domain.FetchUserParams) (*domain.UserProfile, error) {
	users, err := u.registry.UsersRepository.FetchUserForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userService) GetUsers(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error) {
	users, err := u.registry.UsersRepository.FetchUsersForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(*users, func(i, j domain.UserProfile) int {
		return cmp.Compare(i.Name, j.Name)
	})

	return users, nil
}

func (u userService) GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	companies, err := u.registry.CompaniesRepository.FetchCompaniesForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(*companies, func(i, j domain.Company) int {
		return cmp.Compare(i.GetPriority(), j.GetPriority())
	})
	return companies, nil
}
