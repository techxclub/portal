package request

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type RegisterUserV1Request struct {
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phone_number"`
	RegisteredEmail   string  `json:"registered_email"`
	CompanyName       string  `json:"company_name"`
	Designation       string  `json:"designation"`
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

	if _, err := mail.ParseAddress(r.RegisteredEmail); err != nil {
		return errors.ErrInvalidRegisteredEmail
	}

	if _, err := mail.ParseAddress(r.WorkEmail); err != nil {
		return errors.ErrInvalidWorkEmail
	}

	if r.CompanyName == "" {
		return errors.ErrCompanyRequired
	}

	if !utils.IsValidPhoneNumber(r.PhoneNumber) {
		return errors.ErrInvalidPhoneNumber
	}

	return nil
}

func (r RegisterUserV1Request) ToUserDetails() domain.User {
	return domain.User{
		Status: constants.StatusPendingApproval,
		PersonalInformation: domain.PersonalInformation{
			Name:            r.Name,
			RegisteredEmail: r.RegisteredEmail,
			PhoneNumber:     r.PhoneNumber,
			LinkedIn:        r.LinkedIn,
		},
		ProfessionalInformation: domain.ProfessionalInformation{
			CompanyName:       r.CompanyName,
			YearsOfExperience: r.YearsOfExperience,
			WorkEmail:         r.WorkEmail,
			Designation:       r.Designation,
		},
	}
}
