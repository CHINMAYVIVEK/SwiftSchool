package institutes

import (
	"context"
	"database/sql"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
)

type InstitutesRepositoryInterface interface {
	InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error)
}

type InstitutesRepository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *InstitutesRepository {
	return &InstitutesRepository{db: db}
}

func (r *InstitutesRepository) InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// Map domain → SQLC params
	params := db.CreateInstituteParams{
		Name:         institute.Name,
		Code:         institute.Code,
		CurrencyCode: sql.NullString{String: helper.ToStr(institute.CurrencyCode), Valid: institute.CurrencyCode != nil},
		LogoUrl:      sql.NullString{String: helper.ToStr(institute.LogoURL), Valid: institute.LogoURL != nil},
		Website:      sql.NullString{String: helper.ToStr(institute.Website), Valid: institute.Website != nil},
		CreatedBy:    helper.ToNullUUID(institute.CreatedBy),
		CreatedAt:    helper.ToNullTime(&institute.CreatedAt),
		IsActive:     helper.ToNullBool(institute.IsActive),
	}

	coreInstitute, err := q.CreateInstitute(ctx, params)
	if err != nil {
		return nil, err
	}

	base := domain.BaseUUIDModel{
		ID: coreInstitute.ID,
	}

	return &base, nil
}
