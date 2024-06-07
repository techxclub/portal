package response

import (
	"context"
	"net/http"
	"time"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type UserProfile struct {
	UserID            string     `json:"user_id"`
	CreatedAt         *time.Time `json:"created_at"`
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	YearsOfExperience float32    `json:"years_of_experience"`
	PersonalEmail     string     `json:"personal_email"`
	WorkEmail         string     `json:"work_email"`
	PhoneNumber       string     `json:"phone_number"`
	LinkedIn          string     `json:"linkedin"`
	Role              string     `json:"role"`
}

func NewUserProfileResponse(_ context.Context, _ config.Config, user domain.UserProfile) (UserProfile, http.Header) {
	respBody := getUserProfile(user)
	return respBody, http.Header{}
}

func getUserProfile(user domain.UserProfile) UserProfile {
	return UserProfile{
		UserID:            user.UserID,
		CreatedAt:         user.CreatedAt,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		YearsOfExperience: user.YearsOfExperience,
		PersonalEmail:     user.PersonalEmail,
		WorkEmail:         user.WorkEmail,
		PhoneNumber:       user.PhoneNumber,
		LinkedIn:          user.LinkedIn,
		Role:              user.Role,
	}
}
