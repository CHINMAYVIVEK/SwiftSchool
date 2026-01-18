package core

import (
	"context"

	"swiftschool/domain"
	"swiftschool/internal/db"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

// CreateInstitute inserts a new institute record into the database
func (r *Repository) CreateInstitute(ctx context.Context, inst domain.Institute) (*domain.Institute, error) {
	// Validate first
	if err := inst.Validate(); err != nil {
		return nil, err
	}

	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapInstituteDomainToParams(inst)
	row, err := q.CreateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapInstituteRowToDomain(row)
	return &out, nil
}

// DeleteInstitute removes an institute record from the database
func (r *Repository) DeleteInstitute(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return err
	}

	// Based on query definition, it might need more params or just ID.
	// Assuming ID is enough for soft delete or strict delete.
	// If query signature requires params struct:
	params := db.DeleteInstituteParams{
		ID: id,
	}

	return q.DeleteInstitute(ctx, params)
}

// GetInstituteById retrieves an institute by ID from the database
func (r *Repository) GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// GetInstituteById query is specific enough.
	// Check queries.sql.go for exact method name. It is GetInstituteById.
	row, err := q.GetInstituteById(ctx, id)
	if err != nil {
		return nil, err
	}

	out := mapper.MapInstituteRowToDomain(row)
	return &out, nil
}

// GetInstituteByCode retrieves an institute by code from the database
func (r *Repository) GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	row, err := q.GetInstituteByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	out := mapper.MapInstituteRowToDomain(row)
	return &out, nil
}

// ListInstitutes retrieves all institutes from the database
func (r *Repository) ListInstitutes(ctx context.Context) ([]*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	rows, err := q.ListInstitutes(ctx)
	if err != nil {
		return nil, err
	}

	var institutes []*domain.Institute
	for _, row := range rows {
		i := mapper.MapInstituteRowToDomain(row)
		institutes = append(institutes, &i)
	}

	return institutes, nil
}

// UpdateInstitute updates an existing institute record in the database
func (r *Repository) UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapUpdateInstituteParams(arg)
	row, err := q.UpdateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	out := mapper.MapInstituteRowToDomain(row)
	return &out, nil
}
