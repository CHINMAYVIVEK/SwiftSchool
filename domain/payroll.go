package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: payroll.salary_components
type SalaryComponent struct {
	TenantUUIDModel
	Name        string `json:"name" db:"name"`
	IsDeduction bool   `json:"is_deduction" db:"is_deduction"`
}

// Corresponds to schema: payroll.employee_salary_config
type EmployeeSalaryConfig struct {
	BaseUUIDModel
	EmployeeID  uuid.UUID `json:"employee_id" db:"employee_id"`
	ComponentID uuid.UUID `json:"component_id" db:"component_id"`
	Amount      float64   `json:"amount" db:"amount"`
}

// Corresponds to schema: payroll.loans
type Loan struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	InstituteID       uuid.UUID  `json:"institute_id" db:"institute_id"`
	EmployeeID        *uuid.UUID `json:"employee_id,omitempty" db:"employee_id"`
	Amount            float64    `json:"amount" db:"amount"`
	InterestRate      float64    `json:"interest_rate" db:"interest_rate"`
	InstallmentAmount *float64   `json:"installment_amount,omitempty" db:"installment_amount"`
	TotalInstallments *int       `json:"total_installments,omitempty" db:"total_installments"`
	PaidInstallments  int        `json:"paid_installments" db:"paid_installments"`
	Status            LoanStatus `json:"status" db:"status"`
	ApprovedBy        *uuid.UUID `json:"approved_by,omitempty" db:"approved_by"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: payroll.payslips
type Payslip struct {
	TenantUUIDModel
	EmployeeID      uuid.UUID  `json:"employee_id" db:"employee_id"`
	MonthYear       time.Time  `json:"month_year" db:"month_year"`
	GrossSalary     float64    `json:"gross_salary" db:"gross_salary"`
	TotalDeductions float64    `json:"total_deductions" db:"total_deductions"`
	NetSalary       float64    `json:"net_salary" db:"net_salary"`
	IsPaid          bool       `json:"is_paid" db:"is_paid"`
	PaymentDate     *time.Time `json:"payment_date,omitempty" db:"payment_date"`
}
