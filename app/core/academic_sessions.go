package core

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//             ACADEMIC SESSION METHODS            //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE ACADEMIC SESSION ----------------
func (h *Handler) CreateAcademicSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var session domain.AcademicSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateAcademicSession(r.Context(), session)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create academic session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "academic session created successfully", data)
}

// ---------------- LIST ACADEMIC SESSIONS ----------------
func (h *Handler) ListAcademicSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id")
		return
	}

	data, err := h.service.ListAcademicSessions(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list academic sessions: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "academic sessions retrieved successfully", data)
}

// ---------------- GET ACTIVE SESSION ----------------
func (h *Handler) GetActiveSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id")
		return
	}

	data, err := h.service.GetActiveSession(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get active academic session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "active academic session retrieved successfully", data)
}

// ---------------- UPDATE ACADEMIC SESSION ----------------
func (h *Handler) UpdateAcademicSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var session domain.AcademicSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateAcademicSession(r.Context(), session)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update academic session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "academic session updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.CreateAcademicSession(ctx, arg)
}

func (r *Repository) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST =========================
func (s *Service) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error) {
	return s.repo.ListAcademicSessions(ctx, instituteID)
}

func (r *Repository) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= GET ACTIVE =========================
func (s *Service) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error) {
	return s.repo.GetActiveSession(ctx, instituteID)
}

func (r *Repository) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.UpdateAcademicSession(ctx, arg)
}

func (r *Repository) UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}
