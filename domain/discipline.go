package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: discipline.incidents
type Incident struct {
	TenantUUIDModel
	StudentID    *uuid.UUID       `json:"student_id,omitempty" db:"student_id"`
	IncidentDate *time.Time       `json:"incident_date,omitempty" db:"incident_date"`
	Title        *string          `json:"title,omitempty" db:"title"`
	Description  *string          `json:"description,omitempty" db:"description"`
	Severity     IncidentSeverity `json:"severity,omitempty" db:"severity"`
	ActionTaken  *string          `json:"action_taken,omitempty" db:"action_taken"`
	ReportedBy   *uuid.UUID       `json:"reported_by,omitempty" db:"reported_by"`
}

// Corresponds to schema: discipline.actions
type DisciplineAction struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	IncidentID  *uuid.UUID `json:"incident_id,omitempty" db:"incident_id"`
	ActionType  *string    `json:"action_type,omitempty" db:"action_type"`
	StartDate   *time.Time `json:"start_date,omitempty" db:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty" db:"end_date"`
	Remarks     *string    `json:"remarks,omitempty" db:"remarks"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
