package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Register a new employee in the system
// @Tags Core - Employees
// @Accept json
// @Produce json
// @Param employee body dto.CreateEmployeeRequest true "Employee details"
// @Success 201 {object} dto.SuccessResponse{data=dto.EmployeeResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/create [post]
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var employee domain.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		employee.InstituteID = instID
	}

	if employee.InstituteID == [16]byte{} {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute ID is required")
		return
	}

	data, err := h.service.CreateEmployee(r.Context(), employee)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "employee created successfully", data)
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags Core - Employees
// @Produce json
// @Param id query string true "Employee ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/delete [delete]
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteEmployee(r.Context(), instID, id); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee deleted successfully", nil)
}

// GetEmployeeById godoc
// @Summary Get employee by ID
// @Description Retrieve an employee by ID
// @Tags Core - Employees
// @Produce json
// @Param id query string true "Employee ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.EmployeeResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/get [get]
func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetEmployeeById(r.Context(), instID, id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee retrieved successfully", data)
}

// GetEmployeeFullProfile godoc
// @Summary Get employee full profile
// @Description Retrieve complete employee profile including related data
// @Tags Core - Employees
// @Produce json
// @Param id query string true "Employee ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.EmployeeFullProfileResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/profile [get]
func (h *Handler) GetEmployeeFullProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid employee id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetEmployeeFullProfile(r.Context(), instID, id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get employee profile: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee profile retrieved successfully", data)
}

// ListEmployees godoc
// @Summary List employees
// @Description Retrieve all employees for an institute
// @Tags Core - Employees
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.EmployeeResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/list [get]
func (h *Handler) ListEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.ListEmployees(r.Context(), instID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list employees: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employees retrieved successfully", data)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an existing employee's information
// @Tags Core - Employees
// @Accept json
// @Produce json
// @Param employee body dto.UpdateEmployeeRequest true "Updated employee details"
// @Success 200 {object} dto.SuccessResponse{data=dto.EmployeeResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /employees/update [put]
func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var employee domain.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		employee.InstituteID = instID
	}

	data, err := h.service.UpdateEmployee(r.Context(), employee)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update employee: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "employee updated successfully", data)
}
