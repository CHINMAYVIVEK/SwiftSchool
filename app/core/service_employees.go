package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateEmployee creates a new employee
func (s *Service) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	// Business logic: check code uniqueness? Repo handles constraint.
	return s.repo.CreateEmployee(ctx, arg)
}

// DeleteEmployee removes an employee
func (s *Service) DeleteEmployee(ctx context.Context, instituteID, id uuid.UUID) error {
	return s.repo.DeleteEmployee(ctx, instituteID, id)
}

// GetEmployeeById retrieves an employee by ID
func (s *Service) GetEmployeeById(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeById(ctx, instituteID, id)
}

// GetEmployeeFullProfile retrieves complete employee profile
func (s *Service) GetEmployeeFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeFullProfile(ctx, instituteID, id)
}

// ListEmployees retrieves all employees for an institute
func (s *Service) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	return s.repo.ListEmployees(ctx, instituteID)
}

// UpdateEmployee updates an existing employee
func (s *Service) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	return s.repo.UpdateEmployee(ctx, arg)
}
