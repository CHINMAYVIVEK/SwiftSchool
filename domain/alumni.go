package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: alumni.profiles
type AlumniProfile struct {
	TenantUUIDModel
	StudentID           *uuid.UUID `json:"student_id,omitempty" db:"student_id"`
	GraduationYear      *int       `json:"graduation_year,omitempty" db:"graduation_year"`
	CurrentOrganization *string    `json:"current_organization,omitempty" db:"current_organization"`
	Designation         *string    `json:"designation,omitempty" db:"designation"`
	LinkedInURL         *string    `json:"linkedin_url,omitempty" db:"linkedin_url"`
	IsActiveMember      bool       `json:"is_active_member" db:"is_active_member"`
}

// Corresponds to schema: alumni.donations
type Donation struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	InstituteID    uuid.UUID  `json:"institute_id" db:"institute_id"`
	AlumniID       *uuid.UUID `json:"alumni_id,omitempty" db:"alumni_id"`
	Amount         *float64   `json:"amount,omitempty" db:"amount"`
	DonationDate   *time.Time `json:"donation_date,omitempty" db:"donation_date"`
	Purpose        *string    `json:"purpose,omitempty" db:"purpose"`
	TransactionRef *string    `json:"transaction_ref,omitempty" db:"transaction_ref"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}
