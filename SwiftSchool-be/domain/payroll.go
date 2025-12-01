package domain

import (
	"time"

	"github.com/google/uuid"
)

type SalaryComponent struct {
	TenantUUIDModel
	Name        string `json:"name" db:"name"`
	IsDeduction bool   `json:"is_deduction" db:"is_deduction"`
}

type EmployeeSalaryConfig struct {
	BaseUUIDModel
	EmployeeID  uuid.UUID `json:"employee_id" db:"employee_id"`
	ComponentID uuid.UUID `json:"component_id" db:"component_id"`
	Amount      float64   `json:"amount" db:"amount"`
}

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
