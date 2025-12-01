package institutes

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
)

// InstitutesService handles business logic
type InstitutesService struct {
	repo *InstitutesRepository
}

// NewService creates a new InstitutesService instance
func NewService(db *helper.PostgresWrapper) *InstitutesService {
	return &InstitutesService{
		repo: NewRepository(db),
	}
}

// InstitutesRegistration creates a new institute and returns BaseUUIDModel
func (s *InstitutesService) InstitutesRegistration(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error) {
	// Business validations or vault logic can be added here

	data, err := s.repo.InstitutesRegistration(ctx, institute)
	if err != nil {
		return nil, err
	}
	return data, nil
}
