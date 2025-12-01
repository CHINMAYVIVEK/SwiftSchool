package domain

import (
	"github.com/google/uuid"
)

type HostelBuilding struct {
	TenantUUIDModel
	Name *string    `json:"name,omitempty" db:"name"`
	Type HostelType `json:"type,omitempty" db:"type"`
}

// Updated: Uses TenantUUIDModel
type HostelRoom struct {
	TenantUUIDModel
	BuildingID  uuid.UUID `json:"building_id" db:"building_id"`
	RoomNumber  string    `json:"room_number" db:"room_number"`
	Capacity    int       `json:"capacity" db:"capacity"`
	CostPerYear *float64  `json:"cost_per_year,omitempty" db:"cost_per_year"`
}

// Updated: Uses TenantUUIDModel
type HostelAllocation struct {
	TenantUUIDModel
	RoomID            uuid.UUID  `json:"room_id" db:"room_id"`
	StudentID         uuid.UUID  `json:"student_id" db:"student_id"`
	AcademicSessionID *uuid.UUID `json:"academic_session_id,omitempty" db:"academic_session_id"`
	IsActive          bool       `json:"is_active" db:"is_active"`
}
