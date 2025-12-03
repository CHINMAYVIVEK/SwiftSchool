package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//             ACADEMIC SESSION METHODS            //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.CreateAcademicSession(ctx, arg)
}

func (r *Repository) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST =========================
func (s *Service) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error) {
	return s.repo.ListAcademicSessions(ctx, instituteID)
}

func (r *Repository) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= GET ACTIVE =========================
func (s *Service) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error) {
	return s.repo.GetActiveSession(ctx, instituteID)
}

func (r *Repository) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	return s.repo.UpdateAcademicSession(ctx, arg)
}

func (r *Repository) UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {
	// TODO: implement DB logic here
	return nil, nil
}
