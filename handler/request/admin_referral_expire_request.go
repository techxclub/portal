package request

import (
	"net/http"

	"github.com/techx/portal/domain"
)

type AdminExpireReferralRequest struct{}

func NewAdminExpireReferralRequest(_ *http.Request) (*AdminExpireReferralRequest, error) {
	return &AdminExpireReferralRequest{}, nil
}

func (r AdminExpireReferralRequest) Validate() error {
	return nil
}

func (r AdminExpireReferralRequest) ToExpireReferralParams() *domain.Referral {
	return &domain.Referral{}
}
