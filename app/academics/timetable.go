package academics

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                    HANDLER                       //
//////////////////////////////////////////////////////

// ========================= CREATE TIMETABLE ENTRY =========================

// CreateTimetableEntry godoc
// @Summary Create a new timetable entry
// @Description Create a new timetable entry for a class
// @Tags Academics - Timetable
// @Accept json
// @Produce json
// @Param entry body dto.CreateTimetableEntryRequest true "Timetable entry details"
// @Success 201 {object} dto.SuccessResponse{data=dto.TimetableEntryResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /timetable/register [post]
func (h *Handler) CreateTimetableEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var entry domain.TimetableEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateTimetableEntry(r.Context(), entry)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create timetable entry: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "timetable entry created successfully", data)
}

// ========================= GET CLASS TIMETABLE =========================

// GetClassTimetable godoc
// @Summary Get class timetable
// @Description Retrieve the complete timetable for a specific class
// @Tags Academics - Timetable
// @Produce json
// @Param class_id query string true "Class ID"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.TimetableEntryResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /timetable/list [get]
func (h *Handler) GetClassTimetable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	classIDStr := r.URL.Query().Get("class_id")
	dayStr := r.URL.Query().Get("day")

	if instituteIDStr == "" || classIDStr == "" || dayStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id, class_id and day are required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid class_id: "+err.Error())
		return
	}

	day := domain.DayOfWeek(dayStr)

	data, err := h.service.GetClassTimetable(r.Context(), instituteID, classID, day)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch timetable: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "timetable fetched successfully", data)
}

// ========================= SERVICE + REPO =========================

// SERVICE
func (s *Service) CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error) {
	timeTableEntry, err := s.repo.CreateTimetableEntry(ctx, arg)
	if err != nil {
		return nil, err
	}
	return timeTableEntry, nil
}

// REPOSITORY
func (r *Repository) CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error) {

	return nil, nil
}

// SERVICE
func (s *Service) GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error) {
	timeTableEntries, err := s.repo.GetClassTimetable(ctx, instituteID, classID, day)
	if err != nil {
		return nil, err
	}
	return timeTableEntries, nil
}

// REPOSITORY
func (r *Repository) GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error) {

	return nil, nil
}
