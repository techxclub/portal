package response

import (
	"context"

	"github.com/techx/portal/domain"
)

type MentorsListResponse struct {
	Mentors []Mentor `json:"mentors"`
}

type Mentor struct {
	Name              string   `json:"name"`
	Role              string   `json:"role"`
	Company           string   `json:"company"`
	CompanyID         int64    `json:"company_id"`
	YearsOfExperience float32  `json:"years_of_experience"`
	Tags              []string `json:"topics"`
	Domain            string   `json:"domain"`
	CalendalyLink     string   `json:"calandly_link"`
}

func NewMentorsListResponse(_ context.Context, users domain.Users) (MentorsListResponse, HTTPMetadata) {
	result := make([]Mentor, 0)
	for _, u := range users {
		result = append(result, Mentor{
			Name:              u.Name,
			Role:              u.Role,
			Company:           u.CompanyName,
			CompanyID:         u.CompanyID,
			YearsOfExperience: u.YearsOfExperience,
			Tags:              u.MentorConfig.Tags,
			Domain:            u.MentorConfig.Domain,
			CalendalyLink:     u.MentorConfig.CalendalyLink,
		})
	}

	return MentorsListResponse{
		Mentors: result,
	}, HTTPMetadata{}
}
