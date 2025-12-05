package finance

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

// =================================================================================
// HANDLERS
// =================================================================================

func (h *Handler) CreateFeeHead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.FeeHead
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request payload: "+err.Error())
		return
	}

	data, err := h.service.CreateFeeHead(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create fee head: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "fee head created successfully", data)
}

func (h *Handler) ListFeeHeads(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListFeeHeads(r.Context(), instituteID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch fee heads: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "fee heads fetched successfully", data)
}

func (h *Handler) CreateFeeStructure(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.FeeStructure
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request payload: "+err.Error())
		return
	}

	data, err := h.service.CreateFeeStructure(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create fee structure: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "fee structure created successfully", data)
}

func (h *Handler) ListFeeStructures(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	sessionIDStr := r.URL.Query().Get("session_id")

	if instituteIDStr == "" || sessionIDStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id and session_id are required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid session_id: "+err.Error())
		return
	}

	data, err := h.service.ListFeeStructures(r.Context(), instituteID, sessionID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch fee structures: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "fee structures fetched successfully", data)
}

func (h *Handler) CreateFineRule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.FineRule
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request payload: "+err.Error())
		return
	}

	data, err := h.service.CreateFineRule(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create fine rule: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "fine rule created successfully", data)
}

// ========================= CREATE FEE HEAD =========================

// SERVICE
func (s *Service) CreateFeeHead(ctx context.Context, arg domain.FeeHead) (*domain.FeeHead, error) {
	return s.repo.CreateFeeHead(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateFeeHead(ctx context.Context, arg domain.FeeHead) (*domain.FeeHead, error) {

	return &arg, nil
}

// ========================= LIST FEE HEADS =========================

// SERVICE
func (s *Service) ListFeeHeads(ctx context.Context, instituteID uuid.UUID) ([]*domain.FeeHead, error) {
	return s.repo.ListFeeHeads(ctx, instituteID)
}

// REPOSITORY
func (r *Repository) ListFeeHeads(ctx context.Context, instituteID uuid.UUID) ([]*domain.FeeHead, error) {

	return nil, nil
}

// ========================= CREATE FEE STRUCTURE =========================

// SERVICE
func (s *Service) CreateFeeStructure(ctx context.Context, arg domain.FeeStructure) (*domain.FeeStructure, error) {
	return s.repo.CreateFeeStructure(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateFeeStructure(ctx context.Context, arg domain.FeeStructure) (*domain.FeeStructure, error) {

	return &arg, nil
}

// ========================= LIST FEE STRUCTURES =========================

// SERVICE
func (s *Service) ListFeeStructures(ctx context.Context, instituteID, sessionID uuid.UUID) ([]*domain.FeeStructure, error) {
	return s.repo.ListFeeStructures(ctx, instituteID, sessionID)
}

// REPOSITORY
func (r *Repository) ListFeeStructures(ctx context.Context, instituteID, sessionID uuid.UUID) ([]*domain.FeeStructure, error) {

	return nil, nil
}

// ========================= CREATE FINE RULE =========================

// SERVICE
func (s *Service) CreateFineRule(ctx context.Context, arg domain.FineRule) (*domain.FineRule, error) {
	return s.repo.CreateFineRule(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateFineRule(ctx context.Context, arg domain.FineRule) (*domain.FineRule, error) {

	return &arg, nil
}
