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

// ========================= CREATE SUBJECT =========================
func (h *Handler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var subject domain.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateSubject(r.Context(), subject)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create subject: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "subject created successfully", data)
}

// ========================= LIST SUBJECTS =========================
func (h *Handler) ListSubjects(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListSubjects(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch subjects: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "subjects fetched successfully", data)
}

// ========================= SERVICE + REPO =========================

// SERVICE
func (s *Service) CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error) {
	subject, err := s.repo.CreateSubject(ctx, arg)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

// REPOSITORY
func (r *Repository) CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error) {

	return nil, nil
}

// SERVICE
func (s *Service) ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error) {
	subjects, err := s.repo.ListSubjects(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

// REPOSITORY
func (r *Repository) ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error) {

	return nil, nil
}
