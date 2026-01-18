package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateGuardian creates a new guardian
func (s *Service) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	return s.repo.CreateGuardian(ctx, arg)
}

// LinkStudentGuardian links a student to a guardian
func (s *Service) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID, relationship string, isPrimary bool) error {
	return s.repo.LinkStudentGuardian(ctx, studentID, guardianID, relationship, isPrimary)
}
