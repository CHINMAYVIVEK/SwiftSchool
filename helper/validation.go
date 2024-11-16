package helper

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"unicode"

	"github.com/google/uuid"
)

const (
	passwordLength = 12
	passwordChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]~"
)

var (
	emailPattern  = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	mobilePattern = regexp.MustCompile(`^\+\d{1,3}\d{10}$`)
	nanoIDPattern = regexp.MustCompile(`^[a-z0-9_-]{21,}$`)
)

// PatternValidation validates a string value against a predefined pattern (Email or Mobile)
func PatternValidation(patternType, value string) bool {
	switch patternType {
	case "Mobile":
		return mobilePattern.MatchString(value)
	case "Email":
		return emailPattern.MatchString(value)
	case "NanoID":
		return nanoIDPattern.MatchString(value)
	default:
		return false
	}
}

// IsValidID checks if the class ID contains only alphanumeric characters
func IsValidID(classID string) bool {
	for _, r := range classID {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsPasswordStrong checks if the password meets strength criteria
func IsPasswordStrong(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpperCase, hasLowerCase, hasDigit, hasSpecialChar bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		case unicode.IsDigit(char):
			hasDigit = true
		default:
			hasSpecialChar = true
		}

		// If all conditions are met, we can return early
		if hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar {
			return true
		}
	}

	return hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar
}

// GenerateRandomPassword generates a random strong password of length 12
func GenerateRandomPassword() string {
	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		password[i] = passwordChars[index.Int64()]
	}
	return string(password)
}

// IsValidUUID checks if the provided string is a valid UUID
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}