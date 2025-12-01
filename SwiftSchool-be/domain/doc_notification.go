package domain

import (
	"github.com/google/uuid"
)

// Document represents docs.documents
type Document struct {
	TenantUUIDModel
	OwnerID   uuid.UUID    `json:"owner_id" db:"owner_id"`
	OwnerType OwnerType    `json:"owner_type" db:"owner_type"`
	DocType   DocumentType `json:"doc_type" db:"doc_type"`
	FileName  *string      `json:"file_name,omitempty" db:"file_name"`
	FileURL   string       `json:"file_url" db:"file_url"`
}

// Notification represents comms.notifications
type Notification struct {
	TenantUUIDModel
	UserID  *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	Title   *string    `json:"title,omitempty" db:"title"`
	Message *string    `json:"message,omitempty" db:"message"`
	IsRead  bool       `json:"is_read" db:"is_read"`
}
