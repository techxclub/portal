package request

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type RegisterUserV1Request struct {
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phone_number"`
	PersonalEmail     string  `json:"personal_email"`
	Company           string  `json:"company"`
	Role              string  `json:"role"`
	YearsOfExperience float32 `json:"years_of_experience"`
	WorkEmail         string  `json:"work_email"`
	LinkedIn          string  `json:"linkedin"`
}

func NewRegisterUserV1Request(r *http.Request) (*RegisterUserV1Request, error) {
	var req RegisterUserV1Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (r RegisterUserV1Request) Validate() error {
	if r.Name == "" {
		return errors.ErrNameRequired
	}

	if r.YearsOfExperience <= 0 {
		return errors.ErrInvalidYearsOfExperience
	}

	if _, err := mail.ParseAddress(r.PersonalEmail); err != nil {
		return errors.ErrInvalidPersonalEmail
	}

	if _, err := mail.ParseAddress(r.WorkEmail); err != nil {
		return errors.ErrInvalidWorkEmail
	}

	if r.Company == "" {
		return errors.ErrCompanyRequired
	}

	return IsValidPhoneNumber(r.PhoneNumber)
}

func (r RegisterUserV1Request) ToUserDetails() domain.UserProfile {
	return domain.UserProfile{
		Name:              r.Name,
		Company:           r.Company,
		YearsOfExperience: r.YearsOfExperience,
		PersonalEmail:     r.PersonalEmail,
		WorkEmail:         r.WorkEmail,
		PhoneNumber:       r.PhoneNumber,
		LinkedIn:          r.LinkedIn,
		Role:              r.Role,
		Status:            constants.StatusPendingApproval,
	}
}
