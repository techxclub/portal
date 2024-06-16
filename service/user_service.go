package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
	"golang.org/x/sync/errgroup"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error)
	UpdateUserDetails(ctx context.Context, userDetails domain.UserProfile) (*domain.EmptyDomain, error)
	GetProfile(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error)
	GetUsers(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error)
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
	eg := errgroup.Group{}

	var user *domain.UserProfile

	eg.Go(func() (err error) {
		user, err = u.registry.UsersRepository.CreateUser(ctx, userDetails)
		return
	})

	eg.Go(func() error {
		_, err := u.registry.CompaniesRepository.AddCompany(ctx, domain.Company{
			Name:     userDetails.Company,
			Verified: utils.ToPtr(false),
			Popular:  utils.ToPtr(false),
			Actor:    constants.ActorUser,
		})
		if err != nil {
			log.Info().Err(err).Msg("Failed to add company")
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
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
	err := u.registry.UsersRepository.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (u userService) GetProfile(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error) {
	users, err := u.registry.UsersRepository.GetUserForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userService) GetUsers(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error) {
	users, err := u.registry.UsersRepository.GetUsersForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userService) GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	companies, err := u.registry.CompaniesRepository.GetCompaniesForParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
