package domain

import "github.com/google/uuid"

// Corresponds to schema: auth.users
type User struct {
	BaseUUIDModel
	Username       string     `json:"username" db:"username"`
	PasswordHash   string     `json:"-" db:"password_hash"` // Never expose
	RoleType       UserRole   `json:"role_type" db:"role_type"`
	LinkedEntityID uuid.UUID  `json:"linked_entity_id" db:"linked_entity_id"`   // Links to Student/Employee
	InstituteID    *uuid.UUID `json:"institute_id,omitempty" db:"institute_id"` // Nullable for SuperAdmin
	IsActive       bool       `json:"is_active" db:"is_active"`
}
