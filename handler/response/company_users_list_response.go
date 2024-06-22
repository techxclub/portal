package response

import (
	"context"

	"github.com/techx/portal/domain"
)

// swagger:model
type CompanyUsersListResponse struct {
	Users []CompanyUser `json:"users"`
}

type CompanyUser struct {
	Name              string  `json:"name"`
	UserID            string  `json:"user_id"`
	CompanyName       string  `json:"company_name"`
	Role              string  `json:"role"`
	YearsOfExperience float32 `json:"years_of_experience"`
}

func NewCompanyUsersListResponse(_ context.Context, users domain.Users) (CompanyUsersListResponse, HTTPMetadata) {
	result := make([]CompanyUser, 0)
	for _, u := range users {
		result = append(result, CompanyUser{
			UserID:            u.UserID,
			Name:              u.Name,
			CompanyName:       u.CompanyName,
			Role:              u.Role,
			YearsOfExperience: u.YearsOfExperience,
		})
	}

	return CompanyUsersListResponse{
		Users: result,
	}, HTTPMetadata{}
}
