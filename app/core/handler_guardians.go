package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

// CreateGuardian godoc
// @Summary Create a new guardian
// @Description Register a new guardian in the system
// @Tags Core - Guardians
// @Accept json
// @Produce json
// @Param guardian body dto.CreateGuardianRequest true "Guardian details"
// @Success 201 {object} dto.SuccessResponse{data=dto.GuardianResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /guardians/create [post]
func (h *Handler) CreateGuardian(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var guardian domain.Guardian
	if err := json.NewDecoder(r.Body).Decode(&guardian); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Guardian is global/shared, doesn't need InstituteID from header for creation in DB
	// But logically we might want to check permissions.
	// For now, proceed.

	data, err := h.service.CreateGuardian(r.Context(), guardian)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create guardian: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "guardian created successfully", data)
}

// LinkStudentGuardian godoc
// @Summary Link student to guardian
// @Description meaningful description
// @Tags Core - Guardians
// @Accept json
// @Produce json
// @Param link body dto.LinkStudentGuardianRequest true "Link details"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /guardians/link [post]
func (h *Handler) LinkStudentGuardian(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	type LinkRequest struct {
		StudentID    string `json:"student_id"`
		GuardianID   string `json:"guardian_id"`
		Relationship string `json:"relationship"`
		IsPrimary    bool   `json:"is_primary"`
	}

	var req LinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	studentID, err := uuid.Parse(req.StudentID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id")
		return
	}
	guardianID, err := uuid.Parse(req.GuardianID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid guardian id")
		return
	}

	if err := h.service.LinkStudentGuardian(r.Context(), studentID, guardianID, req.Relationship, req.IsPrimary); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to link guardian: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "guardian linked successfully", nil)
}
