package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateDepartment creates a new department
func (s *Service) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.CreateDepartment(ctx, arg)
}

// DeleteDepartment removes a department
func (s *Service) DeleteDepartment(ctx context.Context, instituteID, id uuid.UUID) error {
	return s.repo.DeleteDepartment(ctx, instituteID, id)
}

// ListDepartments retrieves all departments for an institute
func (s *Service) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	return s.repo.ListDepartments(ctx, instituteID)
}

// UpdateDepartment updates an existing department
func (s *Service) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.UpdateDepartment(ctx, arg)
}
