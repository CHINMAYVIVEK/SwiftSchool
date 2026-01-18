package common

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

// ========================= CREATE DOCUMENT =========================

// CreateDocument godoc
// @Summary Create a new document
// @Description Upload and register a new document
// @Tags Common - Documents
// @Accept json
// @Produce json
// @Param document body dto.CreateDocumentRequest true "Document details"
// @Success 201 {object} dto.SuccessResponse{data=dto.DocumentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /common/documents/create [post]
func (h *Handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var doc domain.Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateDocument(r.Context(), doc)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create document: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "document created successfully", data)
}

// ========================= LIST DOCUMENTS =========================

// ListDocuments godoc
// @Summary List documents
// @Description Retrieve all documents for a specific owner
// @Tags Common - Documents
// @Produce json
// @Param institute_id query string true "Institute ID"
// @Param owner_id query string true "Owner ID"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.DocumentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /common/documents/list [get]
func (h *Handler) ListDocuments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	instituteIDStr := r.URL.Query().Get("institute_id")
	ownerIDStr := r.URL.Query().Get("owner_id")

	if instituteIDStr == "" || ownerIDStr == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "institute_id and owner_id are required")
		return
	}

	instituteID, err := uuid.Parse(instituteIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid institute_id: "+err.Error())
		return
	}

	ownerID, err := uuid.Parse(ownerIDStr)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid owner_id: "+err.Error())
		return
	}

	data, err := h.service.ListDocuments(r.Context(), instituteID, ownerID)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to fetch documents: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusOK, "documents fetched successfully", data)
}

//////////////////////////////////////////////////////
// ========================= CREATE DOCUMENT =========================

// SERVICE
func (s *Service) CreateDocument(ctx context.Context, arg domain.Document) (*domain.Document, error) {
	doc, err := s.repo.CreateDocument(ctx, arg)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// REPOSITORY
func (r *Repository) CreateDocument(ctx context.Context, arg domain.Document) (*domain.Document, error) {
	return nil, nil
}

//////////////////////////////////////////////////////
// ========================= LIST DOCUMENTS =========================

// SERVICE
func (s *Service) ListDocuments(ctx context.Context, instituteID, ownerID uuid.UUID) ([]*domain.Document, error) {
	docs, err := s.repo.ListDocuments(ctx, instituteID, ownerID)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

// REPOSITORY
func (r *Repository) ListDocuments(ctx context.Context, instituteID, ownerID uuid.UUID) ([]*domain.Document, error) {
	return nil, nil
}
