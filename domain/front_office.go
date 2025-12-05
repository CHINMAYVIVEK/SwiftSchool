package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: front_office.visitors
type Visitor struct {
	TenantUUIDModel
	VisitorType  VisitorType `json:"visitor_type" db:"visitor_type"`
	Name         string      `json:"name" db:"name"`
	Phone        *string     `json:"phone,omitempty" db:"phone"`
	Purpose      *string     `json:"purpose,omitempty" db:"purpose"`
	PersonToMeet *uuid.UUID  `json:"person_to_meet,omitempty" db:"person_to_meet"`
	CheckIn      time.Time   `json:"check_in" db:"check_in"`
	CheckOut     *time.Time  `json:"check_out,omitempty" db:"check_out"`
	GatePassID   *string     `json:"gate_pass_id,omitempty" db:"gate_pass_id"`
	VehicleNo    *string     `json:"vehicle_no,omitempty" db:"vehicle_no"`
	CreatedBy    *uuid.UUID  `json:"created_by,omitempty" db:"created_by"`
}

// Corresponds to schema: front_office.postal_records
type PostalRecord struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	InstituteID  uuid.UUID  `json:"institute_id" db:"institute_id"`
	Type         string     `json:"type" db:"type"` // inward, outward
	SenderName   *string    `json:"sender_name,omitempty" db:"sender_name"`
	ReceiverName *string    `json:"receiver_name,omitempty" db:"receiver_name"`
	ReferenceNo  *string    `json:"reference_no,omitempty" db:"reference_no"`
	Date         *time.Time `json:"date,omitempty" db:"date"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
}
