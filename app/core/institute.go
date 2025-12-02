package core

import (
	"context"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

func (h *Handler) CreateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data interface{}
	// Success response
	helper.NewSuccessResponse(w, http.StatusOK, "institute created successfully", data)
}

func (s *Service) CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	coreInstitute, err := s.repo.CreateInstitute(ctx, arg)
	if err != nil {
		return nil, err
	}
	return coreInstitute, nil
}

func (s *Service) DeleteInstitute(ctx context.Context, arg db.DeleteInstituteParams) error {
	err := s.repo.DeleteInstitute(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	coreInstitute, err := s.repo.GetInstituteByCode(ctx, code)
	if err != nil {
		return coreInstitute, err
	}
	return coreInstitute, nil
}
func (s *Service) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	coreInstitute, err := s.repo.GetInstituteById(ctx, id)
	if err != nil {
		return coreInstitute, err
	}
	return coreInstitute, nil
}
func (s *Service) ListInstitutes(ctx context.Context) ([]domain.Institute, error) {
	coreInstitutes, err := s.repo.ListInstitutes(ctx)
	if err != nil {
		return coreInstitutes, err
	}
	return coreInstitutes, nil
}
func (s *Service) UpdateInstitute(ctx context.Context, arg db.UpdateInstituteParams) (*domain.Institute, error) {
	coreInstitute, err := s.repo.UpdateInstitute(ctx, arg)
	if err != nil {
		return coreInstitute, err
	}
	return coreInstitute, nil
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
		CurrencyCode: helper.ToNullString(arg.CurrencyCode),
		LogoUrl:      helper.ToNullString(arg.LogoURL),
		Website:      helper.ToNullString(arg.Website),
		IsActive:     helper.ToNullBool(arg.IsActive),
		CreatedBy:    helper.ToNullUUID(arg.CreatedBy),
	}

	coreInstitute, err := q.CreateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	base := domain.BaseUUIDModel{
		ID: coreInstitute.ID,
	}
	institute := domain.Institute{
		BaseUUIDModel: base,
	}

	return &institute, nil
}

func (r *Repository) DeleteInstitute(ctx context.Context, arg db.DeleteInstituteParams) error {

	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return err
	}
	params := db.DeleteInstituteParams{
		ID: uuid.Nil,
		UpdatedBy: uuid.NullUUID{
			UUID:  uuid.Nil,
			Valid: false,
		},
	}

	err = q.DeleteInstitute(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	coreInstitute, err := q.GetInstituteByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	base := domain.BaseUUIDModel{
		ID: coreInstitute.ID,
	}
	institute := domain.Institute{
		BaseUUIDModel: base,
	}

	return &institute, nil
}
func (r *Repository) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	coreInstitute, err := r.GetInstituteById(ctx, id)
	if err != nil {
		return nil, err
	}
	base := domain.BaseUUIDModel{
		ID: coreInstitute.ID,
	}
	institute := domain.Institute{
		BaseUUIDModel: base,
	}

	return &institute, nil
}
func (r *Repository) ListInstitutes(ctx context.Context) ([]domain.Institute, error) {
	coreInstitutes, err := r.ListInstitutes(ctx)
	if err != nil {
		return coreInstitutes, err
	}
	return coreInstitutes, nil
}
func (r *Repository) UpdateInstitute(ctx context.Context, arg db.UpdateInstituteParams) (*domain.Institute, error) {
	coreInstitute, err := r.UpdateInstitute(ctx, arg)
	if err != nil {
		return coreInstitute, err
	}
	return coreInstitute, nil
}
