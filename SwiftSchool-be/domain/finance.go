package domain

import (
	"time"

	"github.com/google/uuid"
)

// --- FEES ---

type FeeHead struct {
	TenantUUIDModel
	Name string `json:"name" db:"name"`
}

type FineRule struct {
	TenantUUIDModel
	Name       string   `json:"name" db:"name"`
	GraceDays  int      `json:"grace_days" db:"grace_days"`
	FineType   FineType `json:"fine_type" db:"fine_type"`
	FineAmount float64  `json:"fine_amount" db:"fine_amount"`
	IsActive   bool     `json:"is_active" db:"is_active"`
}

type FeeStructure struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID    `json:"academic_session_id" db:"academic_session_id"`
	ClassID           *uuid.UUID   `json:"class_id,omitempty" db:"class_id"`
	FeeHeadID         uuid.UUID    `json:"fee_head_id" db:"fee_head_id"`
	Amount            float64      `json:"amount" db:"amount"`
	Frequency         FeeFrequency `json:"frequency" db:"frequency"`
}

type Invoice struct {
	TenantUUIDModel
	InvoiceNo   string     `json:"invoice_no" db:"invoice_no"`
	StudentID   uuid.UUID  `json:"student_id" db:"student_id"`
	TotalAmount float64    `json:"total_amount" db:"total_amount"`
	PaidAmount  float64    `json:"paid_amount" db:"paid_amount"`
	Status      string     `json:"status" db:"status"`
	DueDate     *time.Time `json:"due_date,omitempty" db:"due_date"`
}

// Updated: Uses TenantUUIDModel
type InvoiceItem struct {
	TenantUUIDModel
	InvoiceID   uuid.UUID  `json:"invoice_id" db:"invoice_id"`
	FeeHeadID   *uuid.UUID `json:"fee_head_id,omitempty" db:"fee_head_id"`
	Description *string    `json:"description,omitempty" db:"description"`
	Amount      float64    `json:"amount" db:"amount"`
}

type Transaction struct {
	TenantUUIDModel
	InvoiceID        *uuid.UUID    `json:"invoice_id,omitempty" db:"invoice_id"`
	TransactionRefNo *string       `json:"transaction_ref_no,omitempty" db:"transaction_ref_no"`
	PaymentMethod    PaymentMethod `json:"payment_method,omitempty" db:"payment_method"`
	Amount           float64       `json:"amount" db:"amount"`
	PaymentDate      *time.Time    `json:"payment_date,omitempty" db:"payment_date"`
	Status           string        `json:"status" db:"status"`
	CollectedBy      *uuid.UUID    `json:"collected_by,omitempty" db:"collected_by"`
}

// --- ACCOUNTING ---

type Account struct {
	TenantUUIDModel
	Name        string      `json:"name" db:"name"`
	Code        string      `json:"code" db:"code"`
	Type        AccountType `json:"type" db:"type"`
	Description *string     `json:"description,omitempty" db:"description"`
	IsActive    bool        `json:"is_active" db:"is_active"`
}

type JournalEntry struct {
	TenantUUIDModel
	ReferenceNo     string     `json:"reference_no" db:"reference_no"`
	TransactionDate time.Time  `json:"transaction_date" db:"transaction_date"`
	Description     *string    `json:"description,omitempty" db:"description"`
	IsPosted        bool       `json:"is_posted" db:"is_posted"`
	PostedAt        *time.Time `json:"posted_at,omitempty" db:"posted_at"`
}

type JournalItem struct {
	TenantUUIDModel
	JournalEntryID uuid.UUID `json:"journal_entry_id" db:"journal_entry_id"`
	AccountID      uuid.UUID `json:"account_id" db:"account_id"`
	Debit          float64   `json:"debit" db:"debit"`
	Credit         float64   `json:"credit" db:"credit"`
	Description    *string   `json:"description,omitempty" db:"description"`
}

type Vendor struct {
	TenantUUIDModel
	Name        string  `json:"name" db:"name"`
	ContactName *string `json:"contact_name,omitempty" db:"contact_name"`
	Phone       *string `json:"phone,omitempty" db:"phone"`
	Email       *string `json:"email,omitempty" db:"email"`
	Address     *string `json:"address,omitempty" db:"address"`
	GSTNumber   *string `json:"gst_number,omitempty" db:"gst_number"`
}

type PurchaseOrder struct {
	TenantUUIDModel
	VendorID    uuid.UUID      `json:"vendor_id" db:"vendor_id"`
	OrderDate   time.Time      `json:"order_date" db:"order_date"`
	TotalAmount float64        `json:"total_amount" db:"total_amount"`
	Status      PurchaseStatus `json:"status" db:"status"`
	ReferenceNo *string        `json:"reference_no,omitempty" db:"reference_no"`
}

type PurchaseItem struct {
	TenantUUIDModel
	PurchaseOrderID uuid.UUID `json:"purchase_order_id" db:"purchase_order_id"`
	ItemID          uuid.UUID `json:"item_id" db:"item_id"`
	Quantity        int       `json:"quantity" db:"quantity"`
	UnitPrice       float64   `json:"unit_price" db:"unit_price"`
}
