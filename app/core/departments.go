package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                 DEPARTMENT METHODS              //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.CreateDepartment(ctx, arg)
}

func (r *Repository) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDepartment(ctx, id)
}

func (r *Repository) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= LIST =========================
func (s *Service) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	return s.repo.ListDepartments(ctx, instituteID)
}

func (r *Repository) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	return s.repo.UpdateDepartment(ctx, arg)
}

func (r *Repository) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	// TODO: implement DB logic here
	return nil, nil
}
