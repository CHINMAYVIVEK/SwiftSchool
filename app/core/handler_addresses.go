package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateAddress godoc
// @Summary Create a new address
// @Description Register a new address for an entity (user, student, etc.)
// @Tags Core - Addresses
// @Accept json
// @Produce json
// @Param address body dto.CreateAddressRequest true "Address details"
// @Success 201 {object} dto.SuccessResponse{data=dto.AddressResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /addresses/create [post]
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
