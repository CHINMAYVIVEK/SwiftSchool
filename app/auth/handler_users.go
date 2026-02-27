package auth

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Register a new user in the system
// @Tags Auth - Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User details"
// @Success 201 {object} dto.SuccessResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/register [post]
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

// GetUserByUsername godoc
// @Summary Get user by username
// @Description Retrieve a user by their username
// @Tags Auth - Users
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} dto.SuccessResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/get_by_username [get]
func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	username, err := helper.GetRequiredQueryParam(r, "username")
	if err != nil {
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

// GetUserById godoc
// @Summary Get user by ID
// @Description Retrieve a user by their unique ID
// @Tags Auth - Users
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/get_by_id [get]
func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
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

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update a user's password
// @Tags Auth - Users
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserPasswordRequest true "Password update details"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/update_password [patch]
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

// UpdateUserStatus godoc
// @Summary Update user status
// @Description Activate or deactivate a user account
// @Tags Auth - Users
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserStatusRequest true "Status update details"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/update_status [patch]
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

// ListUsersByRole godoc
// @Summary List users by role
// @Description Retrieve all users with a specific role in an institute
// @Tags Auth - Users
// @Produce json
// @Param institute_id query string true "Institute ID"
// @Param role query string true "User role"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /auth/users/list_by_role [get]
func (h *Handler) ListUsersByRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteID, err := helper.ParseRequiredUUIDFromQuery(r, "institute_id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	roleStr, err := helper.GetRequiredQueryParam(r, "role")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "role is required")
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
