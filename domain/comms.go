package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: docs.documents
// Placed here for cohesion with communications if a separate docs.go is not used.
type Document struct {
	TenantUUIDModel
	OwnerID   uuid.UUID    `json:"owner_id" db:"owner_id"`
	OwnerType OwnerType    `json:"owner_type" db:"owner_type"`
	DocType   DocumentType `json:"doc_type" db:"doc_type"`
	FileName  *string      `json:"file_name,omitempty" db:"file_name"`
	FileURL   string       `json:"file_url" db:"file_url"`
}

// Corresponds to schema: comms.notifications
type Notification struct {
	TenantUUIDModel
	UserID  *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	Title   *string    `json:"title,omitempty" db:"title"`
	Message *string    `json:"message,omitempty" db:"message"`
	IsRead  bool       `json:"is_read" db:"is_read"`
}

// Corresponds to schema: comms.sms_logs
type SMSLog struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	InstituteID    uuid.UUID  `json:"institute_id" db:"institute_id"`
	RecipientPhone *string    `json:"recipient_phone,omitempty" db:"recipient_phone"`
	MessageBody    *string    `json:"message_body,omitempty" db:"message_body"`
	Status         *string    `json:"status,omitempty" db:"status"`
	ProviderID     *string    `json:"provider_id,omitempty" db:"provider_id"`
	SentAt         *time.Time `json:"sent_at,omitempty" db:"sent_at"`
}

// Corresponds to schema: comms.email_logs
type EmailLog struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	InstituteID    uuid.UUID  `json:"institute_id" db:"institute_id"`
	RecipientEmail *string    `json:"recipient_email,omitempty" db:"recipient_email"`
	Subject        *string    `json:"subject,omitempty" db:"subject"`
	Status         *string    `json:"status,omitempty" db:"status"`
	SentAt         *time.Time `json:"sent_at,omitempty" db:"sent_at"`
}
