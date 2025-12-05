package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: finance.bank_accounts
type BankAccount struct {
	TenantUUIDModel
	AccountName    *string  `json:"account_name,omitempty" db:"account_name"`
	AccountNumber  *string  `json:"account_number,omitempty" db:"account_number"`
	BankName       *string  `json:"bank_name,omitempty" db:"bank_name"`
	IFSCCode       *string  `json:"ifsc_code,omitempty" db:"ifsc_code"`
	BranchName     *string  `json:"branch_name,omitempty" db:"branch_name"`
	OpeningBalance *float64 `json:"opening_balance,omitempty" db:"opening_balance"`
	CurrentBalance *float64 `json:"current_balance,omitempty" db:"current_balance"`
	IsActive       bool     `json:"is_active" db:"is_active"`
}

// Corresponds to schema: finance.accounts
type Account struct {
	TenantUUIDModel
	Name            string      `json:"name" db:"name"`
	Code            string      `json:"code" db:"code"`
	ParentAccountID *uuid.UUID  `json:"parent_account_id,omitempty" db:"parent_account_id"`
	Type            AccountType `json:"type" db:"type"` // asset, liability, income...
	IsSystem        bool        `json:"is_system" db:"is_system"`
}

// Corresponds to schema: finance.taxes
type Tax struct {
	ID          uuid.UUID `json:"id" db:"id"`
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
	Name        *string   `json:"name,omitempty" db:"name"`
	Percentage  float64   `json:"percentage" db:"percentage"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Corresponds to schema: finance.fee_heads
type FeeHead struct {
	TenantUUIDModel
	Name              string    `json:"name" db:"name"`
	IsRefundable      bool      `json:"is_refundable" db:"is_refundable"`
	LinkedGLAccountID uuid.UUID `json:"linked_gl_account_id" db:"linked_gl_account_id"`
}

// Corresponds to schema: finance.fine_rules
type FineRule struct {
	TenantUUIDModel
	Name       string   `json:"name" db:"name"`
	GraceDays  int      `json:"grace_days" db:"grace_days"`
	FineType   FineType `json:"fine_type" db:"fine_type"`
	FineAmount float64  `json:"fine_amount" db:"fine_amount"`
	IsActive   bool     `json:"is_active" db:"is_active"`
}

// Corresponds to schema: finance.concessions
type Concession struct {
	TenantUUIDModel
	Name  *string  `json:"name,omitempty" db:"name"`
	Type  string   `json:"type" db:"type"` // flat, percentage
	Value *float64 `json:"value,omitempty" db:"value"`
}

// Corresponds to schema: finance.fee_structures
type FeeStructure struct {
	TenantUUIDModel
	AcademicSessionID uuid.UUID    `json:"academic_session_id" db:"academic_session_id"`
	ClassID           *uuid.UUID   `json:"class_id,omitempty" db:"class_id"`
	FeeHeadID         uuid.UUID    `json:"fee_head_id" db:"fee_head_id"`
	Amount            float64      `json:"amount" db:"amount"`
	Frequency         FeeFrequency `json:"frequency" db:"frequency"`
}

// Corresponds to schema: finance.budgets
type Budget struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	InstituteID       uuid.UUID  `json:"institute_id" db:"institute_id"`
	AcademicSessionID *uuid.UUID `json:"academic_session_id,omitempty" db:"academic_session_id"`
	DepartmentID      *uuid.UUID `json:"department_id,omitempty" db:"department_id"`
	ExpenseCategoryID *uuid.UUID `json:"expense_category_id,omitempty" db:"expense_category_id"`
	AllocatedAmount   float64    `json:"allocated_amount" db:"allocated_amount"`
	UtilizedAmount    float64    `json:"utilized_amount" db:"utilized_amount"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: finance.student_wallets
type StudentWallet struct {
	ID          uuid.UUID `json:"id" db:"id"`
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
	StudentID   uuid.UUID `json:"student_id" db:"student_id"`
	Balance     float64   `json:"balance" db:"balance"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Corresponds to schema: finance.invoices
type Invoice struct {
	TenantUUIDModel
	InvoiceNo      string     `json:"invoice_no" db:"invoice_no"`
	StudentID      uuid.UUID  `json:"student_id" db:"student_id"`
	TotalAmount    float64    `json:"total_amount" db:"total_amount"`
	DiscountAmount float64    `json:"discount_amount" db:"discount_amount"`
	FineAmount     float64    `json:"fine_amount" db:"fine_amount"`
	PaidAmount     float64    `json:"paid_amount" db:"paid_amount"`
	Status         string     `json:"status" db:"status"` // pending, partial, paid
	DueDate        *time.Time `json:"due_date,omitempty" db:"due_date"`
}

// Corresponds to schema: finance.invoice_items
type InvoiceItem struct {
	TenantUUIDModel
	InvoiceID       uuid.UUID  `json:"invoice_id" db:"invoice_id"`
	FeeHeadID       *uuid.UUID `json:"fee_head_id,omitempty" db:"fee_head_id"`
	Amount          float64    `json:"amount" db:"amount"`
	DiscountApplied float64    `json:"discount_applied" db:"discount_applied"`
	ConcessionID    *uuid.UUID `json:"concession_id,omitempty" db:"concession_id"`
	Description     *string    `json:"description,omitempty" db:"description"`
}

// Corresponds to schema: finance.transactions
type Transaction struct {
	TenantUUIDModel
	InvoiceID        *uuid.UUID  `json:"invoice_id,omitempty" db:"invoice_id"`
	StudentID        *uuid.UUID  `json:"student_id,omitempty" db:"student_id"`
	TransactionRefNo *string     `json:"transaction_ref_no,omitempty" db:"transaction_ref_no"`
	PaymentMode      PaymentMode `json:"payment_mode" db:"payment_mode"`
	ChequeNo         *string     `json:"cheque_no,omitempty" db:"cheque_no"`
	ChequeDate       *time.Time  `json:"cheque_date,omitempty" db:"cheque_date"`
	BankName         *string     `json:"bank_name,omitempty" db:"bank_name"`
	ChequeStatus     *string     `json:"cheque_status,omitempty" db:"cheque_status"`
	Amount           float64     `json:"amount" db:"amount"`
	IsWalletUsage    bool        `json:"is_wallet_usage" db:"is_wallet_usage"`
	PaymentDate      *time.Time  `json:"payment_date,omitempty" db:"payment_date"`
	Status           string      `json:"status" db:"status"`
	CollectedBy      *uuid.UUID  `json:"collected_by,omitempty" db:"collected_by"`
}

// Corresponds to schema: finance.refunds
type Refund struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	InstituteID uuid.UUID    `json:"institute_id" db:"institute_id"`
	StudentID   *uuid.UUID   `json:"student_id,omitempty" db:"student_id"`
	InvoiceID   *uuid.UUID   `json:"invoice_id,omitempty" db:"invoice_id"`
	Amount      float64      `json:"amount" db:"amount"`
	Reason      *string      `json:"reason,omitempty" db:"reason"`
	Status      RefundStatus `json:"status" db:"status"`
	RefundDate  *time.Time   `json:"refund_date,omitempty" db:"refund_date"`
	ProcessedBy *uuid.UUID   `json:"processed_by,omitempty" db:"processed_by"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// Corresponds to schema: finance.journal_entries
type JournalEntry struct {
	TenantUUIDModel
	ReferenceNo     string     `json:"reference_no" db:"reference_no"`
	TransactionDate time.Time  `json:"transaction_date" db:"transaction_date"`
	Description     *string    `json:"description,omitempty" db:"description"`
	IsPosted        bool       `json:"is_posted" db:"is_posted"`
	PostedAt        *time.Time `json:"posted_at,omitempty" db:"posted_at"`
}

// Corresponds to schema: finance.journal_items
type JournalItem struct {
	TenantUUIDModel
	JournalEntryID uuid.UUID `json:"journal_entry_id" db:"journal_entry_id"`
	AccountID      uuid.UUID `json:"account_id" db:"account_id"`
	Debit          float64   `json:"debit" db:"debit"`
	Credit         float64   `json:"credit" db:"credit"`
	Description    *string   `json:"description,omitempty" db:"description"`
}

// Corresponds to schema: finance.bank_statement_entries
type BankStatementEntry struct {
	ID                      uuid.UUID  `json:"id" db:"id"`
	InstituteID             uuid.UUID  `json:"institute_id" db:"institute_id"`
	BankAccountID           *uuid.UUID `json:"bank_account_id,omitempty" db:"bank_account_id"`
	TransactionDate         *time.Time `json:"transaction_date,omitempty" db:"transaction_date"`
	ValueDate               *time.Time `json:"value_date,omitempty" db:"value_date"`
	Description             *string    `json:"description,omitempty" db:"description"`
	WithdrawalAmount        float64    `json:"withdrawal_amount" db:"withdrawal_amount"`
	DepositAmount           float64    `json:"deposit_amount" db:"deposit_amount"`
	Balance                 *float64   `json:"balance,omitempty" db:"balance"`
	ReconciledTransactionID *uuid.UUID `json:"reconciled_transaction_id,omitempty" db:"reconciled_transaction_id"`
	IsReconciled            bool       `json:"is_reconciled" db:"is_reconciled"`
	CreatedAt               time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: finance.vendors
type Vendor struct {
	TenantUUIDModel
	Name        string  `json:"name" db:"name"`
	ContactName *string `json:"contact_name,omitempty" db:"contact_name"`
	Phone       *string `json:"phone,omitempty" db:"phone"`
	Email       *string `json:"email,omitempty" db:"email"`
	Address     *string `json:"address,omitempty" db:"address"`
	GSTNumber   *string `json:"gst_number,omitempty" db:"gst_number"`
}

// Corresponds to schema: finance.purchase_orders
type PurchaseOrder struct {
	TenantUUIDModel
	VendorID    uuid.UUID      `json:"vendor_id" db:"vendor_id"`
	OrderDate   time.Time      `json:"order_date" db:"order_date"`
	TotalAmount float64        `json:"total_amount" db:"total_amount"`
	Status      PurchaseStatus `json:"status" db:"status"`
	ReferenceNo *string        `json:"reference_no,omitempty" db:"reference_no"`
}

// Corresponds to schema: finance.purchase_order_items
type PurchaseOrderItem struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	InstituteID     uuid.UUID  `json:"institute_id" db:"institute_id"`
	PurchaseOrderID uuid.UUID  `json:"purchase_order_id" db:"purchase_order_id"`
	ItemID          uuid.UUID  `json:"item_id" db:"item_id"`
	Quantity        int        `json:"quantity" db:"quantity"`
	UnitPrice       float64    `json:"unit_price" db:"unit_price"`
	TaxID           *uuid.UUID `json:"tax_id,omitempty" db:"tax_id"`
	TaxAmount       float64    `json:"tax_amount" db:"tax_amount"`
	TotalAmount     float64    `json:"total_amount" db:"total_amount"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: finance.purchase_items
type PurchaseItem struct {
	TenantUUIDModel
	PurchaseOrderID uuid.UUID `json:"purchase_order_id" db:"purchase_order_id"`
	ItemID          uuid.UUID `json:"item_id" db:"item_id"`
	Quantity        int       `json:"quantity" db:"quantity"`
	UnitPrice       float64   `json:"unit_price" db:"unit_price"`
}
