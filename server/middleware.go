package server

import (
	"net/http"
)

// RequireRole checks if the logged-in user has the required role.
func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := GetUserRoleFromSession(r)
		if !ok || userRole != role {
			// Redirect unauthorized users to login page
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

// Example: session cookie key
const sessionRoleKey = "user_role"

// GetUserRoleFromSession returns the logged-in user's role from the session.
// Returns (role string, ok bool)
func GetUserRoleFromSession(r *http.Request) (string, bool) {
	cookie, err := r.Cookie(sessionRoleKey)
	if err != nil {
		// No cookie found
		return "", false
	}

	// In a real app, you might validate this value, decode JWT, or lookup DB
	role := cookie.Value
	if role == "" {
		return "", false
	}

	return role, true
}
