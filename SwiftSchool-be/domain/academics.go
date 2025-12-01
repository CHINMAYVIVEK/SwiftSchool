package domain

import "github.com/google/uuid"

type Subject struct {
	TenantUUIDModel
	Name string  `json:"name" db:"name"`
	Code *string `json:"code,omitempty" db:"code"`
}

type ClassPeriod struct {
	TenantUUIDModel
	Name      *string `json:"name,omitempty" db:"name"`
	StartTime string  `json:"start_time" db:"start_time"`
	EndTime   string  `json:"end_time" db:"end_time"`
	IsBreak   bool    `json:"is_break" db:"is_break"`
}

type TimetableEntry struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID  `json:"academic_session_id" db:"academic_session_id"`
	ClassID           *uuid.UUID `json:"class_id,omitempty" db:"class_id"`
	DayOfWeek         DayOfWeek  `json:"day_of_week" db:"day_of_week"`
	PeriodID          *uuid.UUID `json:"period_id,omitempty" db:"period_id"`
	SubjectID         *uuid.UUID `json:"subject_id,omitempty" db:"subject_id"`
	TeacherID         *uuid.UUID `json:"teacher_id,omitempty" db:"teacher_id"`
}
