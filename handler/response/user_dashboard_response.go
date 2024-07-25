package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type UserDashboardResponse struct {
	ReferralOverview composers.ReferralOverview `json:"referral_overview"`
}

func NewUserDashboardResponse(_ context.Context, userReferrals domain.UserReferrals) (UserDashboardResponse, composers.HTTPMetadata) {
	return UserDashboardResponse{ReferralOverview: composers.NewReferralOverview(userReferrals)}, composers.HTTPMetadata{}
}
