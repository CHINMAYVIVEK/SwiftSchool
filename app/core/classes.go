package core

import (
	"net/http"
)

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {}

func (s *Service) CreateClass() {}
func (s *Service) DeleteClass() {}

func (r *Repository) CreateClass() {}
func (r *Repository) DeleteClass() {}
