package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: health.medical_records
type MedicalRecord struct {
	TenantUUIDModel
	StudentID       uuid.UUID  `json:"student_id" db:"student_id"`
	BloodGroup      *string    `json:"blood_group,omitempty" db:"blood_group"`
	Allergies       *string    `json:"allergies,omitempty" db:"allergies"`
	Medications     *string    `json:"medications,omitempty" db:"medications"`
	EmergencyNotes  *string    `json:"emergency_notes,omitempty" db:"emergency_notes"`
	LastCheckupDate *time.Time `json:"last_checkup_date,omitempty" db:"last_checkup_date"`
}

// Corresponds to schema: health.infirmary_visits
type InfirmaryVisit struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	StudentID   uuid.UUID  `json:"student_id" db:"student_id"`
	VisitDate   *time.Time `json:"visit_date,omitempty" db:"visit_date"`
	Symptoms    *string    `json:"symptoms,omitempty" db:"symptoms"`
	Treatment   *string    `json:"treatment_given,omitempty" db:"treatment_given"`
	AttendedBy  *uuid.UUID `json:"attended_by,omitempty" db:"attended_by"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: health.vaccinations
type Vaccination struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	InstituteID      uuid.UUID  `json:"institute_id" db:"institute_id"`
	StudentID        uuid.UUID  `json:"student_id" db:"student_id"`
	VaccineName      *string    `json:"vaccine_name,omitempty" db:"vaccine_name"`
	DateAdministered *time.Time `json:"date_administered,omitempty" db:"date_administered"`
	NextDueDate      *time.Time `json:"next_due_date,omitempty" db:"next_due_date"`
	Remarks          *string    `json:"remarks,omitempty" db:"remarks"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
}
