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
