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

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.Account
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateAccount(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create account: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "account created successfully", data)
}

func (h *Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.ListAccounts(r.Context(), id)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch accounts: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "accounts fetched successfully", data)
}

func (h *Handler) CreateJournalEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.JournalEntry
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateJournalEntry(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create journal entry: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "journal entry created successfully", data)
}

func (h *Handler) CreateJournalItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req domain.JournalItem
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateJournalItem(r.Context(), req)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create journal item: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "journal item created successfully", data)
}

func (h *Handler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instStr := r.URL.Query().Get("institute_id")
	accStr := r.URL.Query().Get("account_id")

	if instStr == "" || accStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id and account_id are required")
		return
	}

	inst, err := uuid.Parse(instStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	acc, err := uuid.Parse(accStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid account_id: "+err.Error())
		return
	}

	balance, err := h.service.GetAccountBalance(r.Context(), inst, acc)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch balance: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "account balance fetched successfully", map[string]float64{"balance": balance})
}

// ========================= CREATE ACCOUNT =========================

// SERVICE
func (s *Service) CreateAccount(ctx context.Context, arg domain.Account) (*domain.Account, error) {
	return s.repo.CreateAccount(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateAccount(ctx context.Context, arg domain.Account) (*domain.Account, error) {

	return &arg, nil
}

// ========================= LIST ACCOUNTS =========================

// SERVICE
func (s *Service) ListAccounts(ctx context.Context, instituteID uuid.UUID) ([]*domain.Account, error) {
	return s.repo.ListAccounts(ctx, instituteID)
}

// REPOSITORY
func (r *Repository) ListAccounts(ctx context.Context, instituteID uuid.UUID) ([]*domain.Account, error) {

	return nil, nil
}

// ========================= CREATE JOURNAL ENTRY =========================

// SERVICE
func (s *Service) CreateJournalEntry(ctx context.Context, arg domain.JournalEntry) (*domain.JournalEntry, error) {
	return s.repo.CreateJournalEntry(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateJournalEntry(ctx context.Context, arg domain.JournalEntry) (*domain.JournalEntry, error) {

	return &arg, nil
}

// ========================= CREATE JOURNAL ITEM =========================

// SERVICE
func (s *Service) CreateJournalItem(ctx context.Context, arg domain.JournalItem) (*domain.JournalItem, error) {
	return s.repo.CreateJournalItem(ctx, arg)
}

// REPOSITORY
func (r *Repository) CreateJournalItem(ctx context.Context, arg domain.JournalItem) (*domain.JournalItem, error) {

	return &arg, nil
}

// ========================= GET ACCOUNT BALANCE =========================

// SERVICE
func (s *Service) GetAccountBalance(ctx context.Context, instituteID, accountID uuid.UUID) (float64, error) {
	return s.repo.GetAccountBalance(ctx, instituteID, accountID)
}

// REPOSITORY
func (r *Repository) GetAccountBalance(ctx context.Context, instituteID, accountID uuid.UUID) (float64, error) {
	var balance float64
	return balance, nil
}
