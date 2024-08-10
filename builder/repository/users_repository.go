package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type UsersRepository interface {
	Insert(ctx context.Context, user domain.User) (*domain.User, error)
	Update(ctx context.Context, updatedUser *domain.User) error
	BulkUpdate(ctx context.Context, from, to domain.User) error
	FindByUserNumber(ctx context.Context, userNumber int64) (*domain.User, error)
	FindByUserUUID(ctx context.Context, uuid string) (*domain.User, error)
	FindByRegisteredEmail(ctx context.Context, email string) (*domain.User, error)
	FindByWorkEmail(ctx context.Context, email string) (*domain.User, error)
	FetchUserForParams(ctx context.Context, params domain.FetchUserParams) (*domain.User, error)
	FetchUsersForParams(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error)
}

type usersRepository struct {
	dbClient db.Client
}

func NewUsersRepository(userDB db.Client) UsersRepository {
	return &usersRepository{
		dbClient: userDB,
	}
}

type UsersReturning struct {
	UserNumber int64  `db:"user_number"`
	UserUUID   string `db:"user_uuid"`
}

const (
	insertUserQuery          = `INSERT INTO users (create_time, update_time, status, registered_email, name, phone_number, profile_picture, linkedin, gender, company_id, company_name, work_email, designation, years_of_experience, google_auth_details, technical_information, mentor_config) VALUES (:create_time, :update_time, :status, :registered_email, :name, :phone_number, :profile_picture, :linkedin, :gender, :company_id, :company_name, :work_email, :designation, :years_of_experience, :google_auth_details, :technical_information, :mentor_config) RETURNING user_number, user_uuid`
	namedUpdateUserBaseQuery = `UPDATE users SET `
	getUserSelectorFields    = `user_number, user_uuid, create_time, update_time, status, registered_email, name, phone_number, profile_picture, linkedin, gender, company_id, company_name, work_email, designation, years_of_experience, google_auth_details, technical_information, mentor_config`

	findByUserNumber      = `SELECT ` + getUserSelectorFields + ` FROM users WHERE user_number = $1`
	findByUserUUID        = `SELECT ` + getUserSelectorFields + ` FROM users WHERE user_uuid = $1`
	findByRegisteredEmail = `SELECT ` + getUserSelectorFields + ` FROM users WHERE registered_email = $1`
	findByWorkEmail       = `SELECT ` + getUserSelectorFields + ` FROM users WHERE work_email = $1`

	selectUserBaseQuery = `SELECT ` + getUserSelectorFields + ` FROM users WHERE `
)

func (r usersRepository) Insert(ctx context.Context, user domain.User) (*domain.User, error) {
	var returning UsersReturning
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, insertUserQuery, map[string]interface{}{
			constants.ParamCreateTime:           user.CreatedAt,
			constants.ParamUpdateTime:           user.UpdatedAt,
			constants.ParamStatus:               user.Status,
			constants.ParamRegisteredEmail:      user.RegisteredEmail,
			constants.ParamName:                 user.Name,
			constants.ParamPhoneNumber:          user.PhoneNumber,
			constants.ParamProfilePicture:       user.ProfilePicture,
			constants.ParamLinkedIn:             user.LinkedIn,
			constants.ParamGender:               user.Gender,
			constants.ParamCompanyID:            user.CompanyID,
			constants.ParamCompanyName:          user.CompanyName,
			constants.ParamWorkEmail:            user.WorkEmail,
			constants.ParamDesignation:          user.Designation,
			constants.ParamYearsOfExperience:    user.YearsOfExperience,
			constants.ParamGoogleOAuth:          user.GoogleAuthJSON,
			constants.ParamTechnicalInformation: user.TechnicalInformationJSON,
			constants.ParamMentorConfig:         user.MentorConfigJSON,
		})
	})
	if err != nil {
		return nil, err
	}

	user.UserNumber = returning.UserNumber
	user.UserUUID = returning.UserUUID
	return &user, nil
}

func (r usersRepository) Update(ctx context.Context, updatedUser *domain.User) error {
	nowTime := time.Now()
	updatedUser.UpdatedAt = &nowTime

	nqb := domain.NewSetQueryBuilder()
	nqb.AddEqualCondition(constants.ParamStatus, updatedUser.Status)
	nqb.AddEqualCondition(constants.ParamUpdateTime, updatedUser.UpdatedAt)

	nqb.AddEqualCondition(constants.ParamName, updatedUser.Name)
	nqb.AddEqualCondition(constants.ParamPhoneNumber, updatedUser.PhoneNumber)
	nqb.AddEqualCondition(constants.ParamLinkedIn, updatedUser.LinkedIn)
	nqb.AddEqualCondition(constants.ParamGender, updatedUser.Gender)

	nqb.AddEqualCondition(constants.ParamCompanyID, updatedUser.CompanyID)
	nqb.AddEqualCondition(constants.ParamCompanyName, updatedUser.CompanyName)
	nqb.AddEqualCondition(constants.ParamWorkEmail, updatedUser.WorkEmail)
	nqb.AddEqualCondition(constants.ParamDesignation, updatedUser.Designation)
	nqb.AddEqualCondition(constants.ParamYearsOfExperience, updatedUser.YearsOfExperience)

	nqb.AddEqualCondition(constants.ParamTechnicalInformation, updatedUser.TechnicalInformationJSON)

	namedParams, namedArgs := nqb.BuildNamed()
	whereCondition := `user_uuid = :user_uuid`
	if updatedUser.UserNumber != 0 {
		whereCondition = `user_number = :user_number`
	}

	updateCompanyQuery := namedUpdateUserBaseQuery + namedParams + ` WHERE ` + whereCondition
	namedArgs[constants.ParamUserUUID] = updatedUser.UserUUID
	namedArgs[constants.ParamUserNumber] = updatedUser.UserNumber

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r usersRepository) BulkUpdate(ctx context.Context, from, to domain.User) error {
	setConditionBuilder := domain.NewSetQueryBuilder()
	setConditionBuilder.AddEqualCondition(constants.ParamStatus, to.Status)
	setConditionBuilder.AddEqualCondition(constants.ParamCompanyID, to.CompanyID)
	setConditionBuilder.AddEqualCondition(constants.ParamCompanyName, to.CompanyName)
	setParams, setNamedArgs := setConditionBuilder.BuildNamed()

	whereConditionBuilder := domain.NewWhereQueryBuilder()
	whereConditionBuilder.AddEqualCondition(constants.ParamUserNumber, from.UserNumber)
	whereConditionBuilder.AddEqualCondition(constants.ParamUserUUID, from.UserUUID)
	whereConditionBuilder.AddEqualCondition(constants.ParamRegisteredEmail, from.RegisteredEmail)
	whereConditionBuilder.AddEqualCondition(constants.ParamCompanyID, from.CompanyID)
	whereConditionBuilder.AddEqualCondition(constants.ParamCompanyName, from.CompanyName)
	whereCondition, whereNamedArgs := whereConditionBuilder.BuildNamed()

	namedArgs := utils.MergeMaps(setNamedArgs, whereNamedArgs)
	updateCompanyQuery := namedUpdateUserBaseQuery + setParams + ` WHERE ` + whereCondition

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r usersRepository) FindByUserNumber(ctx context.Context, userNumber int64) (*domain.User, error) {
	return r.fetchUser(ctx, findByUserNumber, userNumber)
}

func (r usersRepository) FindByUserUUID(ctx context.Context, email string) (*domain.User, error) {
	return r.fetchUser(ctx, findByUserUUID, email)
}

func (r usersRepository) FindByRegisteredEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.fetchUser(ctx, findByRegisteredEmail, email)
}

func (r usersRepository) FindByWorkEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.fetchUser(ctx, findByWorkEmail, email)
}

func (r usersRepository) FetchUserForParams(ctx context.Context, params domain.FetchUserParams) (*domain.User, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserNumber, params.UserNumber)
	qb.AddEqualCondition(constants.ParamUserUUID, params.UserUUID)
	qb.AddEqualCondition(constants.ParamName, params.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, params.PhoneNumber)
	qb.AddEqualCondition(constants.ParamRegisteredEmail, params.RegisteredEmail)
	qb.AddEqualCondition(constants.ParamWorkEmail, params.WorkEmail)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrSearchParamRequired
	}

	baseQuery := `SELECT ` + getUserSelectorFields + ` FROM users`
	getUserByParamsQuery := baseQuery + " WHERE " + conditions

	var user domain.User
	err := r.dbClient.DBGet(ctx, &user, getUserByParamsQuery, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrUserNotFound
	}

	return &user, err
}

func (r usersRepository) FetchUsersForParams(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserNumber, params.UserNumber)
	qb.AddEqualCondition(constants.ParamUserUUID, params.UserUUID)
	qb.AddEqualCondition(constants.ParamStatus, params.Status)
	qb.AddEqualCondition(constants.ParamName, params.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, params.PhoneNumber)
	qb.AddEqualCondition(constants.ParamRegisteredEmail, params.RegisteredEmail)
	qb.AddEqualCondition(constants.ParamWorkEmail, params.WorkEmail)
	qb.AddEqualCondition(constants.ParamCompanyID, params.CompanyID)
	qb.AddEqualCondition(constants.ParamCompanyName, params.CompanyName)
	qb.AddEqualCondition(constants.ParamDesignation, params.Designation)
	qb.AddGreaterEqualCondition(constants.ParamCreateTime, params.CreatedAt)
	qb.AddEqualConditionForJSONB(constants.ParamMentorConfig, constants.ParamStatus, params.MentorConfig.Status)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrSearchParamRequired
	}

	var users []domain.User
	getUsersByParamsQuery := selectUserBaseQuery + conditions

	err := r.dbClient.DBSelect(ctx, &users, getUsersByParamsQuery, args...)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.ErrNoDataFound
	}

	result := domain.Users(users)
	return &result, nil
}

func (r usersRepository) fetchUser(ctx context.Context, query string, key interface{}) (*domain.User, error) {
	user := &domain.User{}

	err := r.dbClient.DBGet(ctx, user, query, key)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
