package service

import (
	"cmp"
	"context"
	"slices"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error)
	RegisterMentor(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error)
	UpdateUserDetails(ctx context.Context, userDetails domain.UserProfile) (*domain.EmptyDomain, error)
	GetProfile(ctx context.Context, params domain.FetchUserParams) (*domain.UserProfile, error)
	GetUsers(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error)
	GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error)
	GetCompanyUsers(ctx context.Context, params domain.FetchUserParams) (*domain.CompanyUsersService, error)
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

func (u userService) RegisterMentor(ctx context.Context, userDetails domain.UserProfile) (*domain.Registration, error) {
	user, err := u.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserID: userDetails.UserID})
	if err != nil {
		return nil, err
	}

	if !user.IsApproved() {
		return nil, errors.ErrUserNotApproved
	}

	currentMentorConfig := *user.MentorConfig
	if currentMentorConfig.IsMentor {
		return nil, errors.ErrUserAlreadyMentor
	}

	updatedMentorConfig := currentMentorConfig
	updatedMentorConfig.Tags = userDetails.MentorConfig.Tags
	updatedMentorConfig.CalendalyLink = userDetails.MentorConfig.CalendalyLink
	updatedMentorConfig.Domain = userDetails.MentorConfig.Domain

	if currentMentorConfig.IsPreApproved {
		updatedMentorConfig.IsMentor = true
		updatedMentorConfig.Status = constants.MentorStatusApproved
	} else {
		updatedMentorConfig.IsMentor = false
		updatedMentorConfig.Status = constants.MentorStatusPendingApproval
	}

	user.MentorConfig = &updatedMentorConfig

	err = u.registry.UsersRepository.Update(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &domain.Registration{
		User: user,
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
	userID := apicontext.RequestContextFromContext(ctx).GetUserID()
	user, _ := u.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserID: userID})
	if userID != "" && user == nil {
		return nil, errors.ErrUserNotFound
	}

	users, err := u.registry.UsersRepository.FetchUsersForParams(ctx, params)
	if err != nil {
		return nil, err
	}
	slices.SortStableFunc(*users, func(i, j domain.UserProfile) int {
		return cmp.Compare(i.Name, j.Name)
	})

	if user == nil {
		return users, nil
	}

	filteredUsers := utils.Filter(*users, func(user domain.UserProfile) bool {
		return user.UserID != userID
	})

	if len(filteredUsers) == 0 {
		return nil, errors.ErrNoDataFound
	}

	return &filteredUsers, nil
}

func (u userService) GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	userID := apicontext.RequestContextFromContext(ctx).GetUserID()
	user, _ := u.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserID: userID})
	if userID != "" && user == nil {
		return nil, errors.ErrUserNotFound
	}

	companies, err := u.registry.CompaniesRepository.FetchCompaniesForParams(ctx, params)
	if err != nil {
		return nil, err
	}
	slices.SortStableFunc(*companies, func(i, j domain.Company) int {
		return cmp.Compare(i.GetPriority(), j.GetPriority())
	})

	if user == nil {
		return companies, nil
	}

	filteredCompanies := utils.Filter(*companies, func(company domain.Company) bool {
		return company.ID != user.CompanyID
	})

	if len(filteredCompanies) == 0 {
		return nil, errors.ErrNoDataFound
	}

	return &filteredCompanies, nil
}

func (u userService) GetCompanyUsers(ctx context.Context, params domain.FetchUserParams) (*domain.CompanyUsersService, error) {
	companyUsers, err := u.GetUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	userID := apicontext.RequestContextFromContext(ctx).GetUserID()
	referralParams := domain.ReferralParams{
		RequesterUserID: userID,
		CreatedAt:       utils.ToPtr(time.Now().Add(-7 * 24 * time.Hour)),
	}
	userReferrals, err := u.registry.ReferralsRepository.FetchReferralsForParams(ctx, referralParams)
	if err != nil {
		return nil, err
	}

	return &domain.CompanyUsersService{
		Users:     companyUsers,
		Referrals: userReferrals,
	}, nil
}
