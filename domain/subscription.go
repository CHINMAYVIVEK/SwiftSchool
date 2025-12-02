package domain

import (
	"time"

	"github.com/google/uuid"
)

// BillingCycle Enum
type BillingCycle string

const (
	BillingMonthly BillingCycle = "monthly"
	BillingYearly  BillingCycle = "yearly"
)

// SubscriptionStatus Enum
type SubscriptionStatus string

const (
	SubActive    SubscriptionStatus = "active"
	SubTrial     SubscriptionStatus = "trial"
	SubPastDue   SubscriptionStatus = "past_due"
	SubCancelled SubscriptionStatus = "cancelled"
)

// SaaSInvoiceStatus Enum
type SaaSInvoiceStatus string

const (
	SaaSInvoicePaid    SaaSInvoiceStatus = "paid"
	SaaSInvoicePending SaaSInvoiceStatus = "pending"
	SaaSInvoiceFailed  SaaSInvoiceStatus = "failed"
	SaaSInvoiceVoid    SaaSInvoiceStatus = "void"
)

type Plan struct {
	BaseUUIDModel
	Name            string       `json:"name" db:"name"`
	BasePrice       float64      `json:"base_price" db:"base_price"`
	PricePerStudent float64      `json:"price_per_student" db:"price_per_student"`
	BillingCycle    BillingCycle `json:"billing_cycle" db:"billing_cycle"`
	CurrencyCode    string       `json:"currency_code" db:"currency_code"`
	Features        []byte       `json:"features,omitempty" db:"features"` // JSONB
	IsActive        bool         `json:"is_active" db:"is_active"`
}

type Subscription struct {
	BaseUUIDModel                           // Not TenantUUIDModel because it links TO a tenant, but belongs to SaaS Admin context mostly
	InstituteID          uuid.UUID          `json:"institute_id" db:"institute_id"`
	PlanID               uuid.UUID          `json:"plan_id" db:"plan_id"`
	Status               SubscriptionStatus `json:"status" db:"status"`
	StartDate            time.Time          `json:"start_date" db:"start_date"`
	EndDate              time.Time          `json:"end_date" db:"end_date"`
	NextBillingDate      *time.Time         `json:"next_billing_date,omitempty" db:"next_billing_date"`
	StudentCountSnapshot int                `json:"student_count_snapshot" db:"student_count_snapshot"`
}

type SaaSInvoice struct {
	ID                 uuid.UUID         `json:"id" db:"id"`
	SubscriptionID     uuid.UUID         `json:"subscription_id" db:"subscription_id"`
	Amount             float64           `json:"amount" db:"amount"`
	Status             SaaSInvoiceStatus `json:"status" db:"status"`
	BillingPeriodStart *time.Time        `json:"billing_period_start,omitempty" db:"billing_period_start"`
	BillingPeriodEnd   *time.Time        `json:"billing_period_end,omitempty" db:"billing_period_end"`
	IssuedAt           time.Time         `json:"issued_at" db:"issued_at"`
	PaidAt             *time.Time        `json:"paid_at,omitempty" db:"paid_at"`
}
