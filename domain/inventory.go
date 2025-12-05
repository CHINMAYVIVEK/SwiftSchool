package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: inventory.item_categories
type ItemCategory struct {
	TenantUUIDModel
	Name *string `json:"name,omitempty" db:"name"`
}

// Corresponds to schema: inventory.items
type InventoryItem struct {
	TenantUUIDModel
	Name         string     `json:"name" db:"name"`
	CategoryID   *uuid.UUID `json:"category_id,omitempty" db:"category_id"`
	IsFixedAsset bool       `json:"is_fixed_asset" db:"is_fixed_asset"`
	ReorderLevel int        `json:"reorder_level" db:"reorder_level"`
}

// Corresponds to schema: inventory.locations
type Location struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	Name        *string    `json:"name,omitempty" db:"name"`
	InChargeID  *uuid.UUID `json:"in_charge_id,omitempty" db:"in_charge_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: inventory.stock_levels
type StockLevel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	InstituteID uuid.UUID `json:"institute_id" db:"institute_id"`
	ItemID      uuid.UUID `json:"item_id" db:"item_id"`
	LocationID  uuid.UUID `json:"location_id" db:"location_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Corresponds to schema: inventory.stock_transfers
type StockTransfer struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	InstituteID    uuid.UUID  `json:"institute_id" db:"institute_id"`
	ItemID         uuid.UUID  `json:"item_id" db:"item_id"`
	FromLocationID uuid.UUID  `json:"from_location_id" db:"from_location_id"`
	ToLocationID   uuid.UUID  `json:"to_location_id" db:"to_location_id"`
	Quantity       int        `json:"quantity" db:"quantity"`
	MovedBy        *uuid.UUID `json:"moved_by,omitempty" db:"moved_by"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: inventory.fixed_assets
type FixedAsset struct {
	TenantUUIDModel
	ItemID       uuid.UUID  `json:"item_id" db:"item_id"`
	AssetTagID   *string    `json:"asset_tag_id,omitempty" db:"asset_tag_id"`
	SerialNumber *string    `json:"serial_number,omitempty" db:"serial_number"`
	LocationRoom *string    `json:"location_room,omitempty" db:"location_room"`
	PurchaseDate *time.Time `json:"purchase_date,omitempty" db:"purchase_date"`
	PurchaseCost *float64   `json:"purchase_cost,omitempty" db:"purchase_cost"`
	Status       string     `json:"status" db:"status"` // active, retired
}

// Corresponds to schema: inventory.requisitions
type Requisition struct {
	TenantUUIDModel
	RequestedBy  uuid.UUID  `json:"requested_by" db:"requested_by"`
	DepartmentID *uuid.UUID `json:"department_id,omitempty" db:"department_id"`
	Status       string     `json:"status" db:"status"`
	RequestDate  *time.Time `json:"request_date,omitempty" db:"request_date"`
}

// Corresponds to schema: inventory.requisition_items
type RequisitionItem struct {
	ID            uuid.UUID `json:"id" db:"id"`
	InstituteID   uuid.UUID `json:"institute_id" db:"institute_id"`
	RequisitionID uuid.UUID `json:"requisition_id" db:"requisition_id"`
	ItemName      *string   `json:"item_name,omitempty" db:"item_name"`
	Quantity      *int      `json:"quantity,omitempty" db:"quantity"`
	Remarks       *string   `json:"remarks,omitempty" db:"remarks"`
}

// Corresponds to schema: inventory.transactions
type InventoryTransaction struct {
	TenantUUIDModel
	ItemID          *uuid.UUID               `json:"item_id,omitempty" db:"item_id"`
	TransactionType InventoryTransactionType `json:"transaction_type" db:"transaction_type"`
	Quantity        int                      `json:"quantity" db:"quantity"`
	UnitPrice       *float64                 `json:"unit_price,omitempty" db:"unit_price"`
	TransactionDate *time.Time               `json:"transaction_date,omitempty" db:"transaction_date"`
	Remarks         *string                  `json:"remarks,omitempty" db:"remarks"`
}
