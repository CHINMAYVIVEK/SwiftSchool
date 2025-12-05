package domain

import (
	"time"

	"github.com/google/uuid"
)

// BaseUUIDModel serves as the foundational struct for all entities.
// It handles the primary key (UUID) and standard audit trails.
type BaseUUIDModel struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Soft Delete support
	CreatedBy *uuid.UUID `json:"created_by,omitempty" db:"created_by"` // Audit trail
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"` // Audit trail
}

// TenantUUIDModel enforces Strict Multi-Tenancy.
// CRITICAL: Every business entity struct MUST embed this.
// It ensures that data is always scoped to a specific Institute.
type TenantUUIDModel struct {
	BaseUUIDModel
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
}

// Currency represents the global currency lookup (enums.currency)
type Currency struct {
	ID     int    `json:"id" db:"id"`
	Code   string `json:"code" db:"code"`
	Symbol string `json:"symbol" db:"symbol"`
}

// =========================================================
// ENUMS & CONSTANTS (Single Source of Truth)
// =========================================================

// --- AUTH & ROLES ---

type UserRole string

const (
	RoleSuperAdmin UserRole = "super_admin"
	RoleAdmin      UserRole = "admin"
	RoleTeacher    UserRole = "teacher"
	RoleAccountant UserRole = "accountant"
	RoleLibrarian  UserRole = "librarian"
	RoleDriver     UserRole = "driver"
	RoleEmployee   UserRole = "employee"
	RoleStudent    UserRole = "student"
	RoleGuardian   UserRole = "guardian"
	RoleNurse      UserRole = "nurse"
)

// --- DEMOGRAPHICS ---

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

type LanguageProficiency string

const (
	LangNative       LanguageProficiency = "native"
	LangFluent       LanguageProficiency = "fluent"
	LangIntermediate LanguageProficiency = "intermediate"
	LangBeginner     LanguageProficiency = "beginner"
)

// --- CONTACT & ADDRESS ---

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

type RelationshipType string

const (
	RelFather      RelationshipType = "father"
	RelMother      RelationshipType = "mother"
	RelGuardian    RelationshipType = "guardian"
	RelSpouse      RelationshipType = "spouse"
	RelSibling     RelationshipType = "sibling"
	RelGrandparent RelationshipType = "grandparent"
)

// --- ACADEMIC & ATTENDANCE ---

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
	StatusHalfDay AttendanceStatus = "half_day"
	StatusOnLeave AttendanceStatus = "on_leave"
)

type AdmissionStatus string

const (
	AdmissionStatusOpen      AdmissionStatus = "open"
	AdmissionStatusContacted AdmissionStatus = "contacted"
	AdmissionStatusApplied   AdmissionStatus = "applied"
	AdmissionStatusRejected  AdmissionStatus = "rejected"
	AdmissionStatusConverted AdmissionStatus = "converted"
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

// --- FINANCE ---

type AccountType string

const (
	AccAsset     AccountType = "asset"     // Cash, Buildings, Inventory
	AccLiability AccountType = "liability" // Loans, Accounts Payable
	AccEquity    AccountType = "equity"    // Capital
	AccIncome    AccountType = "income"    // Fee Collection, Grants
	AccExpense   AccountType = "expense"   // Salaries, Purchase, Maintenance
)

type PaymentMode string

const (
	PaymentCash         PaymentMode = "cash"
	PaymentCheque       PaymentMode = "cheque"
	PaymentOnline       PaymentMode = "online"
	PaymentUPI          PaymentMode = "upi"
	PaymentBankTransfer PaymentMode = "bank_transfer"
)

type FeeFrequency string

const (
	FeeOneTime   FeeFrequency = "one_time"
	FeeMonthly   FeeFrequency = "monthly"
	FeeQuarterly FeeFrequency = "quarterly"
	FeeYearly    FeeFrequency = "yearly"
)

type FineType string

const (
	FineFixed      FineType = "fixed"
	FinePercentage FineType = "percentage"
	FineDaily      FineType = "daily"
)

type PurchaseStatus string

const (
	PurchaseDraft     PurchaseStatus = "draft"
	PurchaseOrdered   PurchaseStatus = "ordered"
	PurchaseReceived  PurchaseStatus = "received"
	PurchaseCancelled PurchaseStatus = "cancelled"
)

type RefundStatus string

const (
	RefundRequested RefundStatus = "requested"
	RefundApproved  RefundStatus = "approved"
	RefundProcessed RefundStatus = "processed"
	RefundRejected  RefundStatus = "rejected"
)

// --- HR & OPERATIONS ---

type LeaveStatus string

const (
	LeavePending   LeaveStatus = "pending"
	LeaveApproved  LeaveStatus = "approved"
	LeaveRejected  LeaveStatus = "rejected"
	LeaveCancelled LeaveStatus = "cancelled"
)

type LoanStatus string

const (
	LoanRequested LoanStatus = "requested"
	LoanApproved  LoanStatus = "approved"
	LoanActive    LoanStatus = "active"
	LoanClosed    LoanStatus = "closed"
	LoanRejected  LoanStatus = "rejected"
)

type VerificationStatus string

const (
	VerifyPending  VerificationStatus = "pending"
	VerifyVerified VerificationStatus = "verified"
	VerifyRejected VerificationStatus = "rejected"
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
	InvTypeReturn   InventoryTransactionType = "return"
	InvTypeAdjust   InventoryTransactionType = "adjustment"
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

type MealType string

const (
	MealBreakfast MealType = "breakfast"
	MealLunch     MealType = "lunch"
	MealSnacks    MealType = "snacks"
	MealDinner    MealType = "dinner"
)

// --- CRM & SUPPORT ---

type TicketPriority string

const (
	PriorityLow      TicketPriority = "low"
	PriorityMedium   TicketPriority = "medium"
	PriorityHigh     TicketPriority = "high"
	PriorityCritical TicketPriority = "critical"
)

type TicketStatus string

const (
	TicketOpen       TicketStatus = "open"
	TicketInProgress TicketStatus = "in_progress"
	TicketResolved   TicketStatus = "resolved"
	TicketClosed     TicketStatus = "closed"
)

type IncidentSeverity string

const (
	SeverityMinor    IncidentSeverity = "minor"
	SeverityMajor    IncidentSeverity = "major"
	SeverityCritical IncidentSeverity = "critical"
)

type VisitorType string

const (
	VisitorParent   VisitorType = "parent"
	VisitorVendor   VisitorType = "vendor"
	VisitorOfficial VisitorType = "official"
	VisitorOther    VisitorType = "other"
)

// --- DOCUMENTS ---

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
	DocPrescription     DocumentType = "prescription"
	DocCircular         DocumentType = "circular"
)
