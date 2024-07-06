package response

import (
	"context"
	"time"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/composers"
)

type AdminUserReferralListResponse struct {
	Referrals []composers.Referral `json:"referrals"`
}

type Referral struct {
	ID              int64      `db:"id"`
	CompanyID       int64      `db:"company_id"`
	RequesterUserID string     `db:"requester_user_id"`
	ProviderUserID  string     `db:"provider_user_id"`
	JobLink         string     `db:"job_link"`
	Status          string     `db:"status"`
	CreatedAt       *time.Time `db:"created_time"`
}

func NewAdminUserReferralListResponse(ctx context.Context, referrals domain.Referrals) (AdminUserReferralListResponse, HTTPMetadata) {
	return AdminUserReferralListResponse{
		Referrals: composers.ReferralListResponse(ctx, referrals),
	}, HTTPMetadata{}
}
