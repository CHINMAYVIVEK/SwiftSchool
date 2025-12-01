package domain

import (
	"time"

	"github.com/google/uuid"
)

type BaseUUIDModel struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
}

type TenantUUIDModel struct {
	BaseUUIDModel
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
}

// =========================================================
// ENUMS & CONSTANTS
// =========================================================

type UserRole string

const (
	RoleStudent    UserRole = "student"
	RoleGuardian   UserRole = "guardian"
	RoleSuperAdmin UserRole = "super_admin"
	RoleAdmin      UserRole = "admin"
	RoleTeacher    UserRole = "teacher"
	RoleAccountant UserRole = "accountant"
	RoleLibrarian  UserRole = "librarian"
	RoleDriver     UserRole = "driver"
	RoleEmployee   UserRole = "employee"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

type BloodGroup string

const (
	BloodGroupAPlus   BloodGroup = "A+"
	BloodGroupAMinus  BloodGroup = "A-"
	BloodGroupBPlus   BloodGroup = "B+"
	BloodGroupBMinus  BloodGroup = "B-"
	BloodGroupOPlus   BloodGroup = "O+"
	BloodGroupOMinus  BloodGroup = "O-"
	BloodGroupABPlus  BloodGroup = "AB+"
	BloodGroupABMinus BloodGroup = "AB-"
)

type MaritalStatus string

const (
	MaritalSingle   MaritalStatus = "single"
	MaritalMarried  MaritalStatus = "married"
	MaritalDivorced MaritalStatus = "divorced"
	MaritalWidowed  MaritalStatus = "widowed"
)

type SocialCategory string

const (
	CategoryGeneral SocialCategory = "general"
	CategoryOBC     SocialCategory = "obc"
	CategorySC      SocialCategory = "sc"
	CategoryST      SocialCategory = "st"
	CategoryOther   SocialCategory = "other"
)

type VerificationStatus string

const (
	VerifyPending  VerificationStatus = "pending"
	VerifyVerified VerificationStatus = "verified"
	VerifyRejected VerificationStatus = "rejected"
)

type AddressType string

const (
	AddressCurrent   AddressType = "current"
	AddressPermanent AddressType = "permanent"
	AddressEmergency AddressType = "emergency"
)

type OwnerType string

const (
	OwnerTypeStudent   OwnerType = "student"
	OwnerTypeEmployee  OwnerType = "employee"
	OwnerTypeGuardian  OwnerType = "guardian"
	OwnerTypeInstitute OwnerType = "institute"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
	StatusHalfDay AttendanceStatus = "half_day"
	StatusOnLeave AttendanceStatus = "on_leave"
)

type LeaveStatus string

const (
	LeaveStatusPending   LeaveStatus = "pending"
	LeaveStatusApproved  LeaveStatus = "approved"
	LeaveStatusRejected  LeaveStatus = "rejected"
	LeaveStatusCancelled LeaveStatus = "cancelled"
)

type AdmissionStatus string

const (
	AdmissionStatusOpen      AdmissionStatus = "open"
	AdmissionStatusContacted AdmissionStatus = "contacted"
	AdmissionStatusApplied   AdmissionStatus = "applied"
	AdmissionStatusRejected  AdmissionStatus = "rejected"
	AdmissionStatusConverted AdmissionStatus = "converted"
)

type FeeFrequency string

const (
	FeeOneTime   FeeFrequency = "one_time"
	FeeMonthly   FeeFrequency = "monthly"
	FeeQuarterly FeeFrequency = "quarterly"
	FeeYearly    FeeFrequency = "yearly"
)

type PaymentMethod string

const (
	PaymentCash         PaymentMethod = "cash"
	PaymentCheque       PaymentMethod = "cheque"
	PaymentOnline       PaymentMethod = "online"
	PaymentUPI          PaymentMethod = "upi"
	PaymentBankTransfer PaymentMethod = "bank_transfer"
)

type FineType string

const (
	FineFixed      FineType = "fixed"
	FinePercentage FineType = "percentage"
	FineDaily      FineType = "daily"
)

type DayOfWeek string

const (
	DayMon DayOfWeek = "mon"
	DayTue DayOfWeek = "tue"
	DayWed DayOfWeek = "wed"
	DayThu DayOfWeek = "thu"
	DayFri DayOfWeek = "fri"
	DaySat DayOfWeek = "sat"
	DaySun DayOfWeek = "sun"
)

type RelationshipType string

const (
	RelFather      RelationshipType = "father"
	RelMother      RelationshipType = "mother"
	RelGuardian    RelationshipType = "guardian"
	RelSpouse      RelationshipType = "spouse"
	RelSibling     RelationshipType = "sibling"
	RelGrandparent RelationshipType = "grandparent"
)

type TripType string

const (
	TripPickup    TripType = "pickup"
	TripDrop      TripType = "drop"
	TripFieldTrip TripType = "field_trip"
)

type InventoryTransactionType string

const (
	InvTypePurchase InventoryTransactionType = "purchase"
	InvTypeConsume  InventoryTransactionType = "consume"
	InvTypeDamaged  InventoryTransactionType = "damaged"
)

type MemberType string

const (
	MemberStudent  MemberType = "student"
	MemberEmployee MemberType = "employee"
)

type HostelType string

const (
	HostelBoys  HostelType = "boys"
	HostelGirls HostelType = "girls"
	HostelStaff HostelType = "staff"
)

type DocumentType string

const (
	DocPhoto            DocumentType = "photo"
	DocAadhaar          DocumentType = "aadhaar_card"
	DocPAN              DocumentType = "pan_card"
	DocMedical          DocumentType = "medical_certificate"
	DocBirthCert        DocumentType = "birth_certificate"
	DocTC               DocumentType = "transfer_certificate"
	DocReportCard       DocumentType = "report_card"
	DocMigration        DocumentType = "migration_certificate"
	DocCharacter        DocumentType = "character_certificate"
	DocResume           DocumentType = "resume"
	DocDegreeCert       DocumentType = "degree_certificate"
	DocExperienceLetter DocumentType = "experience_letter"
	DocRelievingLetter  DocumentType = "relieving_letter"
	DocSalarySlip       DocumentType = "salary_slip"
	DocAssignment       DocumentType = "assignment"
	DocStudyMaterial    DocumentType = "study_material"
)

type AccountType string

const (
	AccountAsset     AccountType = "asset"     // Cash, Buildings, Inventory
	AccountLiability AccountType = "liability" // Loans, Accounts Payable
	AccountEquity    AccountType = "equity"    // Capital
	AccountIncome    AccountType = "income"    // Fee Collection, Grants
	AccountExpense   AccountType = "expense"   // Salaries, Purchase, Maintenance
)

type PurchaseStatus string

const (
	PurchaseDraft     PurchaseStatus = "draft"
	PurchaseOrdered   PurchaseStatus = "ordered"
	PurchaseReceived  PurchaseStatus = "received" // Inventory updated
	PurchaseCancelled PurchaseStatus = "cancelled"
)

// NEW: Language Proficiency Enum
type LanguageProficiency string

const (
	LangNative       LanguageProficiency = "native"
	LangFluent       LanguageProficiency = "fluent"
	LangIntermediate LanguageProficiency = "intermediate"
	LangBeginner     LanguageProficiency = "beginner"
)

// Currency represents the enums.currency table
type Currency struct {
	ID     int    `json:"id" db:"id"`
	Code   string `json:"code" db:"code"`
	Symbol string `json:"symbol" db:"symbol"`
}
