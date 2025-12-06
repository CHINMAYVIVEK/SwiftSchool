package fleet

import (
	"context"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// =================================================================================
// HANDLERS
// =================================================================================

func (h *Handler) CreateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) ListRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) CreateRouteStop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) CreateTripLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

// ========================= CREATE ROUTE =========================

// SERVICE
func (s *Service) CreateRoute(ctx context.Context, arg domain.Route) (*domain.Route, error) {
	return s.repo.CreateRoute(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateRoute(ctx context.Context, arg domain.Route) (*domain.Route, error) {
	return &arg, nil
}
