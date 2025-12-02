package server

import (
	"net/http"
	"swiftschool/helper"
)

// Route handlers
func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var message string

	// err := helper.VaultDiagnostic(s.config.App.VaultAddr, s.config.App.VaultToken)
	// if err != nil {
	// 	message = fmt.Sprintf("Vault diagnostic failed: %v", err)
	// } else {
	// 	message = "Vault diagnostic completed successfully ✅"
	// }

	healthData := map[string]string{
		"status":           "healthy",
		"vault diagnostic": message,
	}

	helper.NewSuccessResponse(w, http.StatusOK, "Server is running", healthData)
}
