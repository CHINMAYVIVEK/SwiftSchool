package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== DOCUMENT ====================

// CreateDocumentRequest represents the request body for creating a document
type CreateDocumentRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OwnerID     uuid.UUID `json:"owner_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	OwnerType   string    `json:"owner_type" example:"student"`
	DocType     string    `json:"doc_type" example:"photo"`
	FileName    *string   `json:"file_name,omitempty" example:"student_photo.jpg"`
	FileURL     string    `json:"file_url" example:"https://example.com/documents/student_photo.jpg"`
}

// DocumentResponse represents the response for document operations
type DocumentResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OwnerID     uuid.UUID `json:"owner_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	OwnerType   string    `json:"owner_type" example:"student"`
	DocType     string    `json:"doc_type" example:"photo"`
	FileName    *string   `json:"file_name,omitempty" example:"student_photo.jpg"`
	FileURL     string    `json:"file_url" example:"https://example.com/documents/student_photo.jpg"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== NOTIFICATION ====================

// CreateNotificationRequest represents the request body for creating a notification
type CreateNotificationRequest struct {
	InstituteID uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID      *uuid.UUID `json:"user_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	Title       *string    `json:"title,omitempty" example:"Important Announcement"`
	Message     *string    `json:"message,omitempty" example:"School will remain closed tomorrow due to public holiday"`
	IsRead      bool       `json:"is_read" example:"false"`
}

// NotificationResponse represents the response for notification operations
type NotificationResponse struct {
	ID          uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID      *uuid.UUID `json:"user_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	Title       *string    `json:"title,omitempty" example:"Important Announcement"`
	Message     *string    `json:"message,omitempty" example:"School will remain closed tomorrow due to public holiday"`
	IsRead      bool       `json:"is_read" example:"false"`
	CreatedAt   time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}
