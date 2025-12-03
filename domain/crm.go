package domain

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	TenantUUIDModel
	StudentID       uuid.UUID  `json:"student_id" db:"student_id"`
	BloodGroup      *string    `json:"blood_group,omitempty" db:"blood_group"`
	Allergies       *string    `json:"allergies,omitempty" db:"allergies"`
	Medications     *string    `json:"medications,omitempty" db:"medications"`
	EmergencyNotes  *string    `json:"emergency_notes,omitempty" db:"emergency_notes"`
	LastCheckupDate *time.Time `json:"last_checkup_date,omitempty" db:"last_checkup_date"`
}

type SupportTicket struct {
	TenantUUIDModel
	RaisedBy    uuid.UUID  `json:"raised_by" db:"raised_by"`
	CategoryID  uuid.UUID  `json:"category_id" db:"category_id"`
	Priority    string     `json:"priority" db:"priority"` // Enum
	Status      string     `json:"status" db:"status"`     // Enum
	Subject     string     `json:"subject" db:"subject"`
	Description string     `json:"description" db:"description"`
	AssignedTo  *uuid.UUID `json:"assigned_to,omitempty" db:"assigned_to"`
}

type Visitor struct {
	TenantUUIDModel
	VisitorType  string     `json:"visitor_type" db:"visitor_type"`
	Name         string     `json:"name" db:"name"`
	Phone        *string    `json:"phone,omitempty" db:"phone"`
	Purpose      *string    `json:"purpose,omitempty" db:"purpose"`
	PersonToMeet *uuid.UUID `json:"person_to_meet,omitempty" db:"person_to_meet"`
	CheckIn      time.Time  `json:"check_in" db:"check_in"`
	CheckOut     *time.Time `json:"check_out,omitempty" db:"check_out"`
}
