package domain

import (
	"time"

	"github.com/google/uuid"
)

type ItemCategory struct {
	TenantUUIDModel
	Name *string `json:"name,omitempty" db:"name"`
}

type InventoryItem struct {
	TenantUUIDModel
	Name         string     `json:"name" db:"name"`
	CategoryID   *uuid.UUID `json:"category_id,omitempty" db:"category_id"`
	CurrentStock int        `json:"current_stock" db:"current_stock"`
	ReorderLevel int        `json:"reorder_level" db:"reorder_level"`
}

type InventoryTransaction struct {
	TenantUUIDModel
	ItemID          *uuid.UUID               `json:"item_id,omitempty" db:"item_id"`
	TransactionType InventoryTransactionType `json:"transaction_type" db:"transaction_type"`
	Quantity        int                      `json:"quantity" db:"quantity"`
	UnitPrice       *float64                 `json:"unit_price,omitempty" db:"unit_price"`
	TransactionDate *time.Time               `json:"transaction_date,omitempty" db:"transaction_date"`
	Remarks         *string                  `json:"remarks,omitempty" db:"remarks"`
}
