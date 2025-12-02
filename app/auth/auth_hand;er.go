package auth

import (
	"encoding/json"
	"net/http"
	"swiftschool/helper"
)

type Handler struct {
	service AuthService
}

func NewHandler(service AuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	// Parse incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Validate required fields
	if req.Email == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "email is required")
		return
	}
	if req.Password == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "password is required")
		return
	}

	// Call service
	resp, err := h.service.Login(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "login successful", resp)
}
