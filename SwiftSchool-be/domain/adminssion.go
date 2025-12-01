package domain

import (
	"time"
)

type AdmissionEnquiry struct {
	TenantUUIDModel
	StudentName      *string         `json:"student_name,omitempty" db:"student_name"`
	GuardianName     *string         `json:"guardian_name,omitempty" db:"guardian_name"`
	Phone            *string         `json:"phone,omitempty" db:"phone"`
	Email            *string         `json:"email,omitempty" db:"email"`
	ClassApplyingFor *string         `json:"class_applying_for,omitempty" db:"class_applying_for"`
	PreviousSchool   *string         `json:"previous_school,omitempty" db:"previous_school"`
	Status           AdmissionStatus `json:"status" db:"status"`
	EnquiryDate      *time.Time      `json:"enquiry_date,omitempty" db:"enquiry_date"`
	FollowUpDate     *time.Time      `json:"follow_up_date,omitempty" db:"follow_up_date"`
}
