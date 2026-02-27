package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateDepartment godoc
// @Summary Create a new department
// @Description Register a new department in the system
// @Tags Core - Departments
// @Accept json
// @Produce json
// @Param department body dto.CreateDepartmentRequest true "Department details"
// @Success 201 {object} dto.SuccessResponse{data=dto.DepartmentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /departments/create [post]
func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var department domain.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		department.InstituteID = instID
	}

	if department.InstituteID == [16]byte{} {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute ID is required")
		return
	}

	data, err := h.service.CreateDepartment(r.Context(), department)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "department created successfully", data)
}

// DeleteDepartment godoc
// @Summary Delete a department
// @Description Delete a department by ID
// @Tags Core - Departments
// @Produce json
// @Param id query string true "Department ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /departments/delete [delete]
func (h *Handler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid department id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteDepartment(r.Context(), instID, id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "department deleted successfully", nil)
}

// ListDepartments godoc
// @Summary List departments
// @Description Retrieve all departments for an institute
// @Tags Core - Departments
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.DepartmentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /departments/list [get]
func (h *Handler) ListDepartments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.ListDepartments(r.Context(), instID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list departments: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "departments retrieved successfully", data)
}

// UpdateDepartment godoc
// @Summary Update a department
// @Description Update an existing department's information
// @Tags Core - Departments
// @Accept json
// @Produce json
// @Param department body dto.UpdateDepartmentRequest true "Updated department details"
// @Success 200 {object} dto.SuccessResponse{data=dto.DepartmentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /departments/update [put]
func (h *Handler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var department domain.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		department.InstituteID = instID
	}

	data, err := h.service.UpdateDepartment(r.Context(), department)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update department: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "department updated successfully", data)
}
