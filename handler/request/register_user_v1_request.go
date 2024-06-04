package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type RegisterUserV1Request struct {
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	YearsOfExperience float32 `json:"years_of_experience"`
	PersonalEmail     string  `json:"personal_email"`
	WorkEmail         string  `json:"work_email"`
	PhoneNumber       string  `json:"phone_number"`
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
	if r.YearsOfExperience <= 0 {
		return errors.New("invalid year of experience")
	}

	return nil
}

func (r RegisterUserV1Request) ToUserDetails() domain.User {
	return domain.User{
		FirstName:         r.FirstName,
		LastName:          r.LastName,
		YearsOfExperience: r.YearsOfExperience,
		PersonalEmail:     r.PersonalEmail,
		WorkEmail:         r.WorkEmail,
		PhoneNumber:       r.PhoneNumber,
		LinkedIn:          r.LinkedIn,
	}
}
