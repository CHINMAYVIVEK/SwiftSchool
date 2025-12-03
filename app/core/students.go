package core

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                  STUDENT METHODS                //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	return s.repo.CreateStudent(ctx, arg)
}

func (r *Repository) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStudent(ctx, id)
}

func (r *Repository) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= GET FULL PROFILE =========================
func (s *Service) GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	return s.repo.GetStudentFullProfile(ctx, id)
}

func (r *Repository) GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST BY CLASS =========================
func (s *Service) ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]*domain.Student, error) {
	return s.repo.ListStudentsByClass(ctx, classID)
}

func (r *Repository) ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= SEARCH =========================
func (s *Service) SearchStudents(ctx context.Context, query string) ([]*domain.Student, error) {
	return s.repo.SearchStudents(ctx, query)
}

func (r *Repository) SearchStudents(ctx context.Context, query string) ([]*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	return s.repo.UpdateStudent(ctx, arg)
}

func (r *Repository) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}
