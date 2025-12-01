package domain

import "errors"

// Standard Sentinel Errors
// Use these to check error types using errors.Is(err, domain.ErrNotFound)
var (
	// ErrNotFound indicates a specific resource could not be found
	ErrNotFound = errors.New("resource not found")

	// ErrConflict indicates a violation of unique constraints (e.g. duplicate email)
	ErrConflict = errors.New("resource already exists")

	// ErrInternal indicates an unexpected system failure
	ErrInternal = errors.New("internal system error")

	// ErrUnauthorized indicates missing or invalid credentials
	ErrUnauthorized = errors.New("unauthorized action")

	// ErrForbidden indicates valid credentials but insufficient permissions
	ErrForbidden = errors.New("access forbidden")

	// ErrValidation indicates invalid input data
	ErrValidation = errors.New("validation failed")
)

// AppError is a custom error struct if you need to pass specific codes or context
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"` // Internal error, do not expose to JSON
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Helper constructors
func NewValidationError(msg string) error {
	return &AppError{Code: 400, Message: msg, Err: ErrValidation}
}

func NewNotFoundError(msg string) error {
	return &AppError{Code: 404, Message: msg, Err: ErrNotFound}
}
