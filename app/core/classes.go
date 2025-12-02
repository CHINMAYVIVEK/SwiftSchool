package core

import (
	"context"
	"database/sql"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
)

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.Institute
	data, err := h.service.CreateClass(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create institute: "+err.Error())
		return
	}
	// Success response
	helper.NewSuccessResponse(w, http.StatusOK, "institute created successfully", data)
}

func (s *Service) CreateClass(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
	return s.repo.CreateClass(ctx, institute)
}

func (r *Repository) CreateClass(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
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
