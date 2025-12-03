package inventory

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}

//////////////////////////////////////////////////////
//                    REPOSITORY                    //
//////////////////////////////////////////////////////

type Repository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *Repository {
	return &Repository{db: db}
}

//////////////////////////////////////////////////////
//                     SERVICE                      //
//////////////////////////////////////////////////////

type Service struct {
	repo RepositoryInterface
}

func NewService(db *helper.PostgresWrapper) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

//////////////////////////////////////////////////////
//               REPOSITORY INTERFACE               //
//////////////////////////////////////////////////////

type RepositoryInterface interface {
	CreateItemCategory(ctx context.Context, arg domain.ItemCategory) (*domain.ItemCategory, error)
	CreateInventoryItem(ctx context.Context, arg domain.InventoryItem) (*domain.InventoryItem, error)
	ListInventoryItems(ctx context.Context, instituteID uuid.UUID) ([]*domain.InventoryItem, error)
	UpdateStock(ctx context.Context, itemID, instituteID uuid.UUID, quantityChange int) error
	CreateInventoryTransaction(ctx context.Context, arg domain.InventoryTransaction) (*domain.InventoryTransaction, error)
	ListInventoryTransactions(ctx context.Context, instituteID uuid.UUID) ([]*domain.InventoryTransaction, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateItemCategory(ctx context.Context, arg domain.ItemCategory) (*domain.ItemCategory, error)
	CreateInventoryItem(ctx context.Context, arg domain.InventoryItem) (*domain.InventoryItem, error)
	ListInventoryItems(ctx context.Context, instituteID uuid.UUID) ([]*domain.InventoryItem, error)
	UpdateStock(ctx context.Context, itemID, instituteID uuid.UUID, quantityChange int) error
	CreateInventoryTransaction(ctx context.Context, arg domain.InventoryTransaction) (*domain.InventoryTransaction, error)
	ListInventoryTransactions(ctx context.Context, instituteID uuid.UUID) ([]*domain.InventoryTransaction, error)
}
