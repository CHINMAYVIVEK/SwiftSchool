package helper

import "fmt"

// Common error types for request parsing
var (
	ErrEmptyRequestBody = fmt.Errorf("request body is empty")
)

// ErrMissingParameter returns an error for a missing required parameter
func ErrMissingParameter(paramName string) error {
	return fmt.Errorf("required parameter '%s' is missing", paramName)
}

// ErrInvalidParameter returns an error for an invalid parameter value
func ErrInvalidParameter(paramName, reason string) error {
	return fmt.Errorf("invalid parameter '%s': %s", paramName, reason)
}
