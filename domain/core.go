package domain

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Institute struct {
	BaseUUIDModel
	Name         string  `json:"name" db:"name"`
	Code         string  `json:"code" db:"code"`
	CurrencyCode *string `json:"currency_code" db:"currency_code"`
	LogoURL      *string `json:"logo_url,omitempty" db:"logo_url"`
	Website      *string `json:"website,omitempty" db:"website"`
	IsActive     bool    `json:"is_active" db:"is_active"`
}

type AcademicSession struct {
	TenantUUIDModel
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	IsActive  bool      `json:"is_active" db:"is_active"`
}

type Department struct {
	TenantUUIDModel
	Name string `json:"name" db:"name"`
}

// JSON Structures for DB fields
type LanguageSkill struct {
	Language    string              `json:"language"`
	Proficiency LanguageProficiency `json:"proficiency"`
	CanRead     bool                `json:"can_read"`
	CanWrite    bool                `json:"can_write"`
	CanSpeak    bool                `json:"can_speak"`
}

type SocialMediaHandles map[string]string // e.g., {"twitter": "...", "linkedin": "..."}

type Employee struct {
	TenantUUIDModel
	EmployeeCode       string             `json:"employee_code" db:"employee_code"`
	FirstName          string             `json:"first_name" db:"first_name"`
	LastName           *string            `json:"last_name,omitempty" db:"last_name"`
	DepartmentID       *uuid.UUID         `json:"department_id,omitempty" db:"department_id"`
	Gender             Gender             `json:"gender" db:"gender"`
	MaritalStatus      MaritalStatus      `json:"marital_status,omitempty" db:"marital_status"`
	DateOfJoining      *time.Time         `json:"date_of_joining,omitempty" db:"date_of_joining"`
	Nationality        *string            `json:"nationality,omitempty" db:"nationality"`                   // Added
	PreferredLanguage  *string            `json:"preferred_language,omitempty" db:"preferred_language"`     // Added
	SocialMediaHandles SocialMediaHandles `json:"social_media_handles,omitempty" db:"social_media_handles"` // Added (JSONB)
	LanguageSkills     []LanguageSkill    `json:"language_skills,omitempty" db:"language_skills"`           // Added (JSONB)
	IsActive           bool               `json:"is_active" db:"is_active"`
}

type Class struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID  `json:"academic_session_id" db:"academic_session_id"`
	Name              string     `json:"name" db:"name"`
	Section           string     `json:"section" db:"section"`
	ClassTeacherID    *uuid.UUID `json:"class_teacher_id,omitempty" db:"class_teacher_id"`
}

type Student struct {
	TenantUUIDModel
	AdmissionNo        string             `json:"admission_no" db:"admission_no"`
	FirstName          string             `json:"first_name" db:"first_name"`
	LastName           *string            `json:"last_name,omitempty" db:"last_name"`
	DOB                *time.Time         `json:"dob,omitempty" db:"dob"`
	Gender             Gender             `json:"gender,omitempty" db:"gender"`
	BloodGroup         BloodGroup         `json:"blood_group,omitempty" db:"blood_group"`
	SocialCategory     SocialCategory     `json:"social_category,omitempty" db:"social_category"`
	CurrentClassID     *uuid.UUID         `json:"current_class_id,omitempty" db:"current_class_id"`
	Nationality        *string            `json:"nationality,omitempty" db:"nationality"`                   // Added
	PreferredLanguage  *string            `json:"preferred_language,omitempty" db:"preferred_language"`     // Added
	SocialMediaHandles SocialMediaHandles `json:"social_media_handles,omitempty" db:"social_media_handles"` // Added (JSONB)
	LanguageSkills     []LanguageSkill    `json:"language_skills,omitempty" db:"language_skills"`           // Added (JSONB)
}

type Guardian struct {
	BaseUUIDModel
	FirstName    string   `json:"first_name" db:"first_name"`
	LastName     *string  `json:"last_name,omitempty" db:"last_name"`
	Email        *string  `json:"email,omitempty" db:"email"`
	Phone        *string  `json:"phone,omitempty" db:"phone"`
	Profession   *string  `json:"profession,omitempty" db:"profession"`
	AnnualIncome *float64 `json:"annual_income,omitempty" db:"annual_income"`
}

type StudentGuardianMap struct {
	StudentID        uuid.UUID        `json:"student_id" db:"student_id"`
	GuardianID       uuid.UUID        `json:"guardian_id" db:"guardian_id"`
	Relationship     RelationshipType `json:"relationship" db:"relationship"`
	IsPrimaryContact bool             `json:"is_primary_contact" db:"is_primary_contact"`
	CreatedAt        time.Time        `json:"created_at" db:"created_at"`
}

type Address struct {
	BaseUUIDModel
	OwnerID      uuid.UUID   `json:"owner_id" db:"owner_id"`
	OwnerType    OwnerType   `json:"owner_type" db:"owner_type"`
	AddressType  AddressType `json:"address_type" db:"address_type"`
	AddressLine1 string      `json:"address_line_1" db:"address_line_1"`
	AddressLine2 *string     `json:"address_line_2,omitempty" db:"address_line_2"`
	CountryID    *uuid.UUID  `json:"country_id,omitempty" db:"country_id"`
	StateID      *uuid.UUID  `json:"state_id,omitempty" db:"state_id"`
	DistrictID   *uuid.UUID  `json:"district_id,omitempty" db:"district_id"`
	PostalCode   *string     `json:"postal_code,omitempty" db:"postal_code"`
}

func (i Institute) Validate() error {
	if strings.TrimSpace(i.Name) == "" {
		return errors.New("institute name cannot be empty")
	}

	if strings.TrimSpace(i.Code) == "" {
		return errors.New("institute code cannot be empty")
	}

	if i.CurrencyCode != nil && len(*i.CurrencyCode) != 3 {
		return errors.New("currency code must be 3 characters")
	}

	if i.Website != nil {
		if _, err := url.ParseRequestURI(*i.Website); err != nil {
			return errors.New("website is not a valid URL")
		}
	}

	if i.LogoURL != nil {
		if _, err := url.ParseRequestURI(*i.LogoURL); err != nil {
			return errors.New("logo URL is not valid")
		}
	}

	return nil
}
