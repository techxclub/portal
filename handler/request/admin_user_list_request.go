package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type BaseUserListRequest struct {
	UserID      string
	Status      string
	Name        string
	PhoneNumber string
	CompanyID   string
	CompanyName string
	Role        string
}

type AdminUserListRequest struct {
	BaseUserListRequest
}

func NewAdminUserListRequest(r *http.Request) (*AdminUserListRequest, error) {
	userID := r.URL.Query().Get(constants.ParamUserID)
	status := r.URL.Query().Get(constants.ParamStatus)
	name := r.URL.Query().Get(constants.ParamName)
	phoneNumber := r.URL.Query().Get(constants.ParamPhoneNumber)
	CompanyID := r.URL.Query().Get(constants.ParamCompanyID)
	CompanyName := r.URL.Query().Get(constants.ParamCompanyName)
	Role := r.URL.Query().Get(constants.ParamRole)

	return &AdminUserListRequest{
		BaseUserListRequest{
			UserID:      userID,
			Status:      status,
			Name:        name,
			PhoneNumber: utils.SanitizePhoneNumber(phoneNumber),
			CompanyID:   CompanyID,
			CompanyName: CompanyName,
			Role:        Role,
		},
	}, nil
}

func (r AdminUserListRequest) Validate() error {
	return nil
}

func (r AdminUserListRequest) ToUserProfileParams() domain.UserProfileParams {
	return domain.UserProfileParams{
		UserID:      r.UserID,
		Status:      r.Status,
		PhoneNumber: r.PhoneNumber,
		Name:        r.Name,
		CompanyID:   utils.ParseInt64WithDefault(r.CompanyID, 0),
		CompanyName: r.CompanyName,
		Role:        r.Role,
	}
}
