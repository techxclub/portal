package request

import (
	"net/http"
	"time"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type BaseUserReferralListRequest struct {
	RequesterUserUUID string
	ProviderUserUUID  string
	CompanyID         string
	Status            string
	CreatedAt         *time.Time
}

type AdminUserReferralListRequest struct {
	BaseUserReferralListRequest
}

func NewAdminUserReferralListRequest(r *http.Request) (*AdminUserReferralListRequest, error) {
	requesterUserUUID := r.URL.Query().Get(constants.ParamRequesterID)
	providerUserUUID := r.URL.Query().Get(constants.ParamProviderID)
	status := r.URL.Query().Get(constants.ParamStatus)
	companyID := r.URL.Query().Get(constants.ParamCompanyID)
	createdAtStr := r.URL.Query().Get(constants.ParamCreatedTime)
	createdAt, err := parseCreatedAt(createdAtStr)
	if err != nil {
		return nil, err
	}

	return &AdminUserReferralListRequest{
		BaseUserReferralListRequest{
			RequesterUserUUID: requesterUserUUID,
			ProviderUserUUID:  providerUserUUID,
			CompanyID:         companyID,
			Status:            status,
			CreatedAt:         createdAt,
		},
	}, nil
}

func (r AdminUserReferralListRequest) Validate() error {
	return nil
}

func (r AdminUserReferralListRequest) ToFetchReferralParams() domain.ReferralParams {
	return domain.ReferralParams{
		RequesterUserUUID: r.RequesterUserUUID,
		ProviderUserUUID:  r.ProviderUserUUID,
		CompanyID:         utils.ParseInt64WithDefault(r.CompanyID, 0),
		Status:            r.Status,
		CreatedAt:         r.CreatedAt,
	}
}

func parseCreatedAt(createdAtStr string) (*time.Time, error) {
	if createdAtStr == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
