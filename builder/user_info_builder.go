package builder

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UserInfoBuilder interface {
	NextUserIDNum(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error)
	GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error)
	GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error)
}

type usersInfoBuilder struct {
	dbClient *db.Repository
}

func NewUsersInfoBuilder(userDB *db.Repository) UserInfoBuilder {
	return &usersInfoBuilder{
		dbClient: userDB,
	}
}

const (
	nextUserIDNum = `SELECT nextval('users_user_id_num_seq')`

	interUserQuery = `INSERT INTO users (user_id_num, user_id, created_time, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role) VALUES (:user_id_num, :user_id, :created_time, :first_name, :last_name, :years_of_experience, :personal_email, :work_email, :phone_number, :linkedin, :role)`

	getUserSelectorFields = `user_id_num, user_id, created_time, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role`

	getUserByPhoneQuery = `SELECT ` + getUserSelectorFields + ` FROM users WHERE phone_number = $1`
	getUserByIDQuery    = `SELECT ` + getUserSelectorFields + ` FROM users WHERE user_id = $1`
)

func (u usersInfoBuilder) NextUserIDNum(ctx context.Context) (int64, error) {
	var userID int64

	err := u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return u.dbClient.DBGetInTx(ctx, tx, &userID, nextUserIDNum)
	})
	return userID, err
}

func (u usersInfoBuilder) InsertUser(ctx context.Context, user *domain.UserProfile) error {
	return u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		now := time.Now()
		return u.dbClient.DBNamedExecInTx(ctx, tx, interUserQuery, map[string]interface{}{
			"user_id_num":         user.UserIDNum,
			"user_id":             user.UserID,
			"created_time":        now,
			"first_name":          user.FirstName,
			"last_name":           user.LastName,
			"years_of_experience": user.YearsOfExperience,
			"personal_email":      user.PersonalEmail,
			"work_email":          user.WorkEmail,
			"phone_number":        user.PhoneNumber,
			"linkedin":            user.LinkedIn,
			"role":                user.Role,
		})
	})
}

func (u usersInfoBuilder) CreateUser(ctx context.Context, details domain.UserProfile) (*domain.UserProfile, error) {
	user := &domain.UserProfile{
		FirstName:         details.FirstName,
		LastName:          details.LastName,
		YearsOfExperience: details.YearsOfExperience,
		PersonalEmail:     details.PersonalEmail,
		WorkEmail:         details.WorkEmail,
		PhoneNumber:       details.PhoneNumber,
		LinkedIn:          details.LinkedIn,
		Role:              constants.RoleViewer,
	}

	userIDNum, err := u.NextUserIDNum(ctx)
	if err != nil {
		return nil, err
	}

	user.UserIDNum = userIDNum
	user.UserID = getUserIDFromNum(userIDNum)
	err = u.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u usersInfoBuilder) GetUserForParams(ctx context.Context, params domain.UserProfileParams) (*domain.UserProfile, error) {
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

func (u usersInfoBuilder) GetUsersForParams(ctx context.Context, params domain.UserProfileParams) (*domain.Users, error) {
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
		return nil, errors.New("no users found")
	}

	result := domain.Users(users)
	return &result, nil
}

func getUserIDFromNum(seq int64) string {
	return "U" + strconv.FormatInt(seq, 10)
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
		return "", nil, errors.New("no search parameters provided")
	}

	baseQuery := `SELECT ` + getUserSelectorFields + ` FROM users WHERE `
	query := baseQuery + strings.Join(conditions, " AND ")
	return query, args, nil
}
