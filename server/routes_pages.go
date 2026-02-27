package server

import (
	"net/http"
	"swiftschool/pages"
)

func (s *Server) registerPageRoutes() {
	pageHandler := pages.NewHandler(s.db)

	// Serve static assets for CSS, JS, images
	s.mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("template/assets"))))

	// Helper to wrap routes with optional role middleware
	register := func(path string, handler http.HandlerFunc, role ...string) {
		if len(role) > 0 {
			handler = RequireRole(role[0], handler)
		}
		s.mux.HandleFunc(path, handler)
	}

	// ================= AUTH =================
	register("/", pageHandler.LoginPage)
	// register("/logout", pageHandler.Logout)

	// // ================= STUDENT =================
	// register("/student/dashboard", pageHandler.StudentDashboard, "student")
	// register("/student/profile", pageHandler.StudentProfile, "student")

	// // ================= TEACHER =================
	// register("/teacher/dashboard", pageHandler.TeacherDashboard, "teacher")

	// // ================= ADMIN =================
	// register("/admin/dashboard", pageHandler.AdminDashboard, "admin")
}
