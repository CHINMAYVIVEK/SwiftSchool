package auth

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// CreateUser creates a new user in the system
func (s *Service) CreateUser(ctx context.Context, arg domain.User) (*domain.User, error) {
	return s.repo.CreateUser(ctx, arg)
}

// GetUserByUsername retrieves a user by username
func (s *Service) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}

// GetUserById retrieves a user by ID
func (s *Service) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}

// UpdateUserPassword updates a user's password
func (s *Service) UpdateUserPassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
	return s.repo.UpdateUserPassword(ctx, id, passwordHash)
}

// UpdateUserStatus updates a user's active status
func (s *Service) UpdateUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return s.repo.UpdateUserStatus(ctx, id, isActive)
}

// ListUsersByRole retrieves all users with a specific role in an institute
func (s *Service) ListUsersByRole(ctx context.Context, instituteID uuid.UUID, roleType domain.UserRole) ([]*domain.User, error) {
	return s.repo.ListUsersByRole(ctx, instituteID, roleType)
}
