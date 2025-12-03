package auth

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// ========================= SERVICE =========================

func (s *Service) CreateUser(ctx context.Context, arg domain.User) (*domain.User, error) {
	// Logic: Hash password before saving
	// arg.PasswordHash = hash(arg.Password)
	return s.repo.CreateUser(ctx, arg)
}

func (s *Service) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}

// ========================= REPOSITORY =========================

func (r *Repository) CreateUser(ctx context.Context, arg domain.User) (*domain.User, error) {
	// TODO: Call SQLC queries.CreateUser
	return &arg, nil
}

func (r *Repository) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	// TODO: Call SQLC queries.GetUserById
	return nil, nil
}
