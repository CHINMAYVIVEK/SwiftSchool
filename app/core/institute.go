package core

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                 INSTITUTE METHODS               //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE INSTITUTE ----------------
func (h *Handler) CreateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var institute domain.Institute
	if err := json.NewDecoder(r.Body).Decode(&institute); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateInstitute(r.Context(), institute)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "institute creation failed: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "institute created successfully", data)
}

// ---------------- DELETE INSTITUTE ----------------
func (h *Handler) DeleteInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id")
		return
	}

	if err := h.service.DeleteInstitute(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute deleted successfully", nil)
}

// ---------------- GET INSTITUTE BY ID ----------------
func (h *Handler) GetInstituteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id")
		return
	}

	data, err := h.service.GetInstituteById(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute retrieved successfully", data)
}

// ---------------- GET INSTITUTE BY CODE ----------------
func (h *Handler) GetInstituteByCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute code is required")
		return
	}

	data, err := h.service.GetInstituteByCode(r.Context(), code)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute retrieved successfully", data)
}

// ---------------- LIST INSTITUTES ----------------
func (h *Handler) ListInstitutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	data, err := h.service.ListInstitutes(r.Context())
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list institutes: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institutes retrieved successfully", data)
}

// ---------------- UPDATE INSTITUTE ----------------
func (h *Handler) UpdateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var institute domain.Institute
	if err := json.NewDecoder(r.Body).Decode(&institute); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateInstitute(r.Context(), institute)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "institute update failed: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	return s.repo.CreateInstitute(ctx, arg)
}

func (r *Repository) CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := db.CreateInstituteParams{
		Name:         arg.Name,
		Code:         arg.Code,
		CurrencyCode: helper.PtrToNullString(arg.CurrencyCode),
		LogoUrl:      helper.PtrToNullString(arg.LogoURL),
		Website:      helper.PtrToNullString(arg.Website),
		IsActive:     helper.BoolToNullBool(arg.IsActive),
		CreatedBy:    helper.PtrToNullUUID(arg.CreatedBy),
	}

	coreInstitute, err := q.CreateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	institute := domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{ID: coreInstitute.ID},
	}

	return &institute, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteInstitute(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteInstitute(ctx, id)
}

func (r *Repository) DeleteInstitute(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return err
	}

	// Adjust params as needed
	params := db.DeleteInstituteParams{
		ID: id,
	}

	return q.DeleteInstitute(ctx, params)
}

// ========================= GET BY CODE =========================
func (s *Service) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	return s.repo.GetInstituteByCode(ctx, code)
}

func (r *Repository) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	coreInstitute, err := q.GetInstituteByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	institute := domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{ID: coreInstitute.ID},
	}

	return &institute, nil
}

// ========================= GET BY ID =========================
func (s *Service) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	return s.repo.GetInstituteById(ctx, id)
}

func (r *Repository) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	coreInstitute, err := q.GetInstituteById(ctx, id)
	if err != nil {
		return nil, err
	}

	institute := domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{ID: coreInstitute.ID},
	}

	return &institute, nil
}

// ========================= LIST =========================
func (s *Service) ListInstitutes(ctx context.Context) ([]*domain.Institute, error) {
	return s.repo.ListInstitutes(ctx)
}

func (r *Repository) ListInstitutes(ctx context.Context) ([]*domain.Institute, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	coreInstitutes, err := q.ListInstitutes(ctx)
	if err != nil {
		return nil, err
	}

	institutes := make([]*domain.Institute, len(coreInstitutes))
	for i, ci := range coreInstitutes {
		institutes[i] = &domain.Institute{
			BaseUUIDModel: domain.BaseUUIDModel{ID: ci.ID},
		}
	}

	return institutes, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	return s.repo.UpdateInstitute(ctx, arg)
}

func (r *Repository) UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}
	params := db.UpdateInstituteParams{}
	coreInstitute, err := q.UpdateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	institute := domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{ID: coreInstitute.ID},
	}

	return &institute, nil
}
