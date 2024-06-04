package builder

import (
	"context"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type UserInfoBuilder interface {
	NextUserIDNum(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, details domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*domain.User, error)
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

	interUserQuery = `INSERT INTO users (user_id_num, user_id, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role) VALUES (:user_id_num, :user_id, :first_name, :last_name, :years_of_experience, :personal_email, :work_email, :phone_number, :linkedin, :role)`

	getUserSelectorFields = `user_id_num, user_id, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role`

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

func (u usersInfoBuilder) InsertUser(ctx context.Context, user *domain.User) error {
	return u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return u.dbClient.DBNamedExecInTx(ctx, tx, interUserQuery, map[string]interface{}{
			"user_id_num":         user.UserIDNum,
			"user_id":             user.UserID,
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

func (u usersInfoBuilder) CreateUser(ctx context.Context, details domain.User) (*domain.User, error) {
	user := &domain.User{
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

func (u usersInfoBuilder) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User

	err := u.dbClient.DBGet(ctx, &user, getUserByIDQuery, userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u usersInfoBuilder) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user domain.User

	err := u.dbClient.DBGet(ctx, &user, getUserByPhoneQuery, phone)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func getUserIDFromNum(seq int64) string {
	return "U" + strconv.FormatInt(seq, 10)
}
