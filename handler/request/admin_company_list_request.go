package request

import (
	"net/http"
	"strconv"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type BaseCompanyListRequest struct {
	ID       string
	Name     string
	Priority string
	Verified string
	Popular  string
}

type AdminCompanyListRequest struct {
	BaseCompanyListRequest
}

func NewAdminCompanyListRequest(r *http.Request) (*AdminCompanyListRequest, error) {
	id := r.URL.Query().Get(constants.ParamID)
	name := r.URL.Query().Get(constants.ParamName)
	priority := r.URL.Query().Get(constants.ParamPriority)
	verified := r.URL.Query().Get(constants.ParamVerified)
	popular := r.URL.Query().Get(constants.ParamPopular)

	return &AdminCompanyListRequest{
		BaseCompanyListRequest{
			ID:       id,
			Name:     name,
			Priority: priority,
			Verified: verified,
			Popular:  popular,
		},
	}, nil
}

func (r AdminCompanyListRequest) Validate() error {
	return nil
}

func (r AdminCompanyListRequest) ToFetchCompanyParams() domain.FetchCompanyParams {
	return domain.FetchCompanyParams{
		ID:       parseInt64(r.ID),
		Name:     r.Name,
		Priority: parseInt64(r.Priority),
		Verified: parseBool(r.Verified),
		Popular:  parseBool(r.Popular),
	}
}

func parseBool(str string) *bool {
	if str == "" {
		return nil
	}
	parsedValue, err := strconv.ParseBool(str)
	if err != nil {
		return nil
	}
	return &parsedValue
}

func parseInt64(str string) *int64 {
	if str == "" {
		return nil
	}
	parsedValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil
	}
	return &parsedValue
}
