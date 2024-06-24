package request

import (
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

type MentorsListRequest struct {
	Status string
}

func NewMentorsListRequest(_ *http.Request) (*MentorsListRequest, error) {
	return &MentorsListRequest{
		Status: constants.MentorStatusApproved,
	}, nil
}

func (r MentorsListRequest) Validate() error {
	return nil
}

func (r MentorsListRequest) ToMentorProfileParams() domain.FetchUserParams {
	return domain.FetchUserParams{
		MentorConfig: &domain.MentorConfig{
			Status: r.Status,
		},
	}
}
