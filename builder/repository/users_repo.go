package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UsersRepo interface {
	NextUserIDNum(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error)
	GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error)
	GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error)
	GetCompanies(ctx context.Context) (*domain.Companies, error)
}

type usersRepo struct {
	dbClient *db.Repository
}

func NewUsersRepo(userDB *db.Repository) UsersRepo {
	return &usersRepo{
		dbClient: userDB,
	}
}

type UsersReturning struct {
	UserID string `db:"user_id"`
}

const (
	nextUserIDNum = `SELECT nextval('users_user_id_num_seq')`

	interUserQuery = `INSERT INTO users (user_id_num, created_time, status, name, phone_number, personal_email, company, work_email, role, years_of_experience, linkedin) VALUES (:user_id_num, :created_time, :status, :name, :phone_number, :personal_email, :company, :work_email, :role, :years_of_experience, :linkedin) RETURNING user_id`

	getUserSelectorFields = `user_id_num, user_id, created_time, status, name, phone_number, personal_email, company, work_email, role, years_of_experience, linkedin`

	getDistinctCompanies = `SELECT DISTINCT company as name FROM users`
)

func (u usersRepo) NextUserIDNum(ctx context.Context) (int64, error) {
	var userID int64

	err := u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return u.dbClient.DBGetInTx(ctx, tx, &userID, nextUserIDNum)
	})
	return userID, err
}

func (u usersRepo) CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error) {
	userIDNum, err := u.NextUserIDNum(ctx)
	if err != nil {
		return nil, err
	}

	details.UserIDNum = userIDNum

	var returning UsersReturning
	err = u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return u.dbClient.DBNamedExecReturningInTx(ctx, tx, &returning, interUserQuery, map[string]interface{}{
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
	return &details, nil
}

func (u usersRepo) GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error) {
	getUserByParamsQuery, args, err := getQueryForParams(params)
	if err != nil {
		return nil, err
	}

	var user domain.UserProfile
	err = u.dbClient.DBGet(ctx, &user, getUserByParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u usersRepo) GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error) {
	getUsersByParamsQuery, args, err := getQueryForParams(params)
	if err != nil {
		return nil, err
	}

	var users []domain.UserProfile
	err = u.dbClient.DBSelect(ctx, &users, getUsersByParamsQuery, args...)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.ErrUserNotFound
	}

	result := domain.Users(users)
	return &result, nil
}

func getQueryForParams(params domain.UserProfileParams) (string, []interface{}, error) {
	counter := 1
	args := make([]interface{}, 0)
	conditions := make([]string, 0)
	for key, value := range params.ToMap() {
		if value == "" {
			continue
		}

		condition := fmt.Sprintf("%s = $%d", key, counter)
		conditions = append(conditions, condition)
		args = append(args, value)
		counter++
	}

	if len(conditions) == 0 {
		return "", nil, errors.ErrSearchParamRequired
	}

	baseQuery := `SELECT ` + getUserSelectorFields + ` FROM users WHERE `
	query := baseQuery + strings.Join(conditions, " AND ")
	return query, args, nil
}

func (u usersRepo) GetCompanies(ctx context.Context) (*domain.Companies, error) {
	var companies []domain.Company
	err := u.dbClient.DBSelect(ctx, &companies, getDistinctCompanies)
	if err != nil {
		return nil, err
	}

	if len(companies) == 0 {
		return nil, errors.ErrCompanyNotFound
	}

	result := domain.Companies(companies)
	return &result, nil
}
