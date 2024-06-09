package twilio

import "time"

type CreateVerificationResponse struct {
	To               string        `json:"to,omitempty"`
	Channel          string        `json:"channel,omitempty"`
	Status           string        `json:"status,omitempty"`
	Lookup           interface{}   `json:"lookup,omitempty"`
	SendCodeAttempts []interface{} `json:"send_code_attempts,omitempty"`
	DateCreated      *time.Time    `json:"date_created,omitempty"`
	DateUpdated      *time.Time    `json:"date_updated,omitempty"`
	URL              string        `json:"url,omitempty"`
}

type CheckVerificationResponse struct {
	To                    string         `json:"to,omitempty"`
	Channel               string         `json:"channel,omitempty"`
	Status                string         `json:"status,omitempty"`
	DateCreated           *time.Time     `json:"date_created,omitempty"`
	DateUpdated           *time.Time     `json:"date_updated,omitempty"`
	SnaAttemptsErrorCodes *[]interface{} `json:"sna_attempts_error_codes,omitempty"`
}
