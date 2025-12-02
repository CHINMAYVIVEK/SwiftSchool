package core

import (
	"context"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

func (s *Service) CreateEmployee(ctx context.Context, arg db.CreateEmployeeParams) (db.CoreEmployee, error) {
	coreEmployee, err := s.repo.CreateEmployee(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}
func (s *Service) DeleteEmployee(ctx context.Context, arg db.DeleteEmployeeParams) error {
	err := s.repo.DeleteEmployee(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetEmployeeById(ctx context.Context, arg db.GetEmployeeByIdParams) (db.CoreEmployee, error) {
	coreEmployee, err := s.repo.GetEmployeeById(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}
func (s *Service) GetEmployeeFullProfile(ctx context.Context, arg db.GetEmployeeFullProfileParams) (db.GetEmployeeFullProfileRow, error) {
	getEmployeeFullProfileRow, err := s.repo.GetEmployeeFullProfile(ctx, arg)
	if err != nil {
		return getEmployeeFullProfileRow, err
	}
	return getEmployeeFullProfileRow, nil
}
func (s *Service) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]db.ListEmployeesRow, error) {
	listEmployeesRows, err := s.repo.ListEmployees(ctx, instituteID)
	if err != nil {
		return listEmployeesRows, err
	}
	return listEmployeesRows, nil
}
func (s *Service) UpdateEmployee(ctx context.Context, arg db.UpdateEmployeeParams) (db.CoreEmployee, error) {
	coreEmployee, err := s.repo.UpdateEmployee(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}

func (r *Repository) CreateEmployee(ctx context.Context, arg db.CreateEmployeeParams) (db.CoreEmployee, error) {
	coreEmployee, err := r.CreateEmployee(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}
func (r *Repository) DeleteEmployee(ctx context.Context, arg db.DeleteEmployeeParams) error {
	err := r.DeleteEmployee(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetEmployeeById(ctx context.Context, arg db.GetEmployeeByIdParams) (db.CoreEmployee, error) {
	coreEmployee, err := r.GetEmployeeById(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}
func (r *Repository) GetEmployeeFullProfile(ctx context.Context, arg db.GetEmployeeFullProfileParams) (db.GetEmployeeFullProfileRow, error) {
	getEmployeeFullProfileRow, err := r.GetEmployeeFullProfile(ctx, arg)
	if err != nil {
		return getEmployeeFullProfileRow, err
	}
	return getEmployeeFullProfileRow, nil
}
func (r *Repository) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]db.ListEmployeesRow, error) {
	listEmployeesRows, err := r.ListEmployees(ctx, instituteID)
	if err != nil {
		return listEmployeesRows, err
	}
	return listEmployeesRows, nil
}
func (r *Repository) UpdateEmployee(ctx context.Context, arg db.UpdateEmployeeParams) (db.CoreEmployee, error) {
	coreEmployee, err := r.UpdateEmployee(ctx, arg)
	if err != nil {
		return coreEmployee, err
	}
	return coreEmployee, nil
}
