package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type UsersRepository interface {
	NextUserIDNum(ctx context.Context) (int64, error)
	Insert(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error)
	Update(ctx context.Context, details domain.UserProfile) error
	BulkUpdate(ctx context.Context, from, to domain.UserProfile) error
	FetchUserForParams(ctx context.Context, params domain.FetchUserParams) (*domain.UserProfile, error)
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
	UserID    string     `db:"user_id"`
	CreatedAt *time.Time `db:"created_time"`
}

const (
	nextUserIDNum            = `SELECT nextval('users_user_id_num_seq')`
	interUserQuery           = `INSERT INTO users (user_id_num, created_time, status, name, phone_number, personal_email, company_id, company_name, work_email, role, years_of_experience, mentor_config, linkedin) VALUES (:user_id_num, :created_time, :status, :name, :phone_number, :personal_email, :company_id, :company_name, :work_email, :role, :years_of_experience, :mentor_config, :linkedin) RETURNING user_id, created_time`
	namedUpdateUserBaseQuery = `UPDATE users SET `
	getUserSelectorFields    = `user_id_num, user_id, created_time, status, name, phone_number, personal_email, company_id, company_name, work_email, role, years_of_experience, mentor_config, linkedin`
	selectUserBaseQuery      = `SELECT ` + getUserSelectorFields + ` FROM users WHERE `
)

func (r usersRepository) NextUserIDNum(ctx context.Context) (int64, error) {
	var userID int64

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBGetInTx(ctx, tx, &userID, nextUserIDNum)
	})
	return userID, err
}

func (r usersRepository) Insert(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error) {
	userIDNum, err := r.NextUserIDNum(ctx)
	if err != nil {
		return nil, err
	}

	details.UserIDNum = userIDNum
	details.MentorConfig = &domain.MentorConfig{Status: constants.MentorStatusNotApproved}

	var returning UsersReturning
	err = r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, interUserQuery, map[string]interface{}{
			constants.ParamUserIDNum:         details.UserIDNum,
			constants.ParamCreatedTime:       now,
			constants.ParamStatus:            details.Status,
			constants.ParamName:              details.Name,
			constants.ParamPhoneNumber:       details.PhoneNumber,
			constants.ParamPersonalEmail:     details.PersonalEmail,
			constants.ParamCompanyID:         details.CompanyID,
			constants.ParamCompanyName:       details.CompanyName,
			constants.ParamWorkEmail:         details.WorkEmail,
			constants.ParamRole:              details.Role,
			constants.ParamYearsOfExperience: details.YearsOfExperience,
			constants.ParamMentorConfig:      details.MentorConfig,
			constants.ParamLinkedIn:          details.LinkedIn,
		})
	})
	if err != nil {
		return nil, err
	}

	details.UserID = returning.UserID
	details.CreatedAt = returning.CreatedAt
	return &details, nil
}

func (r usersRepository) Update(ctx context.Context, details domain.UserProfile) error {
	nqb := domain.NewSetQueryBuilder()
	nqb.AddEqualCondition(constants.ParamStatus, details.Status)
	nqb.AddEqualCondition(constants.ParamCompanyName, details.CompanyName)
	nqb.AddEqualCondition(constants.ParamMentorConfig, details.MentorConfig)

	namedParams, namedArgs := nqb.BuildNamed()

	whereCondition := `user_id = :user_id`
	if details.UserIDNum != 0 {
		whereCondition = `user_id_num = :user_id_num`
	}

	updateCompanyQuery := namedUpdateUserBaseQuery + namedParams + ` WHERE ` + whereCondition
	namedArgs[constants.ParamUserID] = details.UserID
	namedArgs[constants.ParamUserIDNum] = details.UserIDNum

	err := r.dbClient.DBRunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r usersRepository) BulkUpdate(ctx context.Context, from, to domain.UserProfile) error {
	setConditionBuilder := domain.NewSetQueryBuilder()
	setConditionBuilder.AddEqualCondition(constants.ParamStatus, to.Status)
	setConditionBuilder.AddEqualCondition(constants.ParamCompanyID, to.CompanyID)
	setConditionBuilder.AddEqualCondition(constants.ParamCompanyName, to.CompanyName)
	setParams, setNamedArgs := setConditionBuilder.BuildNamed()

	whereConditionBuilder := domain.NewWhereQueryBuilder()
	whereConditionBuilder.AddEqualCondition(constants.ParamUserIDNum, from.UserIDNum)
	whereConditionBuilder.AddEqualCondition(constants.ParamUserID, from.UserID)
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

func (r usersRepository) FetchUserForParams(ctx context.Context, params domain.FetchUserParams) (*domain.UserProfile, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserIDNum, params.UserIDNum)
	qb.AddEqualCondition(constants.ParamUserID, params.UserID)
	qb.AddEqualCondition(constants.ParamName, params.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, params.PhoneNumber)
	qb.AddEqualCondition(constants.ParamPersonalEmail, params.PersonalEmail)
	qb.AddEqualCondition(constants.ParamWorkEmail, params.WorkEmail)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrSearchParamRequired
	}

	baseQuery := `SELECT ` + getUserSelectorFields + ` FROM users`
	getUserByParamsQuery := baseQuery + " WHERE " + conditions

	var user domain.UserProfile
	err := r.dbClient.DBGet(ctx, &user, getUserByParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r usersRepository) FetchUsersForParams(ctx context.Context, params domain.FetchUserParams) (*domain.Users, error) {
	qb := domain.NewGetQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserIDNum, params.UserID)
	qb.AddEqualCondition(constants.ParamUserID, params.UserID)
	qb.AddEqualCondition(constants.ParamStatus, params.Status)
	qb.AddEqualCondition(constants.ParamName, params.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, params.PhoneNumber)
	qb.AddEqualCondition(constants.ParamPersonalEmail, params.PersonalEmail)
	qb.AddEqualCondition(constants.ParamWorkEmail, params.WorkEmail)
	qb.AddEqualCondition(constants.ParamCompanyID, params.CompanyID)
	qb.AddEqualCondition(constants.ParamCompanyName, params.CompanyName)
	qb.AddEqualCondition(constants.ParamRole, params.Role)
	qb.AddGreaterEqualCondition(constants.ParamCreatedTime, params.CreatedAt)
	qb.AddEqualConditionForJSONB(constants.ParamMentorConfigStatus, constants.ParamMentorConfig, params.MentorConfig.Status)

	conditions, args := qb.Build()
	if conditions == "" {
		return nil, errors.ErrSearchParamRequired
	}

	var users []domain.UserProfile
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
