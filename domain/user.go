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
	Status            string     `db:"status"`
	Name              string     `db:"name"`
	PhoneNumber       string     `db:"phone_number"`
	PersonalEmail     string     `db:"personal_email"`
	Company           string     `db:"company"`
	WorkEmail         string     `db:"work_email"`
	Role              string     `db:"role"`
	YearsOfExperience float32    `db:"years_of_experience"`
	LinkedIn          string     `db:"linkedin"`
}

type UserProfileParams struct {
	UserID      string
	Status      string
	Name        string
	PhoneNumber string
	Company     string
	Role        string
}

func (u UserProfile) IsApproved() bool {
	return u.Status == constants.StatusApproved || u.Status == constants.StatusAutoApproved
}

func (p UserProfileParams) GetQueryConditions() (string, []interface{}) {
	qb := NewQueryBuilder()
	qb.AddEqualCondition(constants.ParamUserID, p.UserID)
	qb.AddEqualCondition(constants.ParamStatus, p.Status)
	qb.AddEqualCondition(constants.ParamName, p.Name)
	qb.AddEqualCondition(constants.ParamPhoneNumber, p.PhoneNumber)
	qb.AddEqualCondition(constants.ParamCompany, p.Company)
	qb.AddEqualCondition(constants.ParamRole, p.Role)

	return qb.Build()
}
