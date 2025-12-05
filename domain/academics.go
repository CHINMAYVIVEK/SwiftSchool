package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: academics.subjects
type Subject struct {
	TenantUUIDModel
	Name    string  `json:"name" db:"name"`
	Code    *string `json:"code,omitempty" db:"code"`
	Type    string  `json:"type" db:"type"` // mandatory, elective
	Credits float64 `json:"credits" db:"credits"`
}

// Corresponds to schema: academics.student_subjects
type StudentSubject struct {
	ID                uuid.UUID `json:"id" db:"id"`
	InstituteID       uuid.UUID `json:"institute_id" db:"institute_id"`
	AcademicSessionID uuid.UUID `json:"academic_session_id" db:"academic_session_id"`
	StudentID         uuid.UUID `json:"student_id" db:"student_id"`
	SubjectID         uuid.UUID `json:"subject_id" db:"subject_id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}

// Corresponds to schema: academics.class_periods
type ClassPeriod struct {
	TenantUUIDModel
	Name      *string `json:"name,omitempty" db:"name"`
	StartTime string  `json:"start_time" db:"start_time"`
	EndTime   string  `json:"end_time" db:"end_time"`
	IsBreak   bool    `json:"is_break" db:"is_break"`
}

// Corresponds to schema: academics.timetable_entries
type TimetableEntry struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID  `json:"academic_session_id" db:"academic_session_id"`
	ClassID           *uuid.UUID `json:"class_id,omitempty" db:"class_id"`
	DayOfWeek         DayOfWeek  `json:"day_of_week" db:"day_of_week"`
	PeriodID          *uuid.UUID `json:"period_id,omitempty" db:"period_id"`
	SubjectID         *uuid.UUID `json:"subject_id,omitempty" db:"subject_id"`
	TeacherID         *uuid.UUID `json:"teacher_id,omitempty" db:"teacher_id"`
}

// Corresponds to schema: academics.substitutions
type Substitution struct {
	ID                  uuid.UUID  `json:"id" db:"id"`
	InstituteID         uuid.UUID  `json:"institute_id" db:"institute_id"`
	TimetableEntryID    *uuid.UUID `json:"timetable_entry_id,omitempty" db:"timetable_entry_id"`
	OriginalTeacherID   *uuid.UUID `json:"original_teacher_id,omitempty" db:"original_teacher_id"`
	SubstituteTeacherID *uuid.UUID `json:"substitute_teacher_id,omitempty" db:"substitute_teacher_id"`
	Date                time.Time  `json:"date" db:"date"`
	Reason              *string    `json:"reason,omitempty" db:"reason"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: academics.lesson_plans
type LessonPlan struct {
	TenantUUIDModel
	ClassID        uuid.UUID  `json:"class_id" db:"class_id"`
	SubjectID      uuid.UUID  `json:"subject_id" db:"subject_id"`
	TeacherID      uuid.UUID  `json:"teacher_id" db:"teacher_id"`
	Topic          *string    `json:"topic,omitempty" db:"topic"`
	PlannedDate    *time.Time `json:"planned_date,omitempty" db:"planned_date"`
	CompletionDate *time.Time `json:"completion_date,omitempty" db:"completion_date"`
	Status         string     `json:"status" db:"status"` // pending, completed
}

// Corresponds to schema: academics.assignments
type Assignment struct {
	TenantUUIDModel
	ClassID     uuid.UUID  `json:"class_id" db:"class_id"`
	SubjectID   uuid.UUID  `json:"subject_id" db:"subject_id"`
	TeacherID   uuid.UUID  `json:"teacher_id" db:"teacher_id"`
	Title       *string    `json:"title,omitempty" db:"title"`
	Description *string    `json:"description,omitempty" db:"description"`
	DueDate     *time.Time `json:"due_date,omitempty" db:"due_date"`
	MaxMarks    *float64   `json:"max_marks,omitempty" db:"max_marks"`
}

// Corresponds to schema: academics.student_submissions
type StudentSubmission struct {
	TenantUUIDModel
	AssignmentID  uuid.UUID  `json:"assignment_id" db:"assignment_id"`
	StudentID     uuid.UUID  `json:"student_id" db:"student_id"`
	SubmissionURL *string    `json:"submission_url,omitempty" db:"submission_url"`
	SubmittedAt   *time.Time `json:"submitted_at,omitempty" db:"submitted_at"`
	MarksObtained *float64   `json:"marks_obtained,omitempty" db:"marks_obtained"`
	Feedback      *string    `json:"feedback,omitempty" db:"feedback"`
	Status        string     `json:"status" db:"status"` // submitted, graded
}
