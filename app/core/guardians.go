package core

import (
	"context"
	"swiftschool/internal/db"
)

func (s *Service) CreateGuardian(ctx context.Context, arg db.CreateGuardianParams) (db.CoreGuardian, error) {
	createGuardian, err := s.repo.CreateGuardian(ctx, arg)
	if err != nil {
		return createGuardian, err
	}
	return createGuardian, nil
}

func (r *Repository) CreateGuardian(ctx context.Context, arg db.CreateGuardianParams) (db.CoreGuardian, error) {
	createGuardian, err := r.CreateGuardian(ctx, arg)
	if err != nil {
		return createGuardian, err
	}
	return createGuardian, nil
}
