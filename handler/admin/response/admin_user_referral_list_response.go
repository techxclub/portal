package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type AdminUserReferralListResponse struct {
	Referrals []composers.Referral `json:"referrals"`
}

func NewAdminUserReferralListResponse(ctx context.Context, referrals domain.Referrals) (AdminUserReferralListResponse, composers.HTTPMetadata) {
	return AdminUserReferralListResponse{
		Referrals: composers.ReferralListResponse(ctx, referrals),
	}, composers.HTTPMetadata{}
}
