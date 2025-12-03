package core

import (
	"context"
	"swiftschool/domain"
)

//////////////////////////////////////////////////////
//                 ADDRESS METHODS                  //
//////////////////////////////////////////////////////

// ========================= CREATE =========================
func (s *Service) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	return s.repo.CreateAddress(ctx, arg)
}

func (r *Repository) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	// TODO: implement DB logic here
	return nil, nil
}
