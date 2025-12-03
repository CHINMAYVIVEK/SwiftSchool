package core

import (
	"context"
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
func (h *Handler) CreateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data interface{}
	// Success response
	helper.NewSuccessResponse(w, http.StatusOK, "institute created successfully", data)
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
