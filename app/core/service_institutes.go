package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateInstitute creates a new institute in the system
func (s *Service) CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	return s.repo.CreateInstitute(ctx, arg)
}

// DeleteInstitute removes an institute from the system
func (s *Service) DeleteInstitute(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteInstitute(ctx, id)
}

// GetInstituteById retrieves an institute by its ID
func (s *Service) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	return s.repo.GetInstituteById(ctx, id)
}

// GetInstituteByCode retrieves an institute by its unique code
func (s *Service) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	return s.repo.GetInstituteByCode(ctx, code)
}

// ListInstitutes retrieves all institutes
func (s *Service) ListInstitutes(ctx context.Context) ([]*domain.Institute, error) {
	return s.repo.ListInstitutes(ctx)
}

// UpdateInstitute updates an existing institute's information
func (s *Service) UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	return s.repo.UpdateInstitute(ctx, arg)
}
