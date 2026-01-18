package common

import (
	"context"
	"swiftschool/domain"
	"swiftschool/internal/database"

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
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

//////////////////////////////////////////////////////
//                     SERVICE                      //
//////////////////////////////////////////////////////

type Service struct {
	repo RepositoryInterface
}

func NewService(db *database.Database) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

//////////////////////////////////////////////////////
//               REPOSITORY INTERFACE               //
//////////////////////////////////////////////////////

type RepositoryInterface interface {
	// ========================= DOCS =========================
	CreateDocument(ctx context.Context, arg domain.Document) (*domain.Document, error)
	ListDocuments(ctx context.Context, instituteID, ownerID uuid.UUID) ([]*domain.Document, error)

	// ========================= COMMS =========================
	CreateNotification(ctx context.Context, arg domain.Notification) (*domain.Notification, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= DOCS =========================
	CreateDocument(ctx context.Context, arg domain.Document) (*domain.Document, error)
	ListDocuments(ctx context.Context, instituteID, ownerID uuid.UUID) ([]*domain.Document, error)

	// ========================= COMMS =========================
	CreateNotification(ctx context.Context, arg domain.Notification) (*domain.Notification, error)
}
