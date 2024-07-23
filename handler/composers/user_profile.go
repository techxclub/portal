package composers

import (
	"time"

	"github.com/techx/portal/domain"
)

// swagger:model
type UserProfile struct {
	UserNumber              int64                   `json:"user_number"`
	UserUUID                string                  `json:"user_uuid"`
	CreatedAt               *time.Time              `json:"created_at"`
	Status                  string                  `json:"status"`
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
	CompanyID         int64   `json:"company_id"`
	CompanyName       string  `json:"company_name"`
	WorkEmail         string  `json:"work_email"`
	Designation       string  `json:"designation"`
	YearsOfExperience float32 `json:"years_of_experience"`
}

type TechnicalInformation struct {
	Domain string   `json:"domain"`
	Skills []string `json:"skills"`
}

func NewUserProfile(user domain.User) UserProfile {
	technicalInformation := user.TechnicalInformation()

	return UserProfile{
		UserNumber: user.UserNumber,
		UserUUID:   user.UserUUID,
		CreatedAt:  user.CreatedAt,
		Status:     user.Status,
		PersonalInformation: PersonalInformation{
			Name:            user.Name,
			PhoneNumber:     user.PhoneNumber,
			RegisteredEmail: user.RegisteredEmail,
			ProfilePicture:  user.ProfilePicture,
			LinkedIn:        user.LinkedIn,
			Gender:          user.Gender,
		},
		ProfessionalInformation: ProfessionalInformation{
			CompanyID:         user.CompanyID,
			CompanyName:       user.CompanyName,
			WorkEmail:         user.WorkEmail,
			Designation:       user.Designation,
			YearsOfExperience: user.YearsOfExperience,
		},
		TechnicalInformation: TechnicalInformation{
			Domain: technicalInformation.Domain,
			Skills: technicalInformation.Skills,
		},
	}
}
