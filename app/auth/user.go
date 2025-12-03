package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

// ========================= CREATE USER =========================
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create user: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "user created successfully", data)
}

// ========================= GET USER BY USERNAME =========================
func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "username is required")
		return
	}

	data, err := h.service.GetUserByUsername(r.Context(), username)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch user: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "user fetched successfully", data)
}

// ========================= GET USER BY ID =========================
func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "id is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid id: "+err.Error())
		return
	}

	data, err := h.service.GetUserById(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch user: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "user fetched successfully", data)
}

// ========================= UPDATE USER PASSWORD =========================
func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		ID           string `json:"id"`
		PasswordHash string `json:"password_hash"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid id: "+err.Error())
		return
	}

	if err := h.service.UpdateUserPassword(r.Context(), id, req.PasswordHash); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update password: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "password updated successfully", nil)
}

// ========================= UPDATE USER STATUS =========================
func (h *Handler) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		ID       string `json:"id"`
		IsActive bool   `json:"is_active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid id: "+err.Error())
		return
	}

	if err := h.service.UpdateUserStatus(r.Context(), id, req.IsActive); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update user status: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "user status updated successfully", nil)
}

// ========================= LIST USERS BY ROLE =========================
func (h *Handler) ListUsersByRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	roleStr := r.URL.Query().Get("role")
	if instituteIDStr == "" || roleStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id and role are required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	role := domain.UserRole(roleStr)
	data, err := h.service.ListUsersByRole(r.Context(), instituteID, role)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch users: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "users fetched successfully", data)
}

//////////////////////////////////////////////////////
// ========================= CREATE USER =========================

// SERVICE
func (s *Service) CreateUser(ctx context.Context, arg domain.User) (*domain.User, error) {
	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// REPOSITORY
func (r *Repository) CreateUser(ctx context.Context, arg domain.User) (*domain.User, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= GET USER BY USERNAME =========================

// SERVICE
func (s *Service) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// REPOSITORY
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= GET USER BY ID =========================

// SERVICE
func (s *Service) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// REPOSITORY
func (r *Repository) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= UPDATE USER PASSWORD =========================

// SERVICE
func (s *Service) UpdateUserPassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
	return s.repo.UpdateUserPassword(ctx, id, passwordHash)
}

// REPOSITORY
func (r *Repository) UpdateUserPassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
	return nil
}

//////////////////////////////////////////////////////
// ========================= UPDATE USER STATUS =========================

// SERVICE
func (s *Service) UpdateUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return s.repo.UpdateUserStatus(ctx, id, isActive)
}

// REPOSITORY
func (r *Repository) UpdateUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return nil
}

//////////////////////////////////////////////////////
// ========================= LIST USERS BY ROLE =========================

// SERVICE
func (s *Service) ListUsersByRole(ctx context.Context, instituteID uuid.UUID, roleType domain.UserRole) ([]*domain.User, error) {
	users, err := s.repo.ListUsersByRole(ctx, instituteID, roleType)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// REPOSITORY
func (r *Repository) ListUsersByRole(ctx context.Context, instituteID uuid.UUID, roleType domain.UserRole) ([]*domain.User, error) {
	return nil, nil
}
