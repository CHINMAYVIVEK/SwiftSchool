package core

import (
	"context"
	"net/http"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                  CLASS METHODS                   //
//////////////////////////////////////////////////////

// ------------------------ HANDLER ------------------------
func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	// TODO: implement HTTP handler logic here
}

// ========================= CREATE =========================
func (s *Service) CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.CreateClass(ctx, arg)
}

func (r *Repository) CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteClass(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteClass(ctx, id)
}

func (r *Repository) DeleteClass(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= LIST =========================
func (s *Service) ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error) {
	return s.repo.ListClasses(ctx, instituteID)
}

func (r *Repository) ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.UpdateClass(ctx, arg)
}

func (r *Repository) UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	// TODO: implement DB logic here
	return nil, nil
}
