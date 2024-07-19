package response

import (
	"context"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type ReferralListResponse struct {
	Referrals []composers.Referral `json:"referrals"`
}

func NewReferralListResponse(ctx context.Context, referrals domain.Referrals) (ReferralListResponse, HTTPMetadata) {
	return ReferralListResponse{
		Referrals: composers.ReferralListResponse(ctx, referrals),
	}, HTTPMetadata{}
}
