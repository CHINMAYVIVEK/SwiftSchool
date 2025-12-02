package core

import (
	"swiftschool/helper"
)

type CoreService struct {
	repo *CoreRepository
}

// NewService creates a new InstitutesService instance
func NewService(db *helper.PostgresWrapper) *CoreService {
	return &CoreService{
		repo: NewRepository(db),
	}
}
