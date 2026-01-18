package fleet

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
	CreateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error)
	ListVehicles(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vehicle, error)
	UpdateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error)
	DeleteVehicle(ctx context.Context, id, instituteID uuid.UUID) error

	CreateDriverProfile(ctx context.Context, arg domain.DriverProfile) (*domain.DriverProfile, error)

	CreateRoute(ctx context.Context, arg domain.Route) (*domain.Route, error)
	ListRoutes(ctx context.Context, instituteID uuid.UUID) ([]*domain.Route, error)

	CreateRouteStop(ctx context.Context, arg domain.RouteStop) (*domain.RouteStop, error)
	ListRouteStops(ctx context.Context, instituteID uuid.UUID) ([]*domain.RouteStop, error)
	UpdateRouteStop(ctx context.Context, arg domain.RouteStop) (*domain.RouteStop, error)
	DeleteRouteStop(ctx context.Context, id, instituteID uuid.UUID) error
	ListTripLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.TripLog, error)
	CreateTripLog(ctx context.Context, arg domain.TripLog) (*domain.TripLog, error)
	UpdateTripLog(ctx context.Context, arg domain.TripLog) (*domain.TripLog, error)
	// DeleteTripLog(ctx context.Context, id, instituteID uuid.UUID) error

	CreateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error)
	ListFuelLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.FuelLog, error)
	UpdateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error)
	// DeleteFuelLog(ctx context.Context, id, instituteID uuid.UUID) error

	CreateMaintenanceLog(ctx context.Context, arg domain.MaintenanceLog) (*domain.MaintenanceLog, error)
	ListMaintenanceLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.MaintenanceLog, error)
	UpdateMaintenanceLog(ctx context.Context, arg domain.MaintenanceLog) (*domain.MaintenanceLog, error)
	// DeleteMaintenanceLog(ctx context.Context, id, instituteID uuid.UUID) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error)
	ListVehicles(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vehicle, error)
	UpdateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error)
	DeleteVehicle(ctx context.Context, id, instituteID uuid.UUID) error

	CreateDriverProfile(ctx context.Context, arg domain.DriverProfile) (*domain.DriverProfile, error)

	CreateRoute(ctx context.Context, arg domain.Route) (*domain.Route, error)
	ListRoutes(ctx context.Context, instituteID uuid.UUID) ([]*domain.Route, error)

	CreateRouteStop(ctx context.Context, arg domain.RouteStop) (*domain.RouteStop, error)
	ListRouteStops(ctx context.Context, instituteID uuid.UUID) ([]*domain.RouteStop, error)
	UpdateRouteStop(ctx context.Context, arg domain.RouteStop) (*domain.RouteStop, error)
	DeleteRouteStop(ctx context.Context, id, instituteID uuid.UUID) error
	ListTripLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.TripLog, error)
	CreateTripLog(ctx context.Context, arg domain.TripLog) (*domain.TripLog, error)
	UpdateTripLog(ctx context.Context, arg domain.TripLog) (*domain.TripLog, error)
	// DeleteTripLog(ctx context.Context, id, instituteID uuid.UUID) error

	CreateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error)
	ListFuelLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.FuelLog, error)
	UpdateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error)
	// DeleteFuelLog(ctx context.Context, id, instituteID uuid.UUID) error

	CreateMaintenanceLog(ctx context.Context, arg domain.MaintenanceLog) (*domain.MaintenanceLog, error)
	ListMaintenanceLogs(ctx context.Context, instituteID uuid.UUID) ([]*domain.MaintenanceLog, error)
	UpdateMaintenanceLog(ctx context.Context, arg domain.MaintenanceLog) (*domain.MaintenanceLog, error)
	// DeleteMaintenanceLog(ctx context.Context, id, instituteID uuid.UUID) error
}
