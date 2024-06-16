package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UsersRepository interface {
	NextUserIDNum(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error)
	UpdateUser(ctx context.Context, details domain.UserProfile) error
	GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error)
	GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error)
}

type usersRepository struct {
	dbClient *db.Repository
}

func NewUsersRepository(userDB *db.Repository) UsersRepository {
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
	interUserQuery           = `INSERT INTO users (user_id_num, created_time, status, name, phone_number, personal_email, company, work_email, role, years_of_experience, linkedin) VALUES (:user_id_num, :created_time, :status, :name, :phone_number, :personal_email, :company, :work_email, :role, :years_of_experience, :linkedin) RETURNING user_id, created_time`
	namedUpdateUserBaseQuery = `UPDATE users SET `
	getUserSelectorFields    = `user_id_num, user_id, created_time, status, name, phone_number, personal_email, company, work_email, role, years_of_experience, linkedin`
	selectUserBaseQuery      = `SELECT ` + getUserSelectorFields + ` FROM users WHERE `
)

func (r usersRepository) NextUserIDNum(ctx context.Context) (int64, error) {
	var userID int64

	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBGetInTx(ctx, tx, &userID, nextUserIDNum)
	})
	return userID, err
}

func (r usersRepository) CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error) {
	userIDNum, err := r.NextUserIDNum(ctx)
	if err != nil {
		return nil, err
	}

	details.UserIDNum = userIDNum

	var returning UsersReturning
	err = r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return r.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, interUserQuery, map[string]interface{}{
			"user_id_num":         details.UserIDNum,
			"created_time":        now,
			"status":              details.Status,
			"name":                details.Name,
			"phone_number":        details.PhoneNumber,
			"personal_email":      details.PersonalEmail,
			"company":             details.Company,
			"work_email":          details.WorkEmail,
			"role":                details.Role,
			"years_of_experience": details.YearsOfExperience,
			"linkedin":            details.LinkedIn,
		})
	})
	if err != nil {
		return nil, err
	}

	details.UserID = returning.UserID
	details.CreatedAt = returning.CreatedAt
	return &details, nil
}

func (r usersRepository) UpdateUser(ctx context.Context, details domain.UserProfile) error {
	nqb := domain.NewNamedQueryBuilder()
	nqb.AddEqualCondition(constants.ParamStatus, details.Status)
	nqb.AddEqualCondition(constants.ParamCompany, details.Company)
	nqb.AddEqualCondition(constants.ParamRole, details.Role)

	namedParams, namedArgs := nqb.BuildNamedConditions()

	whereCondition := `user_id = :user_id`
	if details.UserIDNum != 0 {
		whereCondition = `user_id_num = :user_id_num`
	}

	updateCompanyQuery := namedUpdateUserBaseQuery + namedParams + ` WHERE ` + whereCondition
	namedArgs[constants.ParamUserID] = details.UserID
	namedArgs[constants.ParamUserIDNum] = details.UserIDNum

	err := r.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.dbClient.DBNamedExecInTx(ctx, tx, updateCompanyQuery, namedArgs)
	})

	return err
}

func (r usersRepository) GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error) {
	qb := domain.NewQueryBuilder()
	if params.UserIDNum != 0 {
		qb.AddEqualCondition(constants.ParamUserIDNum, params.UserIDNum)
	}
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

func (r usersRepository) GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error) {
	qb := domain.NewQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserIDNum, params.UserID)
	qb.AddEqualCondition(constants.ParamUserID, params.UserID)
	qb.AddEqualCondition(constants.ParamStatus, params.Status)
	qb.AddEqualCondition(constants.ParamName, params.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, params.PhoneNumber)
	qb.AddEqualCondition(constants.ParamPersonalEmail, params.PersonalEmail)
	qb.AddEqualCondition(constants.ParamWorkEmail, params.WorkEmail)
	qb.AddEqualCondition(constants.ParamCompany, params.Company)
	qb.AddEqualCondition(constants.ParamRole, params.Role)
	if params.CreatedAt != nil {
		qb.AddGreaterEqualCondition(constants.ParamCreatedTime, params.CreatedAt)
	}

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
