package request

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
)

type UserUpdateProfileRequest struct {
	UserUUID                string                  `json:"-"`
	PersonalInformation     PersonalInformation     `json:"personal_information"`
	ProfessionalInformation ProfessionalInformation `json:"professional_information"`
	TechnicalInformation    TechnicalInformation    `json:"technical_information"`
}

type PersonalInformation struct {
	Name            string `json:"name"`
	PhoneNumber     string `json:"phone_number"`
	RegisteredEmail string `json:"registered_email"`
	ProfilePicture  string `json:"profile_picture"`
	LinkedIn        string `json:"linkedin"`
	Gender          string `json:"gender"`
}

type ProfessionalInformation struct {
	CompanyName       string  `json:"company_name"`
	WorkEmail         string  `json:"work_email"`
	Designation       string  `json:"designation"`
	YearsOfExperience float32 `json:"years_of_experience"`
}

type TechnicalInformation struct {
	Domain string   `json:"domain"`
	Skills []string `json:"skills"`
}

func NewUserUpdateProfileRequest(r *http.Request) (*UserUpdateProfileRequest, error) {
	var req UserUpdateProfileRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	req.UserUUID = r.Header.Get(constants.HeaderXUserUUID)
	return &req, nil
}

func (r UserUpdateProfileRequest) Validate() error {
	if r.UserUUID == "" {
		return errors.ErrInvalidUserID
	}

	return nil
}

func (r UserUpdateProfileRequest) ToUserDomainObject() domain.User {
	user := domain.User{
		UserUUID: r.UserUUID,
		PersonalInformation: domain.PersonalInformation{
			Name:            r.PersonalInformation.Name,
			PhoneNumber:     r.PersonalInformation.PhoneNumber,
			RegisteredEmail: r.PersonalInformation.RegisteredEmail,
			ProfilePicture:  r.PersonalInformation.ProfilePicture,
			LinkedIn:        r.PersonalInformation.LinkedIn,
			Gender:          r.PersonalInformation.Gender,
		},
		ProfessionalInformation: domain.ProfessionalInformation{
			CompanyName:       r.ProfessionalInformation.CompanyName,
			WorkEmail:         r.ProfessionalInformation.WorkEmail,
			Designation:       r.ProfessionalInformation.Designation,
			YearsOfExperience: r.ProfessionalInformation.YearsOfExperience,
		},
	}

	user.SetTechnicalInformation(domain.TechnicalInformation{
		Domain: r.TechnicalInformation.Domain,
		Skills: r.TechnicalInformation.Skills,
	})
	return user
}
