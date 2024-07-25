package request

import (
	"net/http"
	"strconv"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/utils"
)

type BaseCompanyListRequest struct {
	ID             string
	NormalizedName string
	DisplayName    string
	Priority       string
	Verified       string
	Popular        string
}

type AdminCompanyListRequest struct {
	BaseCompanyListRequest
}

func NewAdminCompanyListRequest(r *http.Request) (*AdminCompanyListRequest, error) {
	id := r.URL.Query().Get(constants.ParamID)
	normalizedName := r.URL.Query().Get(constants.ParamNormalizedName)
	displayName := r.URL.Query().Get(constants.ParamDisplayName)
	priority := r.URL.Query().Get(constants.ParamPriority)
	verified := r.URL.Query().Get(constants.ParamVerified)
	popular := r.URL.Query().Get(constants.ParamPopular)

	return &AdminCompanyListRequest{
		BaseCompanyListRequest{
			ID:             id,
			NormalizedName: normalizedName,
			DisplayName:    displayName,
			Priority:       priority,
			Verified:       verified,
			Popular:        popular,
		},
	}, nil
}

func (r AdminCompanyListRequest) Validate() error {
	return nil
}

func (r AdminCompanyListRequest) ToFetchCompanyParams() domain.FetchCompanyParams {
	return domain.FetchCompanyParams{
		ID:             utils.ParseInt64WithDefault(r.ID, 0),
		NormalizedName: r.NormalizedName,
		DisplayName:    r.DisplayName,
		Priority:       utils.ParseInt64WithDefault(r.Priority, 0),
		Verified:       parseBool(r.Verified),
		Popular:        parseBool(r.Popular),
	}
}

func parseBool(str string) *bool {
	parsedValue, err := strconv.ParseBool(str)
	if err != nil {
		return nil
	}
	return &parsedValue
}
