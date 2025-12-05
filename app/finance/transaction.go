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

func (h *Handler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.Invoice
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateInvoice(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create invoice: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "invoice created successfully", data)
}

func (h *Handler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	instStr := r.URL.Query().Get("institute_id")

	if idStr == "" || instStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "id and institute_id are required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid id: "+err.Error())
		return
	}
	inst, err := uuid.Parse(instStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	res, err := h.service.GetInvoiceById(r.Context(), id, inst)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusNotFound, "invoice not found: "+err.Error())
		return
	}
	helper.NewSuccessResponse(w, http.StatusOK, "invoice fetched successfully", res)
}

func (h *Handler) ListStudentInvoices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instStr := r.URL.Query().Get("institute_id")
	studentStr := r.URL.Query().Get("student_id")

	inst, _ := uuid.Parse(instStr)
	student, _ := uuid.Parse(studentStr)

	data, err := h.service.ListStudentInvoices(r.Context(), inst, student)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to list invoices: "+err.Error())
		return
	}
	helper.NewSuccessResponse(w, http.StatusOK, "invoices fetched successfully", data)
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateTransaction(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create transaction: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "transaction created successfully", data)
}

// ========================= CREATE INVOICE =========================

// SERVICE
func (s *Service) CreateInvoice(ctx context.Context, arg domain.Invoice) (*domain.Invoice, error) {
	return s.repo.CreateInvoice(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateInvoice(ctx context.Context, arg domain.Invoice) (*domain.Invoice, error) {

	return &arg, nil
}

// ========================= CREATE INVOICE ITEM =========================

// SERVICE
func (s *Service) CreateInvoiceItem(ctx context.Context, arg domain.InvoiceItem) (*domain.InvoiceItem, error) {
	return s.repo.CreateInvoiceItem(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateInvoiceItem(ctx context.Context, arg domain.InvoiceItem) (*domain.InvoiceItem, error) {

	return &arg, nil
}

// ========================= GET INVOICE BY ID =========================

// SERVICE
func (s *Service) GetInvoiceById(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, error) {
	return s.repo.GetInvoiceById(ctx, id, instituteID)
}

// REPOSITORY
func (r *Repository) GetInvoiceById(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, error) {

	return nil, nil
}

// ========================= GET INVOICE WITH ITEMS =========================

// SERVICE
func (s *Service) GetInvoiceWithItems(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, []*domain.InvoiceItem, error) {
	return s.repo.GetInvoiceWithItems(ctx, id, instituteID)
}

// REPOSITORY
func (r *Repository) GetInvoiceWithItems(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, []*domain.InvoiceItem, error) {

	return nil, nil, nil
}

// ========================= LIST STUDENT INVOICES =========================

// SERVICE
func (s *Service) ListStudentInvoices(ctx context.Context, instituteID, studentID uuid.UUID) ([]*domain.Invoice, error) {
	return s.repo.ListStudentInvoices(ctx, instituteID, studentID)
}

// REPOSITORY
func (r *Repository) ListStudentInvoices(ctx context.Context, instituteID, studentID uuid.UUID) ([]*domain.Invoice, error) {

	return nil, nil
}

// ========================= UPDATE INVOICE STATUS =========================

// SERVICE
func (s *Service) UpdateInvoiceStatus(ctx context.Context, id, instituteID uuid.UUID, amount float64, status domain.SaaSInvoiceStatus) error {
	return s.repo.UpdateInvoiceStatus(ctx, id, instituteID, amount, status)
}

// REPOSITORY
func (r *Repository) UpdateInvoiceStatus(ctx context.Context, id, instituteID uuid.UUID, amount float64, status domain.SaaSInvoiceStatus) error {
	return nil
}

// ========================= GET OVERDUE INVOICES =========================

// SERVICE
func (s *Service) GetOverdueInvoices(ctx context.Context, instituteID uuid.UUID) ([]*domain.Invoice, error) {
	return s.repo.GetOverdueInvoices(ctx, instituteID)
}

// REPOSITORY
func (r *Repository) GetOverdueInvoices(ctx context.Context, instituteID uuid.UUID) ([]*domain.Invoice, error) {

	return nil, nil
}

// ========================= CREATE TRANSACTION =========================

// SERVICE
func (s *Service) CreateTransaction(ctx context.Context, arg domain.Transaction) (*domain.Transaction, error) {
	// Add business logic to update invoice status if payment is complete
	return s.repo.CreateTransaction(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateTransaction(ctx context.Context, arg domain.Transaction) (*domain.Transaction, error) {

	return &arg, nil
}
