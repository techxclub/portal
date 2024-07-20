package domain

import (
	"time"

	"github.com/techx/portal/constants"
)

type Users []User

type User struct {
	UserNumber int64      `db:"user_number"` // Only for internal use
	UserUUID   string     `db:"user_uuid"`
	CreatedAt  *time.Time `db:"created_time"`
	Status     string     `db:"status"`

	PersonalDetails
	ProfessionalDetails

	GoogleAuthJSON      JSONWrapper[GoogleOAuthDetails] `db:"google_auth_details"`
	TechnicalSkillsJSON JSONWrapper[TechnicalSkills]    `db:"technical_skills"`
	MentorConfigJSON    JSONWrapper[MentorConfig]       `db:"mentor_config"`
}

type PersonalDetails struct {
	Name            string `db:"name"`
	PhoneNumber     string `db:"phone_number"`
	RegisteredEmail string `db:"registered_email"`
	ProfilePicture  string `db:"profile_picture"`
	LinkedIn        string `db:"linkedin"`
	Gender          string `db:"gender"`
}

type ProfessionalDetails struct {
	CompanyID         int64   `db:"company_id"`
	CompanyName       string  `db:"company_name"`
	WorkEmail         string  `db:"work_email"`
	Designation       string  `db:"designation"`
	YearsOfExperience float32 `db:"years_of_experience"`
}

type GoogleOAuthDetails struct {
	Email        string    `json:"email"`
	TokenType    string    `json:"token_type"`
	AccessToken  string    `json:"access_token"`
	IDToken      string    `json:"id_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

type TechnicalSkills struct {
	Domain string   `json:"domain"`
	Skills []string `json:"skills"`
}

type MentorConfig struct {
	IsMentor      bool     `json:"is_mentor"`
	IsPreApproved bool     `json:"is_pre_approved"`
	Status        string   `json:"status"`
	CalendalyLink string   `json:"calendaly_link,omitempty"`
	Description   string   `json:"description,omitempty"`
	IsAvailable   bool     `json:"is_available,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Domain        string   `json:"domain,omitempty"`
}

type FetchUserParams struct {
	UserNumber      int64
	UserUUID        string
	Status          string
	Name            string
	PhoneNumber     string
	RegisteredEmail string
	WorkEmail       string
	CompanyID       int64
	CompanyName     string
	MentorConfig    MentorConfig
	Designation     string
	CreatedAt       *time.Time
}

func (u *User) SetGoogleOAuthDetails(data GoogleOAuthDetails) {
	u.GoogleAuthJSON.SetData(data)
}

func (u *User) GoogleOAuthDetails() GoogleOAuthDetails {
	return u.GoogleAuthJSON.GetData()
}

func (u *User) SetTechnicalSkills(data TechnicalSkills) {
	u.TechnicalSkillsJSON.SetData(data)
}

func (u *User) TechnicalSkills() TechnicalSkills {
	return u.TechnicalSkillsJSON.GetData()
}

func (u *User) SetMentorConfig(data MentorConfig) {
	u.MentorConfigJSON.SetData(data)
}

func (u *User) MentorConfig() MentorConfig {
	return u.MentorConfigJSON.GetData()
}

func (u User) IsApproved() bool {
	return u.Status == constants.StatusApproved || u.Status == constants.StatusAutoApproved
}
