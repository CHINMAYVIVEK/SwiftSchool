package domain

import "github.com/google/uuid"

type LeaveType struct {
	TenantUUIDModel
	Name               string `json:"name" db:"name"`
	DaysAllowedPerYear int    `json:"days_allowed_per_year" db:"days_allowed_per_year"`
}

type LeaveApplication struct {
	TenantUUIDModel
	EmployeeID  uuid.UUID   `json:"employee_id" db:"employee_id"`
	LeaveTypeID uuid.UUID   `json:"leave_type_id" db:"leave_type_id"`
	Status      LeaveStatus `json:"status" db:"status"`
	Reason      *string     `json:"reason,omitempty" db:"reason"`
}
