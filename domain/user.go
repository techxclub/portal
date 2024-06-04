package domain

type User struct {
	UserID            int64   `db:"user_id"`
	FirstName         string  `db:"first_name"`
	LastName          string  `db:"last_name"`
	YearsOfExperience float64 `db:"years_of_experience"`
	PersonalEmail     string  `db:"personal_email"`
	WorkEmail         string  `db:"work_email"`
	PhoneNumber       string  `db:"phone_number"`
	LinkedIn          string  `db:"linkedIn"`
	Role              string  `db:"role"`
}
