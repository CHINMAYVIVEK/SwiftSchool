package auth

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
	Login(ctx context.Context, identifier string, userType string) error
	CreateUser(ctx context.Context, arg domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, id uuid.UUID, passwordHash string) error
	UpdateUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error
	ListUsersByRole(ctx context.Context, instituteID uuid.UUID, roleType domain.UserRole) ([]*domain.User, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	Login(ctx context.Context, identifier string, userType string) error
	VerifyOTP(ctx context.Context, identifier string, userType string, otp string) (*domain.User, error)
	CreateUser(ctx context.Context, arg domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, id uuid.UUID, passwordHash string) error
	UpdateUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error
	ListUsersByRole(ctx context.Context, instituteID uuid.UUID, roleType domain.UserRole) ([]*domain.User, error)
}
