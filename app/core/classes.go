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
//                  CLASS METHODS                   //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE CLASS ----------------
func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var class domain.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateClass(r.Context(), class)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "class created successfully", data)
}

// ---------------- DELETE CLASS ----------------
func (h *Handler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid class id")
		return
	}

	if err := h.service.DeleteClass(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "class deleted successfully", nil)
}

// ---------------- LIST CLASSES ----------------
func (h *Handler) ListClasses(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListClasses(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list classes: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "classes retrieved successfully", data)
}

// ---------------- UPDATE CLASS ----------------
func (h *Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var class domain.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateClass(r.Context(), class)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "class updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.CreateClass(ctx, arg)
}

func (r *Repository) CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteClass(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteClass(ctx, id)
}

func (r *Repository) DeleteClass(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= LIST =========================
func (s *Service) ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error) {
	return s.repo.ListClasses(ctx, instituteID)
}

func (r *Repository) ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.UpdateClass(ctx, arg)
}

func (r *Repository) UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}
