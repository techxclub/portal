package utils

import (
	"github.com/nyaruka/phonenumbers"
)

func IsValidPhoneNumber(phoneNumber string) bool {
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return false
	}

	return phonenumbers.IsValidNumber(parsedNumber)
}

func SanitizePhoneNumber(phoneNumber string) string {
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return ""
	}

	return phonenumbers.Format(parsedNumber, phonenumbers.E164)
}
