package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

// swagger:model
type CompanyUsersListResponse struct {
	Users []composers.CompanyUser `json:"users"`
}

func NewCompanyUsersListResponse(ctx context.Context, users domain.CompanyUsersService) (CompanyUsersListResponse, composers.HTTPMetadata) {
	return CompanyUsersListResponse{
		Users: composers.GetCompanyUsers(ctx, users),
	}, composers.HTTPMetadata{}
}
