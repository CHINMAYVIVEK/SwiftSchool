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
//                 EMPLOYEE METHODS                //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE EMPLOYEE ----------------
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var employee domain.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateEmployee(r.Context(), employee)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "employee created successfully", data)
}

// ---------------- DELETE EMPLOYEE ----------------
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	if err := h.service.DeleteEmployee(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee deleted successfully", nil)
}

// ---------------- GET EMPLOYEE BY ID ----------------
func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	data, err := h.service.GetEmployeeById(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee retrieved successfully", data)
}

// ---------------- GET EMPLOYEE FULL PROFILE ----------------
func (h *Handler) GetEmployeeFullProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	data, err := h.service.GetEmployeeFullProfile(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get full profile: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee full profile retrieved successfully", data)
}

// ---------------- LIST EMPLOYEES ----------------
func (h *Handler) ListEmployees(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListEmployees(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list employees: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employees retrieved successfully", data)
}

// ---------------- UPDATE EMPLOYEE ----------------
func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var employee domain.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateEmployee(r.Context(), employee)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	return s.repo.CreateEmployee(ctx, arg)
}

func (r *Repository) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (r *Repository) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= GET BY ID =========================
func (s *Service) GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeById(ctx, id)
}

func (r *Repository) GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= GET FULL PROFILE =========================
func (s *Service) GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeFullProfile(ctx, id)
}

func (r *Repository) GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST =========================
func (s *Service) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	return s.repo.ListEmployees(ctx, instituteID)
}

func (r *Repository) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	return s.repo.UpdateEmployee(ctx, arg)
}

func (r *Repository) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}
