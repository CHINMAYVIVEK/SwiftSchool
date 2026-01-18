package hr

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
	CreateLeaveType(ctx context.Context, arg domain.LeaveType) (*domain.LeaveType, error)
	CreateLeaveApplication(ctx context.Context, arg domain.LeaveApplication) (*domain.LeaveApplication, error)
	ListLeaveApplications(ctx context.Context, instituteID uuid.UUID, status domain.LeaveStatus) ([]*domain.LeaveApplication, error)
	ApproveLeave(ctx context.Context, id, instituteID, approverID uuid.UUID, status domain.LeaveStatus) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateLeaveType(ctx context.Context, arg domain.LeaveType) (*domain.LeaveType, error)
	CreateLeaveApplication(ctx context.Context, arg domain.LeaveApplication) (*domain.LeaveApplication, error)
	ListLeaveApplications(ctx context.Context, instituteID uuid.UUID, status domain.LeaveStatus) ([]*domain.LeaveApplication, error)
	ApproveLeave(ctx context.Context, id, instituteID, approverID uuid.UUID, status domain.LeaveStatus) error
}
