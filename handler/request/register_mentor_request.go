package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type RegisterMentorRequest struct {
	UserID        string   `json:"-"`
	IsMentor      bool     `json:"is_mentor"`
	IsPreApproved bool     `json:"is_pre_approved"`
	Status        string   `json:"status"`
	CalendalyLink string   `json:"calendaly_link,omitempty"`
	Description   string   `json:"description,omitempty"`
	IsAvailable   bool     `json:"is_available,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Domain        string   `json:"domain,omitempty"`
}

func NewRegisterMentorRequest(r *http.Request) (*RegisterMentorRequest, error) {
	var req RegisterMentorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	req.UserID = r.Header.Get(constants.HeaderXUserID)
	return &req, nil
}

func (r RegisterMentorRequest) Validate() error {
	if r.CalendalyLink == "" {
		return errors.ErrCalendalyLinkRequired
	}

	if len(r.Tags) == 0 {
		return errors.ErrTagsRequired
	}

	if len(r.Tags) > 5 {
		return errors.ErrTagsLimitExceededByFive
	}

	if r.Domain == "" {
		return errors.ErrDomainRequired
	}

	return nil
}

func (r RegisterMentorRequest) ToMentorDetails() domain.UserProfile {
	return domain.UserProfile{
		UserID: r.UserID,
		MentorConfig: &domain.MentorConfig{
			IsMentor:      r.IsMentor,
			IsPreApproved: r.IsPreApproved,
			Status:        r.Status,
			CalendalyLink: r.CalendalyLink,
			Description:   r.Description,
			IsAvailable:   r.IsAvailable,
			Tags:          r.Tags,
			Domain:        r.Domain,
		},
	}
}
