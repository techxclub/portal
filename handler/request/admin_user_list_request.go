package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type BaseUserListRequest struct {
	UserNumber  string
	UserUUID    string
	Status      string
	Name        string
	PhoneNumber string
	CompanyID   string
	CompanyName string
	Designation string
}

type AdminUserListRequest struct {
	BaseUserListRequest
}

func NewAdminUserListRequest(r *http.Request) (*AdminUserListRequest, error) {
	userNumber := r.URL.Query().Get(constants.ParamUserNumber)
	userID := r.URL.Query().Get(constants.ParamUserUUID)
	status := r.URL.Query().Get(constants.ParamStatus)
	name := r.URL.Query().Get(constants.ParamName)
	phoneNumber := r.URL.Query().Get(constants.ParamPhoneNumber)
	companyID := r.URL.Query().Get(constants.ParamCompanyID)
	companyName := r.URL.Query().Get(constants.ParamCompanyName)
	designation := r.URL.Query().Get(constants.ParamDesignation)

	return &AdminUserListRequest{
		BaseUserListRequest{
			UserNumber:  userNumber,
			UserUUID:    userID,
			Status:      status,
			Name:        name,
			PhoneNumber: utils.SanitizePhoneNumber(phoneNumber),
			CompanyID:   companyID,
			CompanyName: companyName,
			Designation: designation,
		},
	}, nil
}

func (r AdminUserListRequest) Validate() error {
	return nil
}

func (r AdminUserListRequest) ToFetchUserParams() domain.FetchUserParams {
	return domain.FetchUserParams{
		UserNumber:  utils.ParseInt64WithDefault(r.UserNumber, 0),
		UserUUID:    r.UserUUID,
		Status:      r.Status,
		PhoneNumber: r.PhoneNumber,
		Name:        r.Name,
		CompanyID:   utils.ParseInt64WithDefault(r.CompanyID, 0),
		CompanyName: r.CompanyName,
		Designation: r.Designation,
	}
}
