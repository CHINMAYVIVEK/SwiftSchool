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
//                 DEPARTMENT METHODS              //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE DEPARTMENT ----------------
func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var department domain.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateDepartment(r.Context(), department)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "department created successfully", data)
}

// ---------------- DELETE DEPARTMENT ----------------
func (h *Handler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid department id")
		return
	}

	if err := h.service.DeleteDepartment(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "department deleted successfully", nil)
}

// ---------------- LIST DEPARTMENTS ----------------
func (h *Handler) ListDepartments(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListDepartments(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list departments: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "departments retrieved successfully", data)
}

// ---------------- UPDATE DEPARTMENT ----------------
func (h *Handler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var department domain.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateDepartment(r.Context(), department)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "department updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.CreateDepartment(ctx, arg)
}

func (r *Repository) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDepartment(ctx, id)
}

func (r *Repository) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= LIST =========================
func (s *Service) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	return s.repo.ListDepartments(ctx, instituteID)
}

func (r *Repository) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.UpdateDepartment(ctx, arg)
}

func (r *Repository) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}
