package student

import (
	"text/template"

	"github.com/chinmayvivek/swiftschool/helper"
)

var logger = helper.GetLogger()

type Handler struct {
	repo      *StudentRepository
	templates map[string]*template.Template
}

func NewHandler(repo *StudentRepository) *Handler {
	if repo == nil {
		panic("student repository is required")
	}
	return &Handler{
		repo:      repo,
		templates: make(map[string]*template.Template),
	}
}
