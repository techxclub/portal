package composers

import (
	"time"

	"github.com/techx/portal/domain"
)

// swagger:model
type UserProfile struct {
	UserID            string     `json:"user_id"`
	CreatedAt         *time.Time `json:"created_at"`
	Status            string     `json:"status"`
	Name              string     `json:"name"`
	PhoneNumber       string     `json:"phone_number"`
	PersonalEmail     string     `json:"personal_email"`
	Company           string     `json:"company"`
	WorkEmail         string     `json:"work_email"`
	Role              string     `json:"role"`
	YearsOfExperience float32    `json:"years_of_experience"`
	LinkedIn          string     `json:"linkedin"`
}

func NewUserProfile(user domain.UserProfile) UserProfile {
	return UserProfile{
		UserID:            user.UserID,
		CreatedAt:         user.CreatedAt,
		Status:            user.Status,
		Name:              user.Name,
		PhoneNumber:       user.PhoneNumber,
		PersonalEmail:     user.PersonalEmail,
		Company:           user.Company,
		WorkEmail:         user.WorkEmail,
		Role:              user.Role,
		YearsOfExperience: user.YearsOfExperience,
		LinkedIn:          user.LinkedIn,
	}
}
