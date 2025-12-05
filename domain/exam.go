package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: exam.grade_systems
type GradeSystem struct {
	TenantUUIDModel
	Name string  `json:"name" db:"name"`
	Type *string `json:"type,omitempty" db:"type"` // percentage, gpa
}

// Corresponds to schema: exam.grade_rules
type GradeRule struct {
	ID            uuid.UUID `json:"id" db:"id"`
	InstituteID   uuid.UUID `json:"institute_id" db:"institute_id"`
	GradeSystemID uuid.UUID `json:"grade_system_id" db:"grade_system_id"`
	GradeName     string    `json:"grade_name" db:"grade_name"`
	MinPercentage *float64  `json:"min_percentage,omitempty" db:"min_percentage"`
	MaxPercentage *float64  `json:"max_percentage,omitempty" db:"max_percentage"`
	GradePoint    *float64  `json:"grade_point,omitempty" db:"grade_point"`
	Description   *string   `json:"description,omitempty" db:"description"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// Corresponds to schema: exam.exams
type Exam struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID  `json:"academic_session_id" db:"academic_session_id"`
	Name              string     `json:"name" db:"name"`
	StartDate         *time.Time `json:"start_date,omitempty" db:"start_date"`
	EndDate           *time.Time `json:"end_date,omitempty" db:"end_date"`
	IsPublished       bool       `json:"is_published" db:"is_published"`
}

// Corresponds to schema: exam.schedules
type ExamSchedule struct {
	TenantUUIDModel
	ExamID          uuid.UUID `json:"exam_id" db:"exam_id"`
	ClassID         uuid.UUID `json:"class_id" db:"class_id"`
	SubjectID       uuid.UUID `json:"subject_id" db:"subject_id"`
	ExamDate        time.Time `json:"exam_date" db:"exam_date"`
	DurationMinutes *int      `json:"duration_minutes,omitempty" db:"duration_minutes"`
	MaxMarks        float64   `json:"max_marks" db:"max_marks"`
	MinPassMarks    float64   `json:"min_pass_marks" db:"min_pass_marks"`
	ExamType        string    `json:"exam_type" db:"exam_type"` // theory, practical
}

// Corresponds to schema: exam.marks
type ExamMark struct {
	TenantUUIDModel
	ScheduleID    uuid.UUID `json:"schedule_id" db:"schedule_id"`
	StudentID     uuid.UUID `json:"student_id" db:"student_id"`
	MarksObtained *float64  `json:"marks_obtained,omitempty" db:"marks_obtained"`
	IsAbsent      bool      `json:"is_absent" db:"is_absent"`
	Remarks       *string   `json:"remarks,omitempty" db:"remarks"`
}
