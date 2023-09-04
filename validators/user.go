package validators

import (
	"crud-app/models"
	"errors"
	"regexp"
)

func ValidateUser(u *models.User) error {
	if len(u.Username) < 8 || len(u.Password) < 8 {
		return errors.New("Username and Password must be at least 8 characters long")
	}

	// Check if the Password contains at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(u.Password)
	// Check if the Password contains at least one letter
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(u.Password)

	if !hasDigit || !hasLetter {
		return errors.New("Password must contain both numbers and letters")
	}

	return nil
}
