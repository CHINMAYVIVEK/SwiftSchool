package core

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
	"time"

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

	// Get SQLC queries instance
	q, err := r.db.Queries()
	if err != nil {
		return nil, err
	}
	params := db.CreateStudentParams{
		InstituteID:        arg.InstituteID,
		AdmissionNo:        arg.AdmissionNo,
		FirstName:          arg.FirstName,
		LastName:           helper.ToNullString(arg.LastName),
		Dob:                helper.ToNullTime(arg.DOB),
		Gender:             helper.ToNullString(arg.Gender),
		BloodGroup:         helper.ToNullString(arg.BloodGroup),
		SocialCategory:     helper.ToNullString(arg.SocialCategory),
		CurrentClassID:     helper.ToNullUUID(arg.CurrentClassID),
		Nationality:        helper.ToNullString(arg.Nationality),
		PreferredLanguage:  helper.ToNullString(arg.PreferredLanguage),
		SocialMediaHandles: helper.EncodeJSONB(arg.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(arg.LanguageSkills),
		CreatedBy:          helper.ToNullUUID(arg.CreatedBy),
	}
	coreStudent, err := q.CreateStudent(ctx, params)
	if err != nil {
		return nil, err
	}
	student := domain.Student{
		TenantUUIDModel: domain.TenantUUIDModel{
			InstituteID: coreStudent.InstituteID,
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        coreStudent.ID,
				CreatedAt: *helper.NullToPointer[time.Time](coreStudent.CreatedAt),
				UpdatedAt: *helper.NullToPointer[time.Time](coreStudent.UpdatedAt),
				DeletedAt: helper.NullToPointer[time.Time](coreStudent.DeletedAt),
				CreatedBy: helper.NullToPointer[uuid.UUID](coreStudent.CreatedBy),
				UpdatedBy: helper.NullToPointer[uuid.UUID](coreStudent.UpdatedBy),
			},
		},
		AdmissionNo:        coreStudent.AdmissionNo,
		FirstName:          coreStudent.FirstName,
		LastName:           helper.NullToPointer[string](coreStudent.LastName),
		DOB:                helper.NullToPointer[time.Time](coreStudent.Dob),
		Gender:             helper.NullToValue[domain.Gender](coreStudent.Gender),
		BloodGroup:         helper.NullToValue[domain.BloodGroup](coreStudent.BloodGroup),
		SocialCategory:     helper.NullToValue[domain.SocialCategory](coreStudent.SocialCategory),
		CurrentClassID:     helper.NullToPointer[uuid.UUID](coreStudent.CurrentClassID),
		Nationality:        helper.NullToPointer[string](coreStudent.Nationality),
		PreferredLanguage:  helper.NullToPointer[string](coreStudent.PreferredLanguage),
		SocialMediaHandles: helper.JSONBToValue[domain.SocialMediaHandles](coreStudent.SocialMediaHandles),
		LanguageSkills:     helper.JSONBToValue[[]domain.LanguageSkill](coreStudent.LanguageSkills),
	}

	return &student, nil
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
