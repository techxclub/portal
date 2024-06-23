package domain

import (
	"time"

	"github.com/techx/portal/constants"
)

type Users []UserProfile

type UserProfile struct {
	UserIDNum         int64         `db:"user_id_num"` // Only for internal use
	UserID            string        `db:"user_id"`
	CreatedAt         *time.Time    `db:"created_time"`
	Status            string        `db:"status"`
	Name              string        `db:"name"`
	PhoneNumber       string        `db:"phone_number"`
	PersonalEmail     string        `db:"personal_email"`
	CompanyID         int64         `db:"company_id"`
	CompanyName       string        `db:"company_name"`
	WorkEmail         string        `db:"work_email"`
	Role              string        `db:"role"`
	YearsOfExperience float32       `db:"years_of_experience"`
	LinkedIn          string        `db:"linkedin"`
	MentorConfig      *MentorConfig `db:"mentor_config"`
}

type FetchUserParams struct {
	UserIDNum     int64
	UserID        string
	Status        string
	Name          string
	PhoneNumber   string
	PersonalEmail string
	WorkEmail     string
	CompanyID     int64
	CompanyName   string
	Role          string
	CreatedAt     *time.Time
}

func (u UserProfile) IsApproved() bool {
	return u.Status == constants.StatusApproved || u.Status == constants.StatusAutoApproved
}
