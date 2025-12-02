package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
)

type CoreRepositoryInterface interface {
	InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error)
}

type CoreRepository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *CoreRepository {
	return &CoreRepository{db: db}
}
