package core

import (
	"context"
	"fmt"

	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

// CreateStudent inserts a new student record into the database
func (r *Repository) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapStudentDomainToParams(arg)

	row, err := q.CreateStudent(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapStudentRowToDomain(row)
	return &out, nil
}

// DeleteStudent removes a student record from the database (soft delete)
func (r *Repository) DeleteStudent(ctx context.Context, instituteID, id uuid.UUID) error {
	// TODO: Add DeleteStudent (Soft Delete) to queries.sql
	// For now, if we assume soft deletes are automatic via Update, we implemented Update.
	// But queries.sql is missing DeleteStudent.
	return fmt.Errorf("DeleteStudent logic pending update to queries")
}

// GetStudentFullProfile retrieves complete student information from the database
func (r *Repository) GetStudentFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Student, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := db.GetStudentFullProfileParams{
		ID:          id,
		InstituteID: instituteID,
	}

	row, err := q.GetStudentFullProfile(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapStudentFullProfileRowToDomain(row)
	return &out, nil
}

// ListStudentsByClass retrieves all students in a specific class
func (r *Repository) ListStudentsByClass(ctx context.Context, instituteID, classID uuid.UUID) ([]*domain.Student, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := db.ListStudentsByClassParams{
		InstituteID:    instituteID,
		CurrentClassID: helper.ToNullUUID(classID),
	}

	rows, err := q.ListStudentsByClass(ctx, params)
	if err != nil {
		return nil, err
	}

	var students []*domain.Student
	for _, row := range rows {
		s := mapper.MapStudentRowToDomain(row)
		students = append(students, &s)
	}

	return students, nil
}

// SearchStudents searches for students
func (r *Repository) SearchStudents(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Student, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// SQLC query: SearchStudents(instituteID, admissionNo)
	// Currently regex is not supported in the generated query comment "ILIKE $2".
	// We need to pass the pattern e.g. "%query%"
	params := db.SearchStudentsParams{
		InstituteID: instituteID,
		AdmissionNo: "%" + query + "%", // Assuming simple containment search
	}

	rows, err := q.SearchStudents(ctx, params)
	if err != nil {
		return nil, err
	}

	var students []*domain.Student
	for _, row := range rows {
		s := mapper.MapStudentRowToDomain(row)
		students = append(students, &s)
	}

	return students, nil
}

// UpdateStudent updates an existing student
func (r *Repository) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapUpdateStudentParams(arg)

	row, err := q.UpdateStudent(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapStudentRowToDomain(row)
	return &out, nil
}
