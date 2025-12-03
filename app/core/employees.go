package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                 EMPLOYEE METHODS                //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	return s.repo.CreateEmployee(ctx, arg)
}

func (r *Repository) CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (r *Repository) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= GET BY ID =========================
func (s *Service) GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeById(ctx, id)
}

func (r *Repository) GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= GET FULL PROFILE =========================
func (s *Service) GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetEmployeeFullProfile(ctx, id)
}

func (r *Repository) GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST =========================
func (s *Service) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	return s.repo.ListEmployees(ctx, instituteID)
}

func (r *Repository) ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	return s.repo.UpdateEmployee(ctx, arg)
}

func (r *Repository) UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error) {
	// TODO: implement DB logic here
	return nil, nil
}
