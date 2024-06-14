package handler

import (
	"context"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/handler/request"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
)

func UserProfileHandler(cfg config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewUserProfileRequest,
		func(ctx context.Context, req request.UserProfileRequest) (*domain.UserProfile, error) {
			return serviceRegistry.UserService.GetProfile(ctx, req.ToUserProfileParams())
		},
		func(ctx context.Context, domainObj domain.UserProfile) (response.UserProfile, response.HTTPMetadata) {
			return response.NewUserProfileResponse(ctx, cfg, domainObj)
		},
	)
}
