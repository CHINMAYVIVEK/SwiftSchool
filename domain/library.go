package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: library.categories
type BookCategory struct {
	TenantUUIDModel
	Name *string `json:"name,omitempty" db:"name"`
}

// Corresponds to schema: library.books
type Book struct {
	TenantUUIDModel
	Title           string     `json:"title" db:"title"`
	Author          *string    `json:"author,omitempty" db:"author"`
	ISBN            *string    `json:"isbn,omitempty" db:"isbn"`
	CategoryID      *uuid.UUID `json:"category_id,omitempty" db:"category_id"`
	TotalCopies     int        `json:"total_copies" db:"total_copies"`
	AvailableCopies int        `json:"available_copies" db:"available_copies"`
}

// Corresponds to schema: library.book_issues
type BookIssue struct {
	TenantUUIDModel
	BookID     *uuid.UUID `json:"book_id,omitempty" db:"book_id"`
	MemberID   uuid.UUID  `json:"member_id" db:"member_id"`
	MemberType MemberType `json:"member_type" db:"member_type"`
	IssueDate  *time.Time `json:"issue_date,omitempty" db:"issue_date"`
	DueDate    time.Time  `json:"due_date" db:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty" db:"return_date"`
	FineAmount float64    `json:"fine_amount" db:"fine_amount"`
	Status     string     `json:"status" db:"status"` // issued, returned
}

// Corresponds to schema: library.circulation_rules
type CirculationRule struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	InstituteID    uuid.UUID  `json:"institute_id" db:"institute_id"`
	MemberType     MemberType `json:"member_type" db:"member_type"`
	MaxBooksIssued *int       `json:"max_books_issued,omitempty" db:"max_books_issued"`
	IssueDaysLimit *int       `json:"issue_days_limit,omitempty" db:"issue_days_limit"`
	FinePerDay     *float64   `json:"fine_per_day,omitempty" db:"fine_per_day"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}
