package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: fleet.vehicles
type Vehicle struct {
	TenantUUIDModel
	VehicleNo string  `json:"vehicle_no" db:"vehicle_no"`
	Model     *string `json:"model,omitempty" db:"model"`
	Capacity  *int    `json:"capacity,omitempty" db:"capacity"`
	FuelType  *string `json:"fuel_type,omitempty" db:"fuel_type"` // diesel, petrol
	IsActive  bool    `json:"is_active" db:"is_active"`
}

// Corresponds to schema: fleet.driver_profiles
type DriverProfile struct {
	TenantUUIDModel
	EmployeeID         uuid.UUID          `json:"employee_id" db:"employee_id"`
	LicenseNumber      *string            `json:"license_number,omitempty" db:"license_number"`
	LicenseExpiryDate  *time.Time         `json:"license_expiry_date,omitempty" db:"license_expiry_date"`
	VerificationStatus VerificationStatus `json:"verification_status" db:"verification_status"`
}

// Corresponds to schema: fleet.routes
type Route struct {
	TenantUUIDModel
	Name       string  `json:"name" db:"name"`
	StartPoint *string `json:"start_point,omitempty" db:"start_point"`
	EndPoint   *string `json:"end_point,omitempty" db:"end_point"`
}

// Corresponds to schema: fleet.route_stops
type RouteStop struct {
	TenantUUIDModel
	RouteID       uuid.UUID `json:"route_id" db:"route_id"`
	StopName      *string   `json:"stop_name,omitempty" db:"stop_name"`
	PickupTime    *string   `json:"pickup_time,omitempty" db:"pickup_time"`
	SequenceOrder *int      `json:"sequence_order,omitempty" db:"sequence_order"`
}

// Corresponds to schema: fleet.trip_logs
type TripLog struct {
	TenantUUIDModel
	VehicleID *uuid.UUID `json:"vehicle_id,omitempty" db:"vehicle_id"`
	DriverID  *uuid.UUID `json:"driver_id,omitempty" db:"driver_id"`
	RouteID   *uuid.UUID `json:"route_id,omitempty" db:"route_id"`
	TripType  TripType   `json:"trip_type" db:"trip_type"`
	TripDate  *time.Time `json:"trip_date,omitempty" db:"trip_date"`
	StartTime *time.Time `json:"start_time,omitempty" db:"start_time"`
	EndTime   *time.Time `json:"end_time,omitempty" db:"end_time"`
}

// Corresponds to schema: fleet.fuel_logs
type FuelLog struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	InstituteID     uuid.UUID  `json:"institute_id" db:"institute_id"`
	VehicleID       *uuid.UUID `json:"vehicle_id,omitempty" db:"vehicle_id"`
	FillDate        *time.Time `json:"fill_date,omitempty" db:"fill_date"`
	QuantityLitres  *float64   `json:"quantity_litres,omitempty" db:"quantity_litres"`
	Cost            *float64   `json:"cost,omitempty" db:"cost"`
	OdometerReading *float64   `json:"odometer_reading,omitempty" db:"odometer_reading"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: fleet.maintenance_logs
type MaintenanceLog struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	InstituteID     uuid.UUID  `json:"institute_id" db:"institute_id"`
	VehicleID       *uuid.UUID `json:"vehicle_id,omitempty" db:"vehicle_id"`
	MaintenanceType *string    `json:"maintenance_type,omitempty" db:"maintenance_type"`
	Description     *string    `json:"description,omitempty" db:"description"`
	Cost            *float64   `json:"cost,omitempty" db:"cost"`
	ServiceDate     *time.Time `json:"service_date,omitempty" db:"service_date"`
	NextServiceDue  *time.Time `json:"next_service_due,omitempty" db:"next_service_due"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
}
