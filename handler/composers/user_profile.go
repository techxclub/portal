package composers

import (
	"time"

	"github.com/techx/portal/domain"
)

// swagger:model
type UserProfile struct {
	UserNumber        int64      `json:"user_number"`
	UserID            string     `json:"user_id"`
	CreatedAt         *time.Time `json:"created_at"`
	Status            string     `json:"status"`
	Name              string     `json:"name"`
	PhoneNumber       string     `json:"phone_number"`
	PersonalEmail     string     `json:"personal_email"`
	CompanyID         int64      `json:"company_id"`
	CompanyName       string     `json:"company_name"`
	WorkEmail         string     `json:"work_email"`
	Role              string     `json:"role"`
	YearsOfExperience float32    `json:"years_of_experience"`
	LinkedIn          string     `json:"linkedin"`
}

func NewUserProfile(user domain.UserProfile) UserProfile {
	return UserProfile{
		UserNumber:        user.UserIDNum,
		UserID:            user.UserID,
		CreatedAt:         user.CreatedAt,
		Status:            user.Status,
		Name:              user.Name,
		PhoneNumber:       user.PhoneNumber,
		PersonalEmail:     user.PersonalEmail,
		CompanyID:         user.CompanyID,
		CompanyName:       user.CompanyName,
		WorkEmail:         user.WorkEmail,
		Role:              user.Role,
		YearsOfExperience: user.YearsOfExperience,
		LinkedIn:          user.LinkedIn,
	}
}
