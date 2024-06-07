package request

import (
	"github.com/nyaruka/phonenumbers"
	"github.com/techx/portal/errors"
)

// IsValidPhoneNumber validates phone number using libphonenumber
func IsValidPhoneNumber(phoneNumber string) error {
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return errors.ErrInvalidPhoneNumberFormat
	}
	if !phonenumbers.IsValidNumber(parsedNumber) {
		return errors.ErrInvalidPhoneNumber
	}
	return nil
}
