package response

import (
	"context"
	"time"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type UserProfile struct {
	UserID            string     `json:"user_id"`
	CreatedAt         *time.Time `json:"created_at"`
	Name              string     `json:"name"`
	Company           string     `json:"company"`
	YearsOfExperience float32    `json:"years_of_experience"`
	PersonalEmail     string     `json:"personal_email"`
	WorkEmail         string     `json:"work_email"`
	PhoneNumber       string     `json:"phone_number"`
	LinkedIn          string     `json:"linkedin"`
	Role              string     `json:"role"`
}

func NewUserProfileResponse(_ context.Context, _ config.Config, user domain.UserProfile) (UserProfile, HTTPMetadata) {
	respBody := getUserProfile(user)
	return respBody, HTTPMetadata{}
}

func getUserProfile(user domain.UserProfile) UserProfile {
	return UserProfile{
		UserID:            user.UserID,
		CreatedAt:         user.CreatedAt,
		Name:              user.Name,
		Company:           user.Company,
		YearsOfExperience: user.YearsOfExperience,
		PersonalEmail:     user.PersonalEmail,
		WorkEmail:         user.WorkEmail,
		PhoneNumber:       user.PhoneNumber,
		LinkedIn:          user.LinkedIn,
		Role:              user.Role,
	}
}
