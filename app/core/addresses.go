package core

import (
	"context"
	"swiftschool/domain"
)

func (s *Service) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	coreAddress, err := s.repo.CreateAddress(ctx, arg)
	if err != nil {
		return coreAddress, err
	}
	return coreAddress, nil
}

func (r *Repository) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	return nil, nil
}
