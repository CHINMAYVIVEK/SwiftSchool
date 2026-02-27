package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateAcademicSession godoc
// @Summary Create a new academic session
// @Description Register a new academic session in the system
// @Tags Core - Academic Sessions
// @Accept json
// @Produce json
// @Param session body dto.CreateAcademicSessionRequest true "Session details"
// @Success 201 {object} dto.SuccessResponse{data=dto.AcademicSessionResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /academic_sessions/create [post]
func (h *Handler) CreateAcademicSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var session domain.AcademicSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		session.InstituteID = instID
	}

	if session.InstituteID == [16]byte{} {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute ID is required")
		return
	}

	data, err := h.service.CreateAcademicSession(r.Context(), session)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "session created successfully", data)
}

// GetActiveSession godoc
// @Summary Get active academic session
// @Description Retrieve the active academic session for the institute
// @Tags Core - Academic Sessions
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=dto.AcademicSessionResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /academic_sessions/active [get]
func (h *Handler) GetActiveSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetActiveSession(r.Context(), instID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get active session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "active session retrieved successfully", data)
}

// ListAcademicSessions godoc
// @Summary List academic sessions
// @Description Retrieve all academic sessions for an institute
// @Tags Core - Academic Sessions
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.AcademicSessionResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /academic_sessions/list [get]
func (h *Handler) ListAcademicSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.ListAcademicSessions(r.Context(), instID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list sessions: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "sessions retrieved successfully", data)
}

// UpdateAcademicSession godoc
// @Summary Update an academic session
// @Description Update an existing academic session
// @Tags Core - Academic Sessions
// @Accept json
// @Produce json
// @Param session body dto.UpdateAcademicSessionRequest true "Updated session details"
// @Success 200 {object} dto.SuccessResponse{data=dto.AcademicSessionResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /academic_sessions/update [put]
func (h *Handler) UpdateAcademicSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var session domain.AcademicSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		session.InstituteID = instID
	}

	data, err := h.service.UpdateAcademicSession(r.Context(), session)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update session: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "session updated successfully", data)
}
