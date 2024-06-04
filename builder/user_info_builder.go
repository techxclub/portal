package builder

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type UserInfoBuilder interface {
	NextUserID(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, details domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
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
	nextUserIDSeq = `SELECT nextval('users_user_id_seq')`

	interUserQuery = `INSERT INTO users (user_id, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role) VALUES (:user_id, :first_name, :last_name, :years_of_experience, :personal_email, :work_email, :phone_number, :linkedin, :role)`

	getUserSelectorFields = `user_id, first_name, last_name, years_of_experience, personal_email, work_email, phone_number, linkedin, role`
	getUserByIDQuery      = `SELECT ` + getUserSelectorFields + ` FROM users WHERE user_id = $1`
)

func (u usersInfoBuilder) NextUserID(ctx context.Context) (int64, error) {
	var userID int64

	err := u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return u.dbClient.DBGetInTx(ctx, tx, &userID, nextUserIDSeq)
	})
	return userID, err
}

func (u usersInfoBuilder) InsertUser(ctx context.Context, user *domain.User) error {
	return u.dbClient.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return u.dbClient.DBNamedExecInTx(ctx, tx, interUserQuery, map[string]interface{}{
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
		Role:              constants.RoleSeeker,
	}

	userID, err := u.NextUserID(ctx)
	if err != nil {
		return nil, err
	}

	user.UserID = userID
	err = u.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u usersInfoBuilder) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User

	err := u.dbClient.DBGet(ctx, &user, getUserByIDQuery, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
