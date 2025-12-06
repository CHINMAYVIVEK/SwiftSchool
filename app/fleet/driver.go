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

func (h *Handler) CreateDriverProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

// ========================= CREATE DRIVER PROFILE =========================

// SERVICE
func (s *Service) CreateDriverProfile(ctx context.Context, arg domain.DriverProfile) (*domain.DriverProfile, error) {
	return s.repo.CreateDriverProfile(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateDriverProfile(ctx context.Context, arg domain.DriverProfile) (*domain.DriverProfile, error) {
	return &arg, nil
}
