package academics

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// ========================= SERVICE =========================

func (s *Service) CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error) {
	classPeriod, err := s.repo.CreateClassPeriod(ctx, arg)
	if err != nil {
		return nil, err
	}
	return classPeriod, nil
}

func (s *Service) ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error) {
	classPeriods, err := s.repo.ListClassPeriods(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return classPeriods, nil
}

// ========================= REPOSITORY =========================

func (r *Repository) CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error) {
	return nil, nil
}

func (r *Repository) ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error) {
	return nil, nil
}
