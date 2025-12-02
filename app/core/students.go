package core

import (
	"context"
	"swiftschool/internal/db"
)

func (s *Service) CreateStudent(ctx context.Context, arg db.CreateStudentParams) (db.CoreStudent, error) {
	coreStudent, err := s.repo.CreateStudent(ctx, arg)
	if err != nil {
		return coreStudent, err
	}
	return coreStudent, nil
}

func (s *Service) LinkStudentGuardian(ctx context.Context, arg db.LinkStudentGuardianParams) error {
	err := s.repo.LinkStudentGuardian(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteStudent(ctx context.Context, arg db.DeleteStudentParams) error {
	err := s.repo.DeleteStudent(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetStudentFullProfile(ctx context.Context, arg db.GetStudentFullProfileParams) (db.GetStudentFullProfileRow, error) {
	getStudentFullProfileRow, err := s.repo.GetStudentFullProfile(ctx, arg)
	if err != nil {
		return getStudentFullProfileRow, err
	}
	return getStudentFullProfileRow, nil
}
func (s *Service) ListStudentsByClass(ctx context.Context, arg db.ListStudentsByClassParams) ([]db.CoreStudent, error) {
	coreStudents, err := s.repo.ListStudentsByClass(ctx, arg)
	if err != nil {
		return coreStudents, nil
	}
	return coreStudents, nil
}
func (s *Service) SearchStudents(ctx context.Context, arg db.SearchStudentsParams) ([]db.CoreStudent, error) {
	coreStudents, err := s.repo.SearchStudents(ctx, arg)
	if err != nil {
		return coreStudents, nil
	}
	return coreStudents, nil
}
func (s *Service) UpdateStudent(ctx context.Context, arg db.UpdateStudentParams) (db.CoreStudent, error) {
	coreStudent, err := s.repo.UpdateStudent(ctx, arg)
	if err != nil {
		return coreStudent, err
	}
	return coreStudent, nil
}

func (r *Repository) CreateStudent(ctx context.Context, arg db.CreateStudentParams) (db.CoreStudent, error) {
	coreStudent, err := r.CreateStudent(ctx, arg)
	if err != nil {
		return coreStudent, err
	}
	return coreStudent, nil
}

func (r *Repository) LinkStudentGuardian(ctx context.Context, arg db.LinkStudentGuardianParams) error {
	err := r.LinkStudentGuardian(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) DeleteStudent(ctx context.Context, arg db.DeleteStudentParams) error {
	err := r.DeleteStudent(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetStudentFullProfile(ctx context.Context, arg db.GetStudentFullProfileParams) (db.GetStudentFullProfileRow, error) {
	getStudentFullProfileRow, err := r.GetStudentFullProfile(ctx, arg)
	if err != nil {
		return getStudentFullProfileRow, err
	}
	return getStudentFullProfileRow, nil
}
func (r *Repository) ListStudentsByClass(ctx context.Context, arg db.ListStudentsByClassParams) ([]db.CoreStudent, error) {
	coreStudents, err := r.ListStudentsByClass(ctx, arg)
	if err != nil {
		return coreStudents, nil
	}
	return coreStudents, nil
}
func (r *Repository) SearchStudents(ctx context.Context, arg db.SearchStudentsParams) ([]db.CoreStudent, error) {
	coreStudents, err := r.SearchStudents(ctx, arg)
	if err != nil {
		return coreStudents, nil
	}
	return coreStudents, nil
}
func (r *Repository) UpdateStudent(ctx context.Context, arg db.UpdateStudentParams) (db.CoreStudent, error) {
	coreStudent, err := r.UpdateStudent(ctx, arg)
	if err != nil {
		return coreStudent, err
	}
	return coreStudent, nil
}
