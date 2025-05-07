package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("string length must be between %d and %d characters", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 1000); err != nil {
		return fmt.Errorf("invalid username: %w", err)
	}

	if !isValidUsername(value) {
		return fmt.Errorf("invalid character in username")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	return nil
}

func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 1000); err != nil {
		return fmt.Errorf("invalid fullname: %w", err)
	}

	if !isValidFullname(value) {
		return fmt.Errorf("invalid character in fullname")
	}

	return nil
}
