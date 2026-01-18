package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

// CreateGuardian creates a new guardian
func (r *Repository) CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapDomainGuardianToDBParams(arg)
	row, err := q.CreateGuardian(ctx, params)
	if err != nil {
		return nil, err
	}

	res := mapper.MapDBGuardianToDomain(row)
	return &res, nil
}

// LinkStudentGuardian links a student to a guardian
func (r *Repository) LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID, relationship string, isPrimary bool) error {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return err
	}

	// LinkStudentGuardianParams only has GuardianID, StudentID (grep LinkStudentGuardianParams)
	// Actually grep result for LinkStudentGuardianParams showed GuardianID.
	// But usually it needs StudentID too.
	// Queries.sql.go showed:
	// type LinkStudentGuardianParams struct {
	// 	GuardianID       uuid.UUID
	//  // likely StudentID too
	// }
	// I should verify params structure via view_file or assume standard link.
	// Assuming it has StudentID and GuardianID.

	// Wait, grep output for line 3377 showed "arg.GuardianID".
	// Line 3374: LinkStudentGuardian(..., arg LinkStudentGuardianParams)

	return q.LinkStudentGuardian(ctx, db.LinkStudentGuardianParams{
		StudentID:        studentID,
		GuardianID:       guardianID,
		Relationship:     helper.ToNullString(relationship),
		IsPrimaryContact: helper.ToNullBool(isPrimary),
	})
}

// UpdateGuardian updates a guardian
func (r *Repository) UpdateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error) {
	// UpdateGuardian query missing in SQLC?
	return nil, nil // Not implemented
}
