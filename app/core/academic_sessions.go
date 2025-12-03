package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

func (s *Service) CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error) {

	coreAcademicSession, err := s.repo.CreateAcademicSession(ctx, arg)
	if err != nil {
		return coreAcademicSession, err
	}
	return coreAcademicSession, nil
}

func (s *Service) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]db.CoreAcademicSession, error) {
	coreAcademicSessions, err := s.repo.ListAcademicSessions(ctx, instituteID)
	if err != nil {
		return coreAcademicSessions, err
	}
	return coreAcademicSessions, nil
}
func (s *Service) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (db.CoreAcademicSession, error) {
	coreAdmissionSessoin, err := s.repo.GetActiveSession(ctx, instituteID)
	if err != nil {
		return coreAdmissionSessoin, err
	}
	return coreAdmissionSessoin, nil
}
func (s *Service) UpdateAcademicSession(ctx context.Context, arg db.UpdateAcademicSessionParams) (db.CoreAcademicSession, error) {
	coreAcademicSession, err := s.repo.UpdateAcademicSession(ctx, arg)
	if err != nil {
		return coreAcademicSession, err
	}
	return coreAcademicSession, nil
}

func (r *Repository) CreateAcademicSession(context.Context, domain.AcademicSession) (*domain.AcademicSession, error) {

	return nil, nil
}
func (r *Repository) ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]db.CoreAcademicSession, error) {
	coreAcademicSessions, err := r.ListAcademicSessions(ctx, instituteID)
	if err != nil {
		return coreAcademicSessions, err
	}
	return coreAcademicSessions, nil
}
func (r *Repository) GetActiveSession(ctx context.Context, instituteID uuid.UUID) (db.CoreAcademicSession, error) {
	coreAdmissionSessoin, err := r.GetActiveSession(ctx, instituteID)
	if err != nil {
		return coreAdmissionSessoin, err
	}
	return coreAdmissionSessoin, nil
}
func (r *Repository) UpdateAcademicSession(ctx context.Context, arg db.UpdateAcademicSessionParams) (db.CoreAcademicSession, error) {
	coreAcademicSession, err := r.UpdateAcademicSession(ctx, arg)
	if err != nil {
		return coreAcademicSession, err
	}
	return coreAcademicSession, nil
}
