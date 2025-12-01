package institutes

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	service *InstitutesService
}

func NewHandler(service *InstitutesService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InstitutesRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.Institute

	// Parse incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "name is required")
		return
	}
	if req.Code == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "code is required")
		return
	}
	if req.CreatedBy == nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "created_by is required (uuid)")
		return
	}

	// Auto-fill generated fields
	now := time.Now()
	req.ID = uuid.New()
	req.CreatedAt = now
	req.UpdatedAt = now
	req.UpdatedBy = req.CreatedBy
	req.DeletedAt = nil
	req.IsActive = true

	// Call service
	data, err := h.service.InstitutesRegistration(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create institute: "+err.Error())
		return
	}

	// Success response
	helper.NewSuccessResponse(w, http.StatusOK, "institute created successfully", data)
}
