package academics

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// ========================= SERVICE =========================

func (s *Service) CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error) {
	subject, err := s.repo.CreateSubject(ctx, arg)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (s *Service) ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error) {
	subjects, err := s.repo.ListSubjects(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

// ========================= REPOSITORY =========================

func (r *Repository) CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error) {
	return nil, nil
}

func (r *Repository) ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error) {
	return nil, nil
}
