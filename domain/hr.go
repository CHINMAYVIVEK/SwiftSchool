package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: hr.leave_types
type LeaveType struct {
	TenantUUIDModel
	Name               string `json:"name" db:"name"`
	DaysAllowedPerYear int    `json:"days_allowed_per_year" db:"days_allowed_per_year"`
}

// Corresponds to schema: hr.leave_applications
type LeaveApplication struct {
	TenantUUIDModel
	EmployeeID  uuid.UUID   `json:"employee_id" db:"employee_id"`
	LeaveTypeID uuid.UUID   `json:"leave_type_id" db:"leave_type_id"`
	Status      LeaveStatus `json:"status" db:"status"`
	StartDate   time.Time   `json:"start_date" db:"start_date"`
	EndDate     *time.Time  `json:"end_date,omitempty" db:"end_date"`
	Reason      *string     `json:"reason,omitempty" db:"reason"`
	ApprovedBy  *uuid.UUID  `json:"approved_by,omitempty" db:"approved_by"`
}

// Corresponds to schema: hr.attendance_devices
type AttendanceDevice struct {
	ID          uuid.UUID `json:"id" db:"id"`
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
	DeviceName  *string   `json:"device_name,omitempty" db:"device_name"`
	IPAddress   *string   `json:"ip_address,omitempty" db:"ip_address"`
	Location    *string   `json:"location,omitempty" db:"location"`
	APIKey      *string   `json:"api_key,omitempty" db:"api_key"`
	IsActive    bool      `json:"is_active" db:"is_active"`
}

// Corresponds to schema: hr.biometric_logs
type BiometricLog struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	DeviceID    *uuid.UUID `json:"device_id,omitempty" db:"device_id"`
	UserID      *uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	PunchTime   *time.Time `json:"punch_time,omitempty" db:"punch_time"`
	PunchType   *string    `json:"punch_type,omitempty" db:"punch_type"` // IN, OUT
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
