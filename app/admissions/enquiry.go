package admissions

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

// ========================= CREATE ENQUIRY =========================
func (h *Handler) CreateEnquiry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var enquiry domain.AdmissionEnquiry
	if err := json.NewDecoder(r.Body).Decode(&enquiry); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateEnquiry(r.Context(), enquiry)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create enquiry: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "enquiry created successfully", data)
}

// ========================= LIST ENQUIRIES =========================
func (h *Handler) ListEnquiries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	if instituteIDStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id is required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	data, err := h.service.ListEnquiries(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch enquiries: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "enquiries fetched successfully", data)
}

// ========================= UPDATE ENQUIRY STATUS =========================
func (h *Handler) UpdateEnquiryStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		ID          string                 `json:"id"`
		InstituteID string                 `json:"institute_id"`
		Status      domain.AdmissionStatus `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid enquiry id: "+err.Error())
		return
	}

	instituteID, err := uuid.Parse(req.InstituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute id: "+err.Error())
		return
	}

	if err := h.service.UpdateEnquiryStatus(r.Context(), id, instituteID, req.Status); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update enquiry status: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "enquiry status updated successfully", nil)
}

//////////////////////////////////////////////////////
// ========================= CREATE ENQUIRY =========================

// SERVICE
func (s *Service) CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error) {
	enquiry, err := s.repo.CreateEnquiry(ctx, arg)
	if err != nil {
		return nil, err
	}
	return enquiry, nil
}

// REPOSITORY
func (r *Repository) CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= LIST ENQUIRIES =========================

// SERVICE
func (s *Service) ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error) {
	enquiries, err := s.repo.ListEnquiries(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return enquiries, nil
}

// REPOSITORY
func (r *Repository) ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= UPDATE ENQUIRY STATUS =========================

// SERVICE
func (s *Service) UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error {
	return s.repo.UpdateEnquiryStatus(ctx, id, instituteID, status)
}

// REPOSITORY
func (r *Repository) UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error {
	return nil
}
