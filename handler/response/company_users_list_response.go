package response

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type CompanyUsersListResponse struct {
	Users []CompanyUser `json:"users"`
}

type CompanyUser struct {
	Name    string `json:"name"`
	UserID  string `json:"user_id"`
	Company string `json:"company"`
}

func NewCompanyUsersListResponse(_ context.Context, _ config.Config, users domain.Users) (CompanyUsersListResponse, HTTPMetadata) {
	result := make([]CompanyUser, 0)
	for _, u := range users {
		result = append(result, CompanyUser{
			UserID:  u.UserID,
			Name:    u.Name,
			Company: u.Company,
		})
	}

	return CompanyUsersListResponse{
		Users: result,
	}, HTTPMetadata{}
}
