package core

import (
	"context"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

func (s *Service) CreateDepartment(ctx context.Context, arg db.CreateDepartmentParams) (db.CoreDepartment, error) {
	coreDepartment, err := s.repo.CreateDepartment(ctx, arg)
	if err != nil {
		return coreDepartment, err
	}
	return coreDepartment, nil
}
func (s *Service) DeleteDepartment(ctx context.Context, arg db.DeleteDepartmentParams) error {
	err := s.repo.DeleteDepartment(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) ListDepartments(ctx context.Context, instituteID uuid.NullUUID) ([]db.CoreDepartment, error) {
	listDepartments, err := s.repo.ListDepartments(ctx, instituteID)
	if err != nil {
		return listDepartments, err
	}
	return listDepartments, nil
}
func (s *Service) UpdateDepartment(ctx context.Context, arg db.UpdateDepartmentParams) (db.CoreDepartment, error) {
	coreDepartments, err := s.repo.UpdateDepartment(ctx, arg)
	if err != nil {
		return coreDepartments, err
	}
	return coreDepartments, nil
}

func (r *Repository) CreateDepartment(ctx context.Context, arg db.CreateDepartmentParams) (db.CoreDepartment, error) {
	coreDepartment, err := r.CreateDepartment(ctx, arg)
	if err != nil {
		return coreDepartment, err
	}
	return coreDepartment, nil
}
func (r *Repository) DeleteDepartment(ctx context.Context, arg db.DeleteDepartmentParams) error {
	err := r.DeleteDepartment(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ListDepartments(ctx context.Context, instituteID uuid.NullUUID) ([]db.CoreDepartment, error) {
	listDepartments, err := r.ListDepartments(ctx, instituteID)
	if err != nil {
		return listDepartments, err
	}
	return listDepartments, nil
}
func (r *Repository) UpdateDepartment(ctx context.Context, arg db.UpdateDepartmentParams) (db.CoreDepartment, error) {
	coreDepartments, err := r.UpdateDepartment(ctx, arg)
	if err != nil {
		return coreDepartments, err
	}
	return coreDepartments, nil
}
