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

func (h *Handler) CreateVendor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.Vendor
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateVendor(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create vendor: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "vendor created successfully", data)
}

func (h *Handler) ListVendors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("institute_id")
	if idStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	data, err := h.service.ListVendors(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list vendors: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "vendors fetched successfully", data)
}

func (h *Handler) CreatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.PurchaseOrder
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreatePurchaseOrder(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create purchase order: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "purchase order created successfully", data)
}

func (h *Handler) AddPurchaseItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.PurchaseItem
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.AddPurchaseItem(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to add purchase item: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "purchase item added successfully", data)
}

func (h *Handler) UpdatePurchaseStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	instStr := r.URL.Query().Get("institute_id")

	type StatusReq struct {
		Status domain.PurchaseStatus `json:"status"`
	}
	var req StatusReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid status payload: "+err.Error())
		return
	}

	id, _ := uuid.Parse(idStr)
	inst, _ := uuid.Parse(instStr)

	err := h.service.UpdatePurchaseStatus(r.Context(), id, inst, req.Status)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to update status: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "purchase status updated successfully", nil)
}

// ========================= CREATE VENDOR =========================

// SERVICE
func (s *Service) CreateVendor(ctx context.Context, arg domain.Vendor) (*domain.Vendor, error) {
	return s.repo.CreateVendor(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateVendor(ctx context.Context, arg domain.Vendor) (*domain.Vendor, error) {

	return &arg, nil
}

// ========================= LIST VENDORS =========================

// SERVICE
func (s *Service) ListVendors(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vendor, error) {
	return s.repo.ListVendors(ctx, instituteID)
}

// REPOSITORY
func (r *Repository) ListVendors(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vendor, error) {

	return nil, nil
}

// ========================= CREATE PURCHASE ORDER =========================

// SERVICE
func (s *Service) CreatePurchaseOrder(ctx context.Context, arg domain.PurchaseOrder) (*domain.PurchaseOrder, error) {
	return s.repo.CreatePurchaseOrder(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreatePurchaseOrder(ctx context.Context, arg domain.PurchaseOrder) (*domain.PurchaseOrder, error) {

	return &arg, nil
}

// ========================= ADD PURCHASE ITEM =========================

// SERVICE
func (s *Service) AddPurchaseItem(ctx context.Context, arg domain.PurchaseItem) (*domain.PurchaseItem, error) {
	return s.repo.AddPurchaseItem(ctx, arg)
}

// REPOSITORY
func (r *Repository) AddPurchaseItem(ctx context.Context, arg domain.PurchaseItem) (*domain.PurchaseItem, error) {

	return &arg, nil
}

// ========================= UPDATE PURCHASE STATUS =========================

// SERVICE
func (s *Service) UpdatePurchaseStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.PurchaseStatus) error {
	return s.repo.UpdatePurchaseStatus(ctx, id, instituteID, status)
}

// REPOSITORY
func (r *Repository) UpdatePurchaseStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.PurchaseStatus) error {
	return nil
}
