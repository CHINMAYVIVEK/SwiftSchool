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
//                 GUARDIAN METHODS                //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE GUARDIAN ----------------
func (h *Handler) CreateGuardian(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var guardian domain.Guardian
	if err := json.NewDecoder(r.Body).Decode(&guardian); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateGuardian(r.Context(), guardian)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create guardian: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "guardian created successfully", data)
}

// ---------------- LINK STUDENT-GUARDIAN ----------------
func (h *Handler) LinkStudentGuardian(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	studentIDStr := r.URL.Query().Get("student_id")
	guardianIDStr := r.URL.Query().Get("guardian_id")

	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id")
		return
	}

	guardianID, err := uuid.Parse(guardianIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid guardian id")
		return
	}

	if err := h.service.LinkStudentGuardian(r.Context(), studentID, guardianID); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to link student and guardian: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student linked with guardian successfully", nil)
}

// ========================= CREATE =========================
func (s *Service) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	return s.repo.CreateGuardian(ctx, arg)
}

func (r *Repository) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LINK GUARDIAN =========================
func (s *Service) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error {
	return s.repo.LinkStudentGuardian(ctx, studentID, guardianID)
}

func (r *Repository) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}
