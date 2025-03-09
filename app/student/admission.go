package student

import (
	"net/http"

	"github.com/chinmayvivek/swiftschool/helper"
	"github.com/google/uuid"
)

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	err := h.repo.admissionStatus()
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "Error executing template")
		return // Added return statement after error
	}
	id := uuid.New().String()
	helper.NewSuccessResponse(w, id, "Admission status updated successfully", http.StatusOK)
}
func (s *StudentRepository) admissionStatus() error {
	return nil
}
