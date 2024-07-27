package composers

import (
	"github.com/techx/portal/domain"
)

type ReferralOverview struct {
	RequestReceivedCount int `json:"request_received_count"`
	RequestSentCount     int `json:"request_sent_count"`
}

func NewReferralOverview(userReferrals domain.UserReferrals) ReferralOverview {
	return ReferralOverview{
		RequestSentCount:     len(*userReferrals.RequestedReferrals),
		RequestReceivedCount: len(*userReferrals.ProvidedReferrals),
	}
}
