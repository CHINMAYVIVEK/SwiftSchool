package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateInstitute godoc
// @Summary Create a new institute
// @Description Create a new educational institute in the system
// @Tags Core - Institutes
// @Accept json
// @Produce json
// @Param institute body dto.CreateInstituteRequest true "Institute details"
// @Success 201 {object} dto.SuccessResponse{data=dto.InstituteResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /institutes/register [post]
func (h *Handler) CreateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var institute domain.Institute
	if err := json.NewDecoder(r.Body).Decode(&institute); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateInstitute(r.Context(), institute)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "institute creation failed: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "institute created successfully", data)
}

// DeleteInstitute godoc
// @Summary Delete an institute
// @Description Soft delete an institute by ID
// @Tags Core - Institutes
// @Produce json
// @Param id query string true "Institute ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /institutes/delete [delete]
func (h *Handler) DeleteInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id: "+err.Error())
		return
	}

	if err := h.service.DeleteInstitute(r.Context(), id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute deleted successfully", nil)
}

// GetInstituteById godoc
// @Summary Get institute by ID
// @Description Retrieve an institute by its unique ID
// @Tags Core - Institutes
// @Produce json
// @Param id query string true "Institute ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.InstituteResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /institutes/get [get]
func (h *Handler) GetInstituteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id: "+err.Error())
		return
	}

	data, err := h.service.GetInstituteById(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute retrieved successfully", data)
}

// GetInstituteByCode retrieves an institute by its unique code
func (h *Handler) GetInstituteByCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	code, err := helper.GetRequiredQueryParam(r, "code")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute code is required")
		return
	}

	data, err := h.service.GetInstituteByCode(r.Context(), code)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get institute: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute retrieved successfully", data)
}

// ListInstitutes godoc
// @Summary List all institutes
// @Description Retrieve a list of all institutes in the system
// @Tags Core - Institutes
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.InstituteResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /institutes/list [get]
func (h *Handler) ListInstitutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	data, err := h.service.ListInstitutes(r.Context())
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list institutes: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institutes retrieved successfully", data)
}

// UpdateInstitute godoc
// @Summary Update an institute
// @Description Update an existing institute's information
// @Tags Core - Institutes
// @Accept json
// @Produce json
// @Param institute body dto.UpdateInstituteRequest true "Updated institute details"
// @Success 200 {object} dto.SuccessResponse{data=dto.InstituteResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /institutes/update [put]
func (h *Handler) UpdateInstitute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var institute domain.Institute
	if err := json.NewDecoder(r.Body).Decode(&institute); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.UpdateInstitute(r.Context(), institute)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "institute update failed: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "institute updated successfully", data)
}
