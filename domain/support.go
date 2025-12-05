package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: support.ticket_categories
type TicketCategory struct {
	ID          uuid.UUID `json:"id" db:"id"`
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
	Name        string    `json:"name" db:"name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Corresponds to schema: support.tickets
type SupportTicket struct {
	TenantUUIDModel
	RaisedBy    uuid.UUID      `json:"raised_by" db:"raised_by"`
	CategoryID  uuid.UUID      `json:"category_id" db:"category_id"`
	Priority    TicketPriority `json:"priority" db:"priority"`
	Status      TicketStatus   `json:"status" db:"status"`
	Subject     string         `json:"subject" db:"subject"`
	Description string         `json:"description" db:"description"`
	AssignedTo  *uuid.UUID     `json:"assigned_to,omitempty" db:"assigned_to"`
}

// Corresponds to schema: support.ticket_comments
type TicketComment struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	TicketID  uuid.UUID  `json:"ticket_id" db:"ticket_id"`
	UserID    *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	Comment   *string    `json:"comment,omitempty" db:"comment"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}
