package pages

import (
	"net/http"

	"swiftschool/helper"
	"swiftschool/internal/database"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

type Handler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{db: db}
}

// render executes a template for the given view and page.
func (h *Handler) render(w http.ResponseWriter, view helper.ViewType, page string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := helper.Render(w, view, page, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
