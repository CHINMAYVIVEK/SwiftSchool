package geo

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
	// ========================= COUNTRY =========================
	CreateCountry(ctx context.Context, arg domain.Country) (*domain.Country, error)
	GetCountry(ctx context.Context, id uuid.UUID) (*domain.Country, error)
	ListCountries(ctx context.Context) ([]*domain.Country, error)
	UpdateCountry(ctx context.Context, arg domain.Country) (*domain.Country, error)
	DeleteCountry(ctx context.Context, id uuid.UUID) error

	// ========================= STATE =========================
	CreateState(ctx context.Context, arg domain.State) (*domain.State, error)
	ListStatesByCountry(ctx context.Context, countryID uuid.UUID) ([]*domain.State, error)
	UpdateState(ctx context.Context, arg domain.State) (*domain.State, error)
	DeleteState(ctx context.Context, id uuid.UUID) error

	// ========================= DISTRICT =========================
	CreateDistrict(ctx context.Context, arg domain.District) (*domain.District, error)
	ListDistrictsByState(ctx context.Context, stateID uuid.UUID) ([]*domain.District, error)
	UpdateDistrict(ctx context.Context, arg domain.District) (*domain.District, error)
	DeleteDistrict(ctx context.Context, id uuid.UUID) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= COUNTRY =========================
	CreateCountry(ctx context.Context, arg domain.Country) (*domain.Country, error)
	GetCountry(ctx context.Context, id uuid.UUID) (*domain.Country, error)
	ListCountries(ctx context.Context) ([]*domain.Country, error)
	UpdateCountry(ctx context.Context, arg domain.Country) (*domain.Country, error)
	DeleteCountry(ctx context.Context, id uuid.UUID) error

	// ========================= STATE =========================
	CreateState(ctx context.Context, arg domain.State) (*domain.State, error)
	ListStatesByCountry(ctx context.Context, countryID uuid.UUID) ([]*domain.State, error)
	UpdateState(ctx context.Context, arg domain.State) (*domain.State, error)
	DeleteState(ctx context.Context, id uuid.UUID) error

	// ========================= DISTRICT =========================
	CreateDistrict(ctx context.Context, arg domain.District) (*domain.District, error)
	ListDistrictsByState(ctx context.Context, stateID uuid.UUID) ([]*domain.District, error)
	UpdateDistrict(ctx context.Context, arg domain.District) (*domain.District, error)
	DeleteDistrict(ctx context.Context, id uuid.UUID) error
}
