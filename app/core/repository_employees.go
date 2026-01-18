package core

import (
	"context"
	"fmt"

	"swiftschool/domain"
	"swiftschool/internal/db"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

// CreateEmployee inserts a new employee record into the database
func (r *Repository) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapDomainEmployeeToDBParams(arg)
	row, err := q.CreateEmployee(ctx, params)
	if err != nil {
		return nil, err
	}

	return mapper.MapDBEmployeeToDomain(row)
}

// DeleteEmployee removes an employee record from the database
func (r *Repository) DeleteEmployee(ctx context.Context, instituteID, id uuid.UUID) error {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Assuming UpdateEmployee (soft delete) or DeleteEmployee query exists.
	// If query missing, throw error.
	// Logic similar to students.
	// For now, assume it's missing or TODO.
	return fmt.Errorf("DeleteEmployee not implemented in queries")
}

// GetEmployeeById retrieves an employee by ID
func (r *Repository) GetEmployeeById(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	row, err := q.GetEmployeeById(ctx, db.GetEmployeeByIdParams{
		ID:          id,
		InstituteID: instituteID,
	})
	if err != nil {
		return nil, err
	}

	return mapper.MapDBEmployeeToDomain(row)
}

// GetEmployeeFullProfile retrieves complete employee profile
func (r *Repository) GetEmployeeFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	row, err := q.GetEmployeeFullProfile(ctx, db.GetEmployeeFullProfileParams{
		ID:          id,
		InstituteID: instituteID,
	})
	if err != nil {
		return nil, err
	}

	return mapper.MapEmployeeFullProfileRowToDomain(row), nil
}

// ListEmployees retrieves all employees for an institute
func (r *Repository) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// ListEmployees(ctx, instituteID)
	rows, err := q.ListEmployees(ctx, instituteID)
	if err != nil {
		return nil, err
	}

	var employees []*domain.Employee
	for _, row := range rows {
		e := mapper.MapDBListEmployeesRowToDomain(row)
		employees = append(employees, e)
	}

	return employees, nil
}

// UpdateEmployee updates an existing employee record
func (r *Repository) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapUpdateEmployeeParams(arg)
	row, err := q.UpdateEmployee(ctx, params)
	if err != nil {
		return nil, err
	}

	return mapper.MapDBEmployeeToDomain(row)
}
