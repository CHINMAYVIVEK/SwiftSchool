package core

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/mapper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                  STUDENT METHODS                //
//////////////////////////////////////////////////////

// ========================= HANDLER =========================

// ---------------- CREATE STUDENT ----------------
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

	data, err := h.service.CreateStudent(r.Context(), student)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "student created successfully", data)
}

// ---------------- DELETE STUDENT ----------------
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	studentIDStr := r.URL.Query().Get("id")
	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id")
		return
	}

	if err := h.service.DeleteStudent(r.Context(), studentID); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to delete student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student deleted successfully", nil)
}

// ---------------- GET FULL PROFILE ----------------
func (h *Handler) GetStudentFullProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	studentIDStr := r.URL.Query().Get("id")
	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid student id")
		return
	}

	data, err := h.service.GetStudentFullProfile(r.Context(), studentID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to get student profile: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student profile retrieved successfully", data)
}

// ---------------- LIST STUDENTS BY CLASS ----------------
func (h *Handler) ListStudentsByClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	classIDStr := r.URL.Query().Get("class_id")
	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid class id")
		return
	}

	data, err := h.service.ListStudentsByClass(r.Context(), classID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list students: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "students retrieved successfully", data)
}

// ---------------- SEARCH STUDENTS ----------------
func (h *Handler) SearchStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "search query required")
		return
	}

	data, err := h.service.SearchStudents(r.Context(), query)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to search students: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "students retrieved successfully", data)
}

// ---------------- UPDATE STUDENT ----------------
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

	data, err := h.service.UpdateStudent(r.Context(), student)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update student: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "student updated successfully", data)
}

// ========================= CREATE =========================
func (s *Service) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	return s.repo.CreateStudent(ctx, arg)
}

func (r *Repository) CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {

	// Apply timeout
	ctx, cancel := r.db.WithTimeout(ctx)
	defer cancel()

	// Queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}

	// ------------------------------------
	// Build SQLC params (SAFE & CLEAN)
	// ------------------------------------

	params := mapper.MapStudentDomainToParams(arg)
	// ------------------------------------
	// Insert student
	// ------------------------------------
	row, err := q.CreateStudent(ctx, params)
	if err != nil {
		return nil, err
	}

	// ------------------------------------
	// Map SQL → Domain
	// ------------------------------------
	out := mapper.MapStudentRowToDomain(row)

	return &out, nil
}

// ========================= DELETE =========================
func (s *Service) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStudent(ctx, id)
}

func (r *Repository) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	// TODO: implement DB logic here
	return nil
}

// ========================= GET FULL PROFILE =========================
func (s *Service) GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	return s.repo.GetStudentFullProfile(ctx, id)
}

func (r *Repository) GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= LIST BY CLASS =========================
func (s *Service) ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]*domain.Student, error) {
	return s.repo.ListStudentsByClass(ctx, classID)
}

func (r *Repository) ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= SEARCH =========================
func (s *Service) SearchStudents(ctx context.Context, query string) ([]*domain.Student, error) {
	return s.repo.SearchStudents(ctx, query)
}

func (r *Repository) SearchStudents(ctx context.Context, query string) ([]*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}

// ========================= UPDATE =========================
func (s *Service) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	return s.repo.UpdateStudent(ctx, arg)
}

func (r *Repository) UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error) {
	// TODO: implement DB logic here
	return nil, nil
}
