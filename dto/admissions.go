package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== ENQUIRY ====================

// CreateEnquiryRequest represents the request body for creating an admission enquiry
type CreateEnquiryRequest struct {
	InstituteID      uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StudentName      *string    `json:"student_name,omitempty" example:"Jane Smith"`
	GuardianName     *string    `json:"guardian_name,omitempty" example:"Michael Smith"`
	Phone            *string    `json:"phone,omitempty" example:"+91-9876543210"`
	Email            *string    `json:"email,omitempty" example:"michael.smith@example.com"`
	ClassApplyingFor *string    `json:"class_applying_for,omitempty" example:"Grade 5"`
	PreviousSchool   *string    `json:"previous_school,omitempty" example:"ABC Public School"`
	Status           string     `json:"status" example:"open"`
	EnquiryDate      *time.Time `json:"enquiry_date,omitempty" example:"2024-01-15T00:00:00Z"`
	FollowUpDate     *time.Time `json:"follow_up_date,omitempty" example:"2024-01-20T00:00:00Z"`
}

// EnquiryResponse represents the response for enquiry operations
type EnquiryResponse struct {
	ID               uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID      uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StudentName      *string    `json:"student_name,omitempty" example:"Jane Smith"`
	GuardianName     *string    `json:"guardian_name,omitempty" example:"Michael Smith"`
	Phone            *string    `json:"phone,omitempty" example:"+91-9876543210"`
	Email            *string    `json:"email,omitempty" example:"michael.smith@example.com"`
	ClassApplyingFor *string    `json:"class_applying_for,omitempty" example:"Grade 5"`
	PreviousSchool   *string    `json:"previous_school,omitempty" example:"ABC Public School"`
	Status           string     `json:"status" example:"open"`
	EnquiryDate      *time.Time `json:"enquiry_date,omitempty" example:"2024-01-15T00:00:00Z"`
	FollowUpDate     *time.Time `json:"follow_up_date,omitempty" example:"2024-01-20T00:00:00Z"`
	CreatedAt        time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt        time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// UpdateEnquiryStatusRequest represents the request body for updating enquiry status
type UpdateEnquiryStatusRequest struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Status      string    `json:"status" example:"contacted"`
}
