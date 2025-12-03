package core

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

//////////////////////////////////////////////////////
//                 ADDRESS METHODS                  //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE ADDRESS ----------------
func (h *Handler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var address domain.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateAddress(r.Context(), address)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create address: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "address created successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	return s.repo.CreateAddress(ctx, arg)
}

func (r *Repository) CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error) {
	// TODO: implement DB logic here
	return nil, nil
}
