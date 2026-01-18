package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateAcademicSession creates a new academic session
func (s *Service) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.CreateAcademicSession(ctx, arg)
}

// GetActiveSession retrieves the active session for an institute
func (s *Service) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error) {
	return s.repo.GetActiveSession(ctx, instituteID)
}

// ListAcademicSessions retrieves all sessions for an institute
func (s *Service) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error) {
	return s.repo.ListAcademicSessions(ctx, instituteID)
}

// UpdateAcademicSession updates a session
func (s *Service) UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.UpdateAcademicSession(ctx, arg)
}
