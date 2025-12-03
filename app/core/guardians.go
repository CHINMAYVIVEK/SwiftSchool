package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                 GUARDIAN METHODS                //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	return s.repo.CreateGuardian(ctx, arg)
}

func (r *Repository) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LINK GUARDIAN =========================
func (s *Service) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error {
	return s.repo.LinkStudentGuardian(ctx, studentID, guardianID)
}

func (r *Repository) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}
