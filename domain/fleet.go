package domain

import (
	"github.com/google/uuid"
)

type Vehicle struct {
	TenantUUIDModel
	VehicleNo string  `json:"vehicle_no" db:"vehicle_no"`
	Capacity  *int    `json:"capacity,omitempty" db:"capacity"`
	Model     *string `json:"model,omitempty" db:"model"`
	IsActive  bool    `json:"is_active" db:"is_active"`
}

type DriverProfile struct {
	TenantUUIDModel
	EmployeeID         uuid.UUID          `json:"employee_id" db:"employee_id"`
	LicenseNumber      *string            `json:"license_number,omitempty" db:"license_number"`
	VerificationStatus VerificationStatus `json:"verification_status" db:"verification_status"`
}

type Route struct {
	TenantUUIDModel
	Name       string  `json:"name" db:"name"`
	StartPoint *string `json:"start_point,omitempty" db:"start_point"`
	EndPoint   *string `json:"end_point,omitempty" db:"end_point"`
}

// Updated: Uses TenantUUIDModel
type RouteStop struct {
	TenantUUIDModel
	RouteID       uuid.UUID `json:"route_id" db:"route_id"`
	StopName      *string   `json:"stop_name,omitempty" db:"stop_name"`
	PickupTime    *string   `json:"pickup_time,omitempty" db:"pickup_time"`
	SequenceOrder *int      `json:"sequence_order,omitempty" db:"sequence_order"`
}

type TripLog struct {
	TenantUUIDModel
	VehicleID *uuid.UUID `json:"vehicle_id,omitempty" db:"vehicle_id"`
	DriverID  *uuid.UUID `json:"driver_id,omitempty" db:"driver_id"`
	RouteID   *uuid.UUID `json:"route_id,omitempty" db:"route_id"`
	TripType  TripType   `json:"trip_type" db:"trip_type"`
}
