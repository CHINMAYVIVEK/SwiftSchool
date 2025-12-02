package core

import (
	"context"
	"swiftschool/internal/db"
)

func (s *Service) CreateAddress(ctx context.Context, arg db.CreateAddressParams) (db.CoreAddress, error) {
	coreAddress, err := s.repo.CreateAddress(ctx, arg)
	if err != nil {
		return coreAddress, err
	}
	return coreAddress, nil
}

func (r *Repository) CreateAddress(ctx context.Context, arg db.CreateAddressParams) (db.CoreAddress, error) {
	coreAddress, err := r.CreateAddress(ctx, arg)
	if err != nil {
		return coreAddress, err
	}
	return coreAddress, nil
}
