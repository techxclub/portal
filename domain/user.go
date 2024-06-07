package domain

import (
	"time"

	"github.com/techx/portal/constants"
)

type Users []UserProfile

type UserProfile struct {
	UserIDNum         int64      `db:"user_id_num"` // Only for internal use
	UserID            string     `db:"user_id"`
	CreatedAt         *time.Time `db:"created_time"`
	FirstName         string     `db:"first_name"`
	LastName          string     `db:"last_name"`
	YearsOfExperience float32    `db:"years_of_experience"`
	PersonalEmail     string     `db:"personal_email"`
	WorkEmail         string     `db:"work_email"`
	PhoneNumber       string     `db:"phone_number"`
	LinkedIn          string     `db:"linkedin"`
	Role              string     `db:"role"`
}

type UserProfileParams struct {
	UserID      string
	PhoneNumber string
	FirstName   string
	LastName    string
	Company     string
	Role        string
}

func (p UserProfileParams) ToMap() map[string]string {
	return map[string]string{
		constants.ParamUserID:      p.UserID,
		constants.ParamPhoneNumber: p.PhoneNumber,
		constants.ParamFirstName:   p.FirstName,
		constants.ParamLastName:    p.LastName,
		constants.ParamCompany:     p.Company,
		constants.ParamRole:        p.Role,
	}
}
