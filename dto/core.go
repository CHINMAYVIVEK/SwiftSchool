package dto

import (
	"time"

	"github.com/google/uuid"
)

// ==================== INSTITUTE ====================

// CreateInstituteRequest represents the request body for creating an institute
type CreateInstituteRequest struct {
	Name               string  `json:"name" example:"Green Valley International School"`
	Code               string  `json:"code" example:"GVIS-BLR"`
	CurrencyCode       *string `json:"currency_code,omitempty" example:"INR"`
	LogoURL            *string `json:"logo_url,omitempty" example:"https://example.com/logo.png"`
	Website            *string `json:"website,omitempty" example:"https://greenvalley.edu"`
	Timezone           string  `json:"timezone" example:"Asia/Kolkata"`
	FiscalYearStartMon int     `json:"fiscal_year_start_month" example:"4"`
	IsActive           bool    `json:"is_active" example:"true"`
}

// UpdateInstituteRequest represents the request body for updating an institute
type UpdateInstituteRequest struct {
	ID                 uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name               string    `json:"name" example:"Green Valley International School"`
	Code               string    `json:"code" example:"GVIS-BLR"`
	CurrencyCode       *string   `json:"currency_code,omitempty" example:"INR"`
	LogoURL            *string   `json:"logo_url,omitempty" example:"https://example.com/logo.png"`
	Website            *string   `json:"website,omitempty" example:"https://greenvalley.edu"`
	Timezone           string    `json:"timezone" example:"Asia/Kolkata"`
	FiscalYearStartMon int       `json:"fiscal_year_start_month" example:"4"`
	IsActive           bool      `json:"is_active" example:"true"`
}

// InstituteResponse represents the response for institute operations
type InstituteResponse struct {
	ID                 uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name               string     `json:"name" example:"Green Valley International School"`
	Code               string     `json:"code" example:"GVIS-BLR"`
	CurrencyCode       *string    `json:"currency_code,omitempty" example:"INR"`
	LogoURL            *string    `json:"logo_url,omitempty" example:"https://example.com/logo.png"`
	Website            *string    `json:"website,omitempty" example:"https://greenvalley.edu"`
	Timezone           string     `json:"timezone" example:"Asia/Kolkata"`
	FiscalYearStartMon int        `json:"fiscal_year_start_month" example:"4"`
	IsActive           bool       `json:"is_active" example:"true"`
	CreatedAt          time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt          time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

// ==================== CLASS ====================

// CreateClassRequest represents the request body for creating a class
type CreateClassRequest struct {
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AcademicSessionID uuid.UUID  `json:"academic_session_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Name              string     `json:"name" example:"Grade 10"`
	Section           string     `json:"section" example:"A"`
	ClassTeacherID    *uuid.UUID `json:"class_teacher_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
}

// UpdateClassRequest represents the request body for updating a class
type UpdateClassRequest struct {
	ID                uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AcademicSessionID uuid.UUID  `json:"academic_session_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Name              string     `json:"name" example:"Grade 10"`
	Section           string     `json:"section" example:"A"`
	ClassTeacherID    *uuid.UUID `json:"class_teacher_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
}

// ClassResponse represents the response for class operations
type ClassResponse struct {
	ID                uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AcademicSessionID uuid.UUID  `json:"academic_session_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Name              string     `json:"name" example:"Grade 10"`
	Section           string     `json:"section" example:"A"`
	ClassTeacherID    *uuid.UUID `json:"class_teacher_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	CreatedAt         time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== DEPARTMENT ====================

// CreateDepartmentRequest represents the request body for creating a department
type CreateDepartmentRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"Science Department"`
}

// UpdateDepartmentRequest represents the request body for updating a department
type UpdateDepartmentRequest struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"Science Department"`
}

// DepartmentResponse represents the response for department operations
type DepartmentResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"Science Department"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== STUDENT ====================

// CreateStudentRequest represents the request body for creating a student
type CreateStudentRequest struct {
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AdmissionNo       string     `json:"admission_no" example:"STU2024001"`
	FirstName         string     `json:"first_name" example:"John"`
	LastName          *string    `json:"last_name,omitempty" example:"Doe"`
	DOB               *time.Time `json:"dob,omitempty" example:"2010-05-15T00:00:00Z"`
	Gender            string     `json:"gender,omitempty" example:"male"`
	BloodGroup        string     `json:"blood_group,omitempty" example:"O+"`
	SocialCategory    string     `json:"social_category,omitempty" example:"general"`
	CurrentClassID    *uuid.UUID `json:"current_class_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	Nationality       *string    `json:"nationality,omitempty" example:"Indian"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"English"`
}

// UpdateStudentRequest represents the request body for updating a student
type UpdateStudentRequest struct {
	ID                uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AdmissionNo       string     `json:"admission_no" example:"STU2024001"`
	FirstName         string     `json:"first_name" example:"John"`
	LastName          *string    `json:"last_name,omitempty" example:"Doe"`
	DOB               *time.Time `json:"dob,omitempty" example:"2010-05-15T00:00:00Z"`
	Gender            string     `json:"gender,omitempty" example:"male"`
	BloodGroup        string     `json:"blood_group,omitempty" example:"O+"`
	SocialCategory    string     `json:"social_category,omitempty" example:"general"`
	CurrentClassID    *uuid.UUID `json:"current_class_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	Nationality       *string    `json:"nationality,omitempty" example:"Indian"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"English"`
}

// StudentResponse represents the response for student operations
type StudentResponse struct {
	ID                uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID       uuid.UUID  `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AdmissionNo       string     `json:"admission_no" example:"STU2024001"`
	FirstName         string     `json:"first_name" example:"John"`
	LastName          *string    `json:"last_name,omitempty" example:"Doe"`
	DOB               *time.Time `json:"dob,omitempty" example:"2010-05-15T00:00:00Z"`
	Gender            string     `json:"gender,omitempty" example:"male"`
	BloodGroup        string     `json:"blood_group,omitempty" example:"O+"`
	SocialCategory    string     `json:"social_category,omitempty" example:"general"`
	CurrentClassID    *uuid.UUID `json:"current_class_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	Nationality       *string    `json:"nationality,omitempty" example:"Indian"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"English"`
	CreatedAt         time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// StudentFullProfileResponse represents the full profile response for a student
type StudentFullProfileResponse struct {
	Student   StudentResponse    `json:"student"`
	Guardians []GuardianResponse `json:"guardians,omitempty"`
	Addresses []AddressResponse  `json:"addresses,omitempty"`
	Class     *ClassResponse     `json:"class,omitempty"`
}

// SearchStudentsRequest represents the request for searching students
type SearchStudentsRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Query       string    `json:"query" example:"John"`
}

// ==================== GUARDIAN ====================

// CreateGuardianRequest represents the request body for creating a guardian
type CreateGuardianRequest struct {
	FirstName    string   `json:"first_name" example:"Robert"`
	LastName     *string  `json:"last_name,omitempty" example:"Doe"`
	Email        *string  `json:"email,omitempty" example:"robert.doe@example.com"`
	Phone        *string  `json:"phone,omitempty" example:"+91-9876543210"`
	Profession   *string  `json:"profession,omitempty" example:"Engineer"`
	AnnualIncome *float64 `json:"annual_income,omitempty" example:"1200000"`
}

// GuardianResponse represents the response for guardian operations
type GuardianResponse struct {
	ID           uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName    string    `json:"first_name" example:"Robert"`
	LastName     *string   `json:"last_name,omitempty" example:"Doe"`
	Email        *string   `json:"email,omitempty" example:"robert.doe@example.com"`
	Phone        *string   `json:"phone,omitempty" example:"+91-9876543210"`
	Profession   *string   `json:"profession,omitempty" example:"Engineer"`
	AnnualIncome *float64  `json:"annual_income,omitempty" example:"1200000"`
	CreatedAt    time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// LinkStudentGuardianRequest represents the request for linking a student to a guardian
type LinkStudentGuardianRequest struct {
	StudentID        uuid.UUID `json:"student_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	GuardianID       uuid.UUID `json:"guardian_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Relationship     string    `json:"relationship" example:"father"`
	IsPrimaryContact bool      `json:"is_primary_contact" example:"true"`
}

// ==================== ACADEMIC SESSION ====================

// CreateAcademicSessionRequest represents the request body for creating an academic session
type CreateAcademicSessionRequest struct {
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"2024-2025"`
	StartDate   time.Time `json:"start_date" example:"2024-04-01T00:00:00Z"`
	EndDate     time.Time `json:"end_date" example:"2025-03-31T23:59:59Z"`
	IsActive    bool      `json:"is_active" example:"true"`
}

// UpdateAcademicSessionRequest represents the request body for updating an academic session
type UpdateAcademicSessionRequest struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"2024-2025"`
	StartDate   time.Time `json:"start_date" example:"2024-04-01T00:00:00Z"`
	EndDate     time.Time `json:"end_date" example:"2025-03-31T23:59:59Z"`
	IsActive    bool      `json:"is_active" example:"true"`
}

// AcademicSessionResponse represents the response for academic session operations
type AcademicSessionResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	InstituteID uuid.UUID `json:"institute_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string    `json:"name" example:"2024-2025"`
	StartDate   time.Time `json:"start_date" example:"2024-04-01T00:00:00Z"`
	EndDate     time.Time `json:"end_date" example:"2025-03-31T23:59:59Z"`
	IsActive    bool      `json:"is_active" example:"true"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ==================== ADDRESS ====================

// CreateAddressRequest represents the request body for creating an address
type CreateAddressRequest struct {
	OwnerID      uuid.UUID  `json:"owner_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OwnerType    string     `json:"owner_type" example:"student"`
	AddressType  string     `json:"address_type" example:"current"`
	AddressLine1 string     `json:"address_line_1" example:"123 Main Street"`
	AddressLine2 *string    `json:"address_line_2,omitempty" example:"Apartment 4B"`
	CountryID    *uuid.UUID `json:"country_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	StateID      *uuid.UUID `json:"state_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	DistrictID   *uuid.UUID `json:"district_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440003"`
	PostalCode   *string    `json:"postal_code,omitempty" example:"560001"`
}

// AddressResponse represents the response for address operations
type AddressResponse struct {
	ID           uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OwnerID      uuid.UUID  `json:"owner_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OwnerType    string     `json:"owner_type" example:"student"`
	AddressType  string     `json:"address_type" example:"current"`
	AddressLine1 string     `json:"address_line_1" example:"123 Main Street"`
	AddressLine2 *string    `json:"address_line_2,omitempty" example:"Apartment 4B"`
	CountryID    *uuid.UUID `json:"country_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	StateID      *uuid.UUID `json:"state_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	DistrictID   *uuid.UUID `json:"district_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440003"`
	PostalCode   *string    `json:"postal_code,omitempty" example:"560001"`
	CreatedAt    time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt    time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}
