package request

import (
	"net/http"
	"time"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type ReferralListRequest struct {
	RequesterUserID string
	ProviderUserID  string
	CompanyID       string
	Status          string
	CreatedAt       *time.Time
}

func NewReferralListRequest(r *http.Request) (*ReferralListRequest, error) {
	requesterUserID := r.URL.Query().Get(constants.ParamRequesterID)
	providerUserID := r.URL.Query().Get(constants.ParamProviderID)
	status := r.URL.Query().Get(constants.ParamStatus)
	companyID := r.URL.Query().Get(constants.ParamCompanyID)
	createdAtStr := r.URL.Query().Get(constants.ParamCreatedTime)
	createdAt, err := parseCreatedAt(createdAtStr)
	if err != nil {
		return nil, err
	}

	return &ReferralListRequest{
		RequesterUserID: requesterUserID,
		ProviderUserID:  providerUserID,
		CompanyID:       companyID,
		Status:          status,
		CreatedAt:       createdAt,
	}, nil
}

func (r ReferralListRequest) Validate() error {
	if r.RequesterUserID == "" {
		return errors.ErrRequesterIDRequired
	}
	return nil
}

func (r ReferralListRequest) ToReferralListParams() domain.ReferralParams {
	return domain.ReferralParams{
		RequesterUserID: r.RequesterUserID,
		ProviderUserID:  r.ProviderUserID,
		CompanyID:       utils.ParseInt64WithDefault(r.CompanyID, 0),
		Status:          r.Status,
		CreatedAt:       r.CreatedAt,
	}
}
