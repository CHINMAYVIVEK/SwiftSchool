package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateStudent creates a new student in the system
func (s *Service) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	// CreateStudent in repo takes domain.Student which already has InstituteID
	return s.repo.CreateStudent(ctx, arg)
}

// DeleteStudent removes a student from the system
func (s *Service) DeleteStudent(ctx context.Context, instituteID, id uuid.UUID) error {
	return s.repo.DeleteStudent(ctx, instituteID, id)
}

// GetStudentFullProfile retrieves complete student information including related data
func (s *Service) GetStudentFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Student, error) {
	return s.repo.GetStudentFullProfile(ctx, instituteID, id)
}

// ListStudentsByClass retrieves all students in a specific class
func (s *Service) ListStudentsByClass(ctx context.Context, instituteID, classID uuid.UUID) ([]*domain.Student, error) {
	return s.repo.ListStudentsByClass(ctx, instituteID, classID)
}

// SearchStudents searches for students by name or admission number
func (s *Service) SearchStudents(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Student, error) {
	return s.repo.SearchStudents(ctx, instituteID, query)
}

// UpdateStudent updates an existing student's information
func (s *Service) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	return s.repo.UpdateStudent(ctx, arg)
}
