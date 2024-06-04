package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

// swagger:model
type UserDetailsResponse struct {
	UserID            int64   `json:"user_id"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	YearsOfExperience float32 `json:"years_of_experience"`
	PersonalEmail     string  `json:"personal_email"`
	WorkEmail         string  `json:"work_email"`
	PhoneNumber       string  `json:"phone_number"`
	LinkedIn          string  `json:"linkedin"`
	Role              string  `json:"role"`
}

func NewGetUserResponse(_ context.Context, _ config.Config, user domain.User) (UserDetailsResponse, http.Header) {
	respBody := UserDetailsResponse{
		UserID:            user.UserID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		YearsOfExperience: user.YearsOfExperience,
		PersonalEmail:     user.PersonalEmail,
		WorkEmail:         user.WorkEmail,
		PhoneNumber:       user.PhoneNumber,
		LinkedIn:          user.LinkedIn,
		Role:              user.Role,
	}

	return respBody, http.Header{}
}
