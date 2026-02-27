package core

import (
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Register a new student in the system
// @Tags Core - Students
// @Accept json
// @Produce json
// @Param student body dto.CreateStudentRequest true "Student details"
// @Success 201 {object} dto.SuccessResponse{data=dto.StudentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/register [post]
func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var student domain.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// We can trust the InstituteID in the body OR override it from the request context/header
	// For security, it should ideally come from the authenticated session.
	// For now, let's override/ensure it's set if we have it in header.
	instID, err := helper.GetInstituteID(r)
	if err == nil {
		student.InstituteID = instID
	} else if student.InstituteID == [16]byte{} { // Check if empty UUID
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute ID is required")
		return
	}

	data, err := h.service.CreateStudent(r.Context(), student)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "student created successfully", data)
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete a student by ID
// @Tags Core - Students
// @Produce json
// @Param id query string true "Student ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/delete [delete]
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	studentID, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteStudent(r.Context(), instID, studentID); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student deleted successfully", nil)
}

// GetStudentFullProfile godoc
// @Summary Get student full profile
// @Description Retrieve complete student profile including guardians and addresses
// @Tags Core - Students
// @Produce json
// @Param id query string true "Student ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.StudentFullProfileResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/profile [get]
func (h *Handler) GetStudentFullProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	studentID, err := helper.ParseRequiredUUIDFromQuery(r, "id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.GetStudentFullProfile(r.Context(), instID, studentID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get student profile: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student profile retrieved successfully", data)
}

// ListStudentsByClass godoc
// @Summary List students by class
// @Description Retrieve all students in a specific class
// @Tags Core - Students
// @Produce json
// @Param class_id query string true "Class ID"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.StudentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/list_by_class [get]
func (h *Handler) ListStudentsByClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	classID, err := helper.ParseRequiredUUIDFromQuery(r, "class_id")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid class id: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.ListStudentsByClass(r.Context(), instID, classID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list students: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "students retrieved successfully", data)
}

// SearchStudents godoc
// @Summary Search students
// @Description Search for students by name or admission number
// @Tags Core - Students
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.StudentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/search [get]
func (h *Handler) SearchStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query, err := helper.GetRequiredQueryParam(r, "q")
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "search query required")
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.SearchStudents(r.Context(), instID, query)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to search students: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "students retrieved successfully", data)
}

// UpdateStudent godoc
// @Summary Update a student
// @Description Update an existing student's information
// @Tags Core - Students
// @Accept json
// @Produce json
// @Param student body dto.UpdateStudentRequest true "Updated student details"
// @Success 200 {object} dto.SuccessResponse{data=dto.StudentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /students/update [put]
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var student domain.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	instID, err := helper.GetInstituteID(r)
	if err == nil {
		student.InstituteID = instID
	}

	data, err := h.service.UpdateStudent(r.Context(), student)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student updated successfully", data)
}
