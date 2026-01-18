package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== AUTH ====================

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Identifier string `json:"identifier" example:"john.doe@example.com"`
	UserType   string `json:"user_type" example:"admin"`
}

// LoginResponse represents the response for login operation
type LoginResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"OTP sent successfully"`
}

// VerifyOTPRequest represents the request body for OTP verification
type VerifyOTPRequest struct {
	Identifier string `json:"identifier" example:"john.doe@example.com"`
	UserType   string `json:"user_type" example:"admin"`
	OTP        string `json:"otp" example:"123456"`
}

// VerifyOTPResponse represents the response for OTP verification
type VerifyOTPResponse struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"Login successful"`
	User    UserResponse `json:"user"`
	Token   string       `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ==================== USER ====================

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username       string     `json:"username" example:"john.doe"`
	Password       string     `json:"password" example:"SecurePassword123!"`
	RoleType       string     `json:"role_type" example:"admin"`
	LinkedEntityID uuid.UUID  `json:"linked_entity_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID    *uuid.UUID `json:"institute_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	IsActive       bool       `json:"is_active" example:"true"`
}

// UserResponse represents the response for user operations
type UserResponse struct {
	ID             uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username       string     `json:"username" example:"john.doe"`
	RoleType       string     `json:"role_type" example:"admin"`
	LinkedEntityID uuid.UUID  `json:"linked_entity_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID    *uuid.UUID `json:"institute_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	IsActive       bool       `json:"is_active" example:"true"`
	CreatedAt      time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt      time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// UpdateUserPasswordRequest represents the request body for updating user password
type UpdateUserPasswordRequest struct {
	ID       uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Password string    `json:"password" example:"NewSecurePassword123!"`
}

// UpdateUserStatusRequest represents the request body for updating user status
type UpdateUserStatusRequest struct {
	ID       uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	IsActive bool      `json:"is_active" example:"false"`
}
