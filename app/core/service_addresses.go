package core

import (
	"context"
	"swiftschool/domain"
)

// CreateAddress creates a new address
func (s *Service) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	return s.repo.CreateAddress(ctx, arg)
}
