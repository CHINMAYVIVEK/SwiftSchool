package core

import (
	"context"
	"net/http"
	"swiftschool/internal/db"
)

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {}

func (s *Service) CreateClass(ctx context.Context, arg db.CreateClassParams) (db.CoreClass, error) {
	coreClass, err := s.repo.CreateClass(ctx, arg)
	if err != nil {
		return coreClass, err
	}
	return coreClass, nil
}
func (s *Service) DeleteClass(ctx context.Context, arg db.DeleteClassParams) error {
	err := s.repo.DeleteClass(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) ListClasses(ctx context.Context, arg db.ListClassesParams) ([]db.ListClassesRow, error) {
	listClassesRows, err := s.repo.ListClasses(ctx, arg)
	if err != nil {
		return listClassesRows, err
	}
	return listClassesRows, nil
}
func (s *Service) UpdateClass(ctx context.Context, arg db.UpdateClassParams) (db.CoreClass, error) {
	coreClass, err := s.repo.UpdateClass(ctx, arg)
	if err != nil {
		return coreClass, err
	}
	return coreClass, nil
}

func (r *Repository) CreateClass(ctx context.Context, arg db.CreateClassParams) (db.CoreClass, error) {
	coreClass, err := r.CreateClass(ctx, arg)
	if err != nil {
		return coreClass, err
	}
	return coreClass, nil
}
func (r *Repository) DeleteClass(ctx context.Context, arg db.DeleteClassParams) error {
	err := r.DeleteClass(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ListClasses(ctx context.Context, arg db.ListClassesParams) ([]db.ListClassesRow, error) {
	listClassesRows, err := r.ListClasses(ctx, arg)
	if err != nil {
		return listClassesRows, err
	}
	return listClassesRows, nil
}
func (r *Repository) UpdateClass(ctx context.Context, arg db.UpdateClassParams) (db.CoreClass, error) {
	coreClass, err := r.UpdateClass(ctx, arg)
	if err != nil {
		return coreClass, err
	}
	return coreClass, nil
}
