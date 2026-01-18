package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateClass godoc
// @Summary Create a new class
// @Description Register a new class in the system
// @Tags Core - Classes
// @Accept json
// @Produce json
// @Param class body dto.CreateClassRequest true "Class details"
// @Success 201 {object} dto.SuccessResponse{data=dto.ClassResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /classes/create [post]
func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var class domain.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		class.InstituteID = instID
	}

	if class.InstituteID == [16]byte{} {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute ID is required")
		return
	}

	data, err := h.service.CreateClass(r.Context(), class)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "class created successfully", data)
}

// DeleteClass godoc
// @Summary Delete a class
// @Description Delete a class by ID
// @Tags Core - Classes
// @Produce json
// @Param id query string true "Class ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /classes/delete [delete]
func (h *Handler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid class id: "+err.Error())
		return
	}

	// DeleteClass service currently accepts only ID.
	// We might want to verify tenant ownership first?
	// But service interface is DeleteClass(ctx, id).
	// For strictness, if service changes, we need to pass instituteID.
	// Current service: func (s *Service) DeleteClass(ctx context.Context, id uuid.UUID) error

	if err := h.service.DeleteClass(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "class deleted successfully", nil)
}

// ListClasses godoc
// @Summary List classes
// @Description Retrieve all classes for an institute
// @Tags Core - Classes
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.ClassResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /classes/list [get]
func (h *Handler) ListClasses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.ListClasses(r.Context(), instID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list classes: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "classes retrieved successfully", data)
}

// UpdateClass godoc
// @Summary Update a class
// @Description Update an existing class's information
// @Tags Core - Classes
// @Accept json
// @Produce json
// @Param class body dto.UpdateClassRequest true "Updated class details"
// @Success 200 {object} dto.SuccessResponse{data=dto.ClassResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /classes/update [put]
func (h *Handler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var class domain.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		class.InstituteID = instID
	}

	data, err := h.service.UpdateClass(r.Context(), class)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update class: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "class updated successfully", data)
}
