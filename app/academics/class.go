package academics

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

func (h *Handler) CreateClassPeriod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var period domain.ClassPeriod
	if err := json.NewDecoder(r.Body).Decode(&period); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateClassPeriod(r.Context(), period)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create class period: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "class period created successfully", data)
}

func (h *Handler) ListClassPeriods(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	if instituteIDStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id is required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	data, err := h.service.ListClassPeriods(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch class periods: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "class periods fetched successfully", data)
}

// ========================= CREATE CLASS PERIOD =========================

// SERVICE
func (s *Service) CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error) {
	classPeriod, err := s.repo.CreateClassPeriod(ctx, arg)
	if err != nil {
		return nil, err
	}
	return classPeriod, nil
}

// REPOSITORY
func (r *Repository) CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error) {
	return nil, nil
}

// ========================= LIST CLASS PERIODS =========================

// SERVICE
func (s *Service) ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error) {
	classPeriods, err := s.repo.ListClassPeriods(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return classPeriods, nil
}

// REPOSITORY
func (r *Repository) ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error) {

	return nil, nil
}
