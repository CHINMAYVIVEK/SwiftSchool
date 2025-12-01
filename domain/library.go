package domain

import (
	"time"

	"github.com/google/uuid"
)

type BookCategory struct {
	TenantUUIDModel
	Name *string `json:"name,omitempty" db:"name"`
}

type Book struct {
	TenantUUIDModel
	Title           string     `json:"title" db:"title"`
	Author          *string    `json:"author,omitempty" db:"author"`
	ISBN            *string    `json:"isbn,omitempty" db:"isbn"`
	CategoryID      *uuid.UUID `json:"category_id,omitempty" db:"category_id"`
	TotalCopies     int        `json:"total_copies" db:"total_copies"`
	AvailableCopies int        `json:"available_copies" db:"available_copies"`
}

type BookIssue struct {
	TenantUUIDModel
	BookID     *uuid.UUID `json:"book_id,omitempty" db:"book_id"`
	MemberID   uuid.UUID  `json:"member_id" db:"member_id"`
	MemberType MemberType `json:"member_type" db:"member_type"`
	IssueDate  *time.Time `json:"issue_date,omitempty" db:"issue_date"`
	DueDate    time.Time  `json:"due_date" db:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty" db:"return_date"`
	FineAmount float64    `json:"fine_amount" db:"fine_amount"`
	Status     string     `json:"status" db:"status"`
}
