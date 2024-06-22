package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
)

type AdminCompanyUpdateRequest struct {
	ID             int64  `json:"id"`
	NormalizedName string `json:"normalized_name"`
	DisplayName    string `json:"display_name"`
	SmallLogo      string `json:"small_logo"`
	BigLogo        string `json:"big_logo"`
	Priority       int64  `json:"priority"`
	Verified       *bool  `json:"verified"`
	Popular        *bool  `json:"popular"`
}

func NewAdminCompanyUpdateRequest(r *http.Request) (*AdminCompanyUpdateRequest, error) {
	var req AdminCompanyUpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r AdminCompanyUpdateRequest) Validate() error {
	return nil
}

func (r AdminCompanyUpdateRequest) ToCompanyProfileParams() *domain.Company {
	return &domain.Company{
		ID:             r.ID,
		NormalizedName: r.NormalizedName,
		DisplayName:    r.DisplayName,
		SmallLogo:      r.SmallLogo,
		BigLogo:        r.BigLogo,
		Priority:       r.Priority,
		Verified:       r.Verified,
		Popular:        r.Popular,
	}
}
