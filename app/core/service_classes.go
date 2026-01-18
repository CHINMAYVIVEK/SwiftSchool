package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateClass creates a new class
func (s *Service) CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.CreateClass(ctx, arg)
}

// DeleteClass removes a class
func (s *Service) DeleteClass(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteClass(ctx, id)
}

// ListClasses retrieves all classes for an institute
func (s *Service) ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error) {
	return s.repo.ListClasses(ctx, instituteID)
}

// UpdateClass updates an existing class
func (s *Service) UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error) {
	return s.repo.UpdateClass(ctx, arg)
}
