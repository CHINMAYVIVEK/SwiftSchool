package fleet

import (
	"context"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

// =================================================================================
// HANDLERS
// =================================================================================

func (h *Handler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

// ========================= CREATE VEHICLE =========================

// SERVICE
func (s *Service) CreateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error) {
	return s.repo.CreateVehicle(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error) {
	return &arg, nil
}

// ========================= LIST VEHICLES =========================

// SERVICE
func (s *Service) ListRoutes(ctx context.Context, instituteID uuid.UUID) ([]*domain.Route, error) {
	return s.repo.ListRoutes(ctx, instituteID)
}

// REPOSITORY
func (r *Repository) ListRoutes(ctx context.Context, instituteID uuid.UUID) ([]*domain.Route, error) {
	return nil, nil
}

// ========================= UPDATE VEHICLE =========================

// SERVICE
func (s *Service) UpdateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error) {
	return s.repo.UpdateVehicle(ctx, arg)
}

// REPOSITORY
func (r *Repository) UpdateVehicle(ctx context.Context, arg domain.Vehicle) (*domain.Vehicle, error) {
	return &arg, nil
}

// ========================= DELETE VEHICLE =========================

// SERVICE
func (s *Service) DeleteVehicle(ctx context.Context, id, instituteID uuid.UUID) error {
	return s.repo.DeleteVehicle(ctx, id, instituteID)
}

// REPOSITORY
func (r *Repository) DeleteVehicle(ctx context.Context, id, instituteID uuid.UUID) error {
	return nil
}

func (s *Service) CreateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error) {
	return nil, nil
}

func (r *Repository) CreateFuelLog(ctx context.Context, arg domain.FuelLog) (*domain.FuelLog, error) {
	return nil, nil
}
