package isEmail

import (
	"net/mail"
)

// IsEmail checks if the given string is a valid email address.
func IsEmail(email string) bool {
	// Regular expression for validating an email
	_, err := mail.ParseAddress(email)
	if err != nil {

		return false
	}

	return true
}
