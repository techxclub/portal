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

func MentorsListHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return Handler(
		request.NewMentorsListRequest,
		func(ctx context.Context, req request.MentorsListRequest) (*domain.Users, error) {
			return serviceRegistry.UserService.GetUsers(ctx, req.ToMentorProfileParams())
		},
		response.NewMentorsListResponse,
	)
}
