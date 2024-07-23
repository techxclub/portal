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
	RegisterUser(ctx context.Context, userDetails domain.User) (*domain.Registration, error)
	RegisterMentor(ctx context.Context, userDetails domain.User) (*domain.Registration, error)
	UpdateUser(ctx context.Context, userDetails domain.User) (*domain.User, error)
	GetUser(ctx context.Context, params domain.FetchUserParams) (*domain.User, error)
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

func (us userService) RegisterUser(ctx context.Context, userDetails domain.User) (*domain.Registration, error) {
	var user *domain.User
	var err error

	normalizedCompanyName := strcase.ToScreamingSnake(strings.ToUpper(userDetails.CompanyName))
	companyDetails, fetchCompanyErr := us.registry.CompaniesRepository.FetchCompanyForParams(ctx, domain.FetchCompanyParams{NormalizedName: normalizedCompanyName})
	if fetchCompanyErr != nil {
		companyDetails, err = us.registry.CompaniesRepository.InsertCompany(ctx, domain.Company{
			NormalizedName: normalizedCompanyName,
			DisplayName:    userDetails.CompanyName,
			Verified:       utils.ToPtr(false),
			Popular:        utils.ToPtr(false),
			Actor:          constants.ActorUser,
		})
		if err != nil {
			log.Info().Err(err).Msg("Failed to add new company")
			return nil, err
		}
	}

	userDetails.CompanyID = companyDetails.ID
	userDetails.CompanyName = companyDetails.DisplayName
	user, err = us.registry.UsersRepository.Insert(ctx, userDetails)
	if err != nil {
		return nil, err
	}

	authToken, err := domain.GenerateToken(user.UserUUID, us.cfg.Auth)
	if err != nil {
		return nil, err
	}

	return &domain.Registration{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (us userService) RegisterMentor(ctx context.Context, userDetails domain.User) (*domain.Registration, error) {
	user, err := us.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserUUID: userDetails.UserUUID})
	if err != nil {
		return nil, err
	}

	if !user.IsApproved() {
		return nil, errors.ErrUserNotApproved
	}

	currentMentorConfig := user.MentorConfig()
	if currentMentorConfig.IsMentor {
		return nil, errors.ErrUserAlreadyMentor
	}

	updatedMentorConfig := currentMentorConfig
	updatedMentorConfig.Tags = userDetails.MentorConfig().Tags
	updatedMentorConfig.CalendalyLink = userDetails.MentorConfig().CalendalyLink
	updatedMentorConfig.Domain = userDetails.MentorConfig().Domain

	if currentMentorConfig.IsPreApproved {
		updatedMentorConfig.IsMentor = true
		updatedMentorConfig.Status = constants.MentorStatusApproved
	} else {
		updatedMentorConfig.IsMentor = false
		updatedMentorConfig.Status = constants.MentorStatusPendingApproval
	}

	user.SetMentorConfig(updatedMentorConfig)
	err = us.registry.UsersRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &domain.Registration{
		User: user,
	}, nil
}

func (us userService) UpdateUser(ctx context.Context, updatedUser domain.User) (*domain.User, error) {
	storedUser, err := us.registry.UsersRepository.FindByUserUUID(ctx, updatedUser.UserUUID)
	if err != nil {
		return nil, err
	}

	if err := us.validateUserUpdateDetails(ctx, *storedUser, updatedUser); err != nil {
		return nil, err
	}

	updatedUser.CreatedAt = storedUser.CreatedAt
	updatedUser.Status = storedUser.Status
	if updatedUser.CompanyName != storedUser.CompanyName {
		updatedUser, err = us.handleCompanyUpdate(ctx, updatedUser)
		if err != nil {
			return nil, err
		}
	} else {
		updatedUser.CompanyID = storedUser.CompanyID
	}

	if updatedUser.Status == constants.StatusIncompleteProfile && updatedUser.IsProfileComplete() {
		updatedUser.Status = constants.StatusPendingApproval
	}

	err = us.registry.UsersRepository.Update(ctx, &updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (us userService) GetUser(ctx context.Context, params domain.FetchUserParams) (*domain.User, error) {
	user, err := us.registry.UsersRepository.FindByUserUUID(ctx, params.UserUUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us userService) GetUsers(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error) {
	userID := apicontext.RequestContextFromContext(ctx).GetUserUUID()
	user, err := us.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserUUID: userID})
	if userID != "" && err != nil {
		return nil, errors.ErrUserNotFound
	}

	users, err := us.registry.UsersRepository.FetchUsersForParams(ctx, params)
	if err != nil {
		return nil, err
	}
	slices.SortStableFunc(*users, func(i, j domain.User) int {
		return cmp.Compare(i.Name, j.Name)
	})

	if user == nil {
		return users, nil
	}

	filteredUsers := utils.Filter(*users, func(user domain.User) bool {
		return user.UserUUID != userID
	})

	if len(filteredUsers) == 0 {
		return nil, errors.ErrNoDataFound
	}

	return &filteredUsers, nil
}

func (us userService) GetCompanies(ctx context.Context, params domain.FetchCompanyParams) (*domain.Companies, error) {
	userID := apicontext.RequestContextFromContext(ctx).GetUserUUID()
	user, _ := us.registry.UsersRepository.FetchUserForParams(ctx, domain.FetchUserParams{UserUUID: userID})
	// ToDo: Update this flow to be used by only approved users
	if userID != "" && user == nil {
		return nil, errors.ErrUserNotFound
	}

	companies, err := us.registry.CompaniesRepository.FetchCompaniesForParams(ctx, params)
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

func (us userService) GetCompanyUsers(ctx context.Context, params domain.FetchUserParams) (*domain.CompanyUsersService, error) {
	companyUsers, err := us.GetUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	userID := apicontext.RequestContextFromContext(ctx).GetUserUUID()
	referralParams := domain.ReferralParams{
		Referral: domain.Referral{
			RequesterUserUUID: userID,
			CreatedAt:         utils.ToPtr(time.Now().Add(-7 * 24 * time.Hour)),
		},
	}
	userReferrals, err := us.registry.ReferralsRepository.FetchReferralsForParams(ctx, referralParams)
	if err != nil {
		return nil, err
	}

	return &domain.CompanyUsersService{
		Users:     companyUsers,
		Referrals: userReferrals,
	}, nil
}

func (us userService) handleCompanyUpdate(ctx context.Context, user domain.User) (domain.User, error) {
	normalizedCompanyName := strcase.ToScreamingSnake(strings.ToUpper(user.CompanyName))
	companyDetails, fetchCompanyErr := us.registry.CompaniesRepository.FetchCompanyForParams(ctx, domain.FetchCompanyParams{NormalizedName: normalizedCompanyName})
	if fetchCompanyErr != nil {
		var err error
		companyDetails, err = us.registry.CompaniesRepository.InsertCompany(ctx, domain.Company{
			NormalizedName: normalizedCompanyName,
			DisplayName:    user.CompanyName,
			Verified:       utils.ToPtr(false),
			Popular:        utils.ToPtr(false),
			Actor:          constants.ActorUser,
		})
		if err != nil {
			log.Info().Err(err).Msg("Failed to add new company")
			return user, err
		}
	}

	user.CompanyID = companyDetails.ID
	user.CompanyName = companyDetails.DisplayName
	return user, nil
}

func (us userService) validateUserUpdateDetails(ctx context.Context, storedUser, updatedUser domain.User) error {
	if storedUser.RegisteredEmail != updatedUser.RegisteredEmail {
		return errors.ErrInvalidUserUpdate
	}

	if storedUser.WorkEmail != updatedUser.WorkEmail {
		isUpdatedWorkEmailVerified := us.registry.OTPBuilder.Check(ctx, updatedUser.WorkEmail)
		if !isUpdatedWorkEmailVerified {
			return errors.ErrWorkEmailNotVerified
		}
	}

	if utils.UpdatedToZeroValue(storedUser.Name, updatedUser.Name) {
		return errors.ErrEmptyName
	}

	if utils.UpdatedToZeroValue(storedUser.PhoneNumber, updatedUser.PhoneNumber) {
		return errors.ErrEmptyPhoneNumber
	}

	if utils.UpdatedToZeroValue(storedUser.LinkedIn, updatedUser.LinkedIn) {
		return errors.ErrEmptyLinkedIn
	}

	if utils.UpdatedToZeroValue(storedUser.CompanyName, updatedUser.CompanyName) {
		return errors.ErrEmptyCompanyName
	}

	if utils.UpdatedToZeroValue(storedUser.WorkEmail, updatedUser.WorkEmail) {
		return errors.ErrEmptyWorkEmail
	}

	if utils.UpdatedToZeroValue(storedUser.Designation, updatedUser.Designation) {
		return errors.ErrEmptyDesignation
	}

	if utils.UpdatedToZeroValue(storedUser.YearsOfExperience, updatedUser.YearsOfExperience) {
		return errors.ErrEmptyYearsOfExperience
	}

	return nil
}
