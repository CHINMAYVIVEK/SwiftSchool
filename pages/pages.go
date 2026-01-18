package pages

import (
	"net/http"

	"swiftschool/helper"
)

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.render(w, helper.WebsiteView, "login", map[string]any{
		"Title": "Swfit School::Login",
	})
}

func (h *Handler) StudentDashboard(w http.ResponseWriter, r *http.Request) {
	h.render(w, helper.StudentView, "dashboard", map[string]any{
		"Title": "Student Dashboard",
	})
}

func (h *Handler) TeacherDashboard(w http.ResponseWriter, r *http.Request) {
	h.render(w, helper.SchoolView, "dashboard", map[string]any{
		"Title": "Teacher Dashboard",
	})
}

func (h *Handler) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	h.render(w, helper.SchoolView, "dashboard", map[string]any{
		"Title": "Admin Dashboard",
	})
}
