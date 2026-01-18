package core

import (
	"context"

	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

// CreateDepartment inserts a new department record into the database
func (r *Repository) CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapDomainDepartmentToDBParams(arg)
	row, err := q.CreateDepartment(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapDBDepartmentToDomain(row)
	return &out, nil
}

// DeleteDepartment removes a department record from the database
func (r *Repository) DeleteDepartment(ctx context.Context, instituteID, id uuid.UUID) error {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return err
	}

	// Try if DeleteDepartment accepts params
	return q.DeleteDepartment(ctx, db.DeleteDepartmentParams{
		ID:          id,
		InstituteID: uuid.NullUUID{UUID: instituteID, Valid: true},
	})
}

// ListDepartments retrieves all departments for an institute
func (r *Repository) ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// Wait, ListDepartments query might take (instituteID) or (instituteID, search).
	// SQLC generated ListDepartments takes uuid.NullUUID directly, not a struct
	rows, err := q.ListDepartments(ctx, helper.ToNullUUID(instituteID))
	if err != nil {
		return nil, err
	}

	var departments []*domain.Department
	for _, row := range rows {
		d := mapper.MapDBDepartmentToDomain(row)
		departments = append(departments, &d)
	}

	return departments, nil
}

// UpdateDepartment updates an existing department record
func (r *Repository) UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapUpdateDepartmentParams(arg)
	row, err := q.UpdateDepartment(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapDBDepartmentToDomain(row)
	return &out, nil
}
