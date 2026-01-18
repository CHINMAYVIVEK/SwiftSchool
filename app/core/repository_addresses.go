package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/mapper"
)

// CreateAddress creates a new address
func (r *Repository) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	params := mapper.MapDomainAddressToDBParams(arg)
	row, err := q.CreateAddress(ctx, params)
	if err != nil {
		return nil, err
	}

	res := mapper.MapDBAddressToDomain(row)
	return &res, nil
}

// UpdateAddress updates an address
func (r *Repository) UpdateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	// UpdateAddress query missing?
	return nil, nil // Not implemented
}
