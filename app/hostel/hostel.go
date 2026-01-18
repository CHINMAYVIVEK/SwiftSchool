package hostel

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
	CreateHostelBuilding(ctx context.Context, arg domain.HostelBuilding) (*domain.HostelBuilding, error)
	CreateHostelRoom(ctx context.Context, arg domain.HostelRoom) (*domain.HostelRoom, error)
	AllocateRoom(ctx context.Context, arg domain.HostelAllocation) (*domain.HostelAllocation, error)
	VacateRoom(ctx context.Context, id, instituteID uuid.UUID) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateHostelBuilding(ctx context.Context, arg domain.HostelBuilding) (*domain.HostelBuilding, error)
	CreateHostelRoom(ctx context.Context, arg domain.HostelRoom) (*domain.HostelRoom, error)
	AllocateRoom(ctx context.Context, arg domain.HostelAllocation) (*domain.HostelAllocation, error)
	VacateRoom(ctx context.Context, id, instituteID uuid.UUID) error
}
