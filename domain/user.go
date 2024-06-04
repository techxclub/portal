package domain

type User struct {
	UserIDNum         int64   `db:"user_id_num"` // Only for internal use
	UserID            string  `db:"user_id"`
	FirstName         string  `db:"first_name"`
	LastName          string  `db:"last_name"`
	YearsOfExperience float32 `db:"years_of_experience"`
	PersonalEmail     string  `db:"personal_email"`
	WorkEmail         string  `db:"work_email"`
	PhoneNumber       string  `db:"phone_number"`
	LinkedIn          string  `db:"linkedin"`
	Role              string  `db:"role"`
}

type UserDetailsRequest struct {
	UserID string
	Phone  string
}
