package auth

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
)

type AuthRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type AuthRepository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// Call SQLC function
	u, err := q.GetUserByUsername(ctx, email)
	if err != nil {
		return nil, err
	}

	basemodel := domain.BaseUUIDModel{
		ID: u.ID,
	}
	// Map SQLC → Domain
	user := domain.User{
		BaseUUIDModel: basemodel,
		Username:      u.Username,
		RoleType:      domain.UserRole(u.RoleType.String),
	}

	return &user, nil
}
