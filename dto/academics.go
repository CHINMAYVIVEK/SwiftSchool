package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== SUBJECT ====================

// CreateSubjectRequest represents the request body for creating a subject
type CreateSubjectRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"Mathematics"`
	Code        *string   `json:"code,omitempty" example:"MATH101"`
	Type        string    `json:"type" example:"mandatory"`
	Credits     float64   `json:"credits" example:"4.0"`
}

// SubjectResponse represents the response for subject operations
type SubjectResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"Mathematics"`
	Code        *string   `json:"code,omitempty" example:"MATH101"`
	Type        string    `json:"type" example:"mandatory"`
	Credits     float64   `json:"credits" example:"4.0"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== CLASS PERIOD ====================

// CreateClassPeriodRequest represents the request body for creating a class period
type CreateClassPeriodRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        *string   `json:"name,omitempty" example:"Period 1"`
	StartTime   string    `json:"start_time" example:"09:00:00"`
	EndTime     string    `json:"end_time" example:"09:45:00"`
	IsBreak     bool      `json:"is_break" example:"false"`
}

// ClassPeriodResponse represents the response for class period operations
type ClassPeriodResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        *string   `json:"name,omitempty" example:"Period 1"`
	StartTime   string    `json:"start_time" example:"09:00:00"`
	EndTime     string    `json:"end_time" example:"09:45:00"`
	IsBreak     bool      `json:"is_break" example:"false"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== TIMETABLE ====================

// CreateTimetableEntryRequest represents the request body for creating a timetable entry
type CreateTimetableEntryRequest struct {
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AcademicSessionID uuid.UUID  `json:"academic_session_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	ClassID           *uuid.UUID `json:"class_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	DayOfWeek         string     `json:"day_of_week" example:"mon"`
	PeriodID          *uuid.UUID `json:"period_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440003"`
	SubjectID         *uuid.UUID `json:"subject_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440004"`
	TeacherID         *uuid.UUID `json:"teacher_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440005"`
}

// TimetableEntryResponse represents the response for timetable entry operations
type TimetableEntryResponse struct {
	ID                uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AcademicSessionID uuid.UUID  `json:"academic_session_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	ClassID           *uuid.UUID `json:"class_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	DayOfWeek         string     `json:"day_of_week" example:"mon"`
	PeriodID          *uuid.UUID `json:"period_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440003"`
	SubjectID         *uuid.UUID `json:"subject_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440004"`
	TeacherID         *uuid.UUID `json:"teacher_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440005"`
	CreatedAt         time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}
