package auth

import (
	"context"
	"errors"
	"swiftschool/domain"
	"swiftschool/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *AuthRepository
}

func NewService(db *helper.PostgresWrapper) *AuthService {
	return &AuthService{
		repo: NewRepository(db),
	}
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// Get user by email
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Validate password
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token}, nil
}

var jwtSecret = []byte("CHANGE_ME_SECRET")

func (s *AuthService) generateToken(u *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"userId":      u.ID,
		"role":        u.RoleType,
		"instituteId": u.InstituteID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
