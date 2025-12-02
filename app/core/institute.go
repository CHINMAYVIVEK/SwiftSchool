package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
	"time"

	"github.com/google/uuid"
)

func (h *CoreHandler) InstitutesRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.Institute

	// Parse incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "name is required")
		return
	}
	if req.Code == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "code is required")
		return
	}
	if req.CreatedBy == nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "created_by is required (uuid)")
		return
	}

	// Auto-fill generated fields
	now := time.Now()
	req.ID = uuid.New()
	req.CreatedAt = now
	req.UpdatedAt = now
	req.UpdatedBy = req.CreatedBy
	req.DeletedAt = nil
	req.IsActive = true

	// Call service
	data, err := h.service.InstitutesRegistration(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create institute: "+err.Error())
		return
	}

	// Success response
	helper.NewSuccessResponse(w, http.StatusOK, "institute created successfully", data)
}

func (r *CoreRepository) InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// Map domain → SQLC params
	params := db.CreateInstituteParams{
		Name:         institute.Name,
		Code:         institute.Code,
		CurrencyCode: sql.NullString{String: helper.ToStr(institute.CurrencyCode), Valid: institute.CurrencyCode != nil},
		LogoUrl:      sql.NullString{String: helper.ToStr(institute.LogoURL), Valid: institute.LogoURL != nil},
		Website:      sql.NullString{String: helper.ToStr(institute.Website), Valid: institute.Website != nil},
		CreatedBy:    helper.ToNullUUID(institute.CreatedBy),
		CreatedAt:    helper.ToNullTime(&institute.CreatedAt),
		IsActive:     helper.ToNullBool(institute.IsActive),
	}

	coreInstitute, err := q.CreateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	base := domain.BaseUUIDModel{
		ID: coreInstitute.ID,
	}

	return &base, nil
}

func (s *CoreService) InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
	// Business validations or vault logic can be added here

	data, err := s.repo.InstitutesRegistration(ctx, institute)
	if err != nil {
		return nil, err
	}
	return data, nil
}
