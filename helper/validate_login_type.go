package helper

import (
	"regexp"
)

type LoginType string

const (
	LoginEmail   LoginType = "email"
	LoginPhone   LoginType = "phone"
	LoginInvalid LoginType = "invalid"
)

// ValidateLoginType checks whether username is a valid email or phone
func ValidateLoginType(username string) LoginType {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)

	switch {
	case emailRegex.MatchString(username):
		return LoginEmail
	case phoneRegex.MatchString(username):
		return LoginPhone
	default:
		return LoginInvalid
	}
}
