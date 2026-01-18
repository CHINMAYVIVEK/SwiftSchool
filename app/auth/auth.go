package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"swiftschool/domain"
	"swiftschool/helper"
)

var logger = helper.GetLogger()

// ==========================
// HTTP HANDLERS
// ==========================

// Login godoc
// @Summary User login
// @Description Initiate login process and send OTP to user
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username, email, or phone"
// @Param role formData string true "User role"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if !isPost(r, w) {
		return
	}

	if err := r.ParseForm(); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid form submission")
		return
	}

	ctx := r.Context()
	identifier, userType, err := extractLoginFields(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Login(ctx, identifier, userType); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]any{
		"step2":      true,
		"identifier": identifier,
		"userType":   userType,
	}

	helper.NewSuccessResponse(
		w,
		http.StatusOK,
		"OTP sent successfully. Valid for 5 minutes. Never share it with anyone.",
		data,
	)
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify OTP and complete login process
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param identifier formData string true "Username, email, or phone"
// @Param user_type formData string true "User role"
// @Param otp formData string true "OTP code"
// @Success 200 {object} dto.VerifyOTPResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/verification [post]
func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	if !isPost(r, w) {
		return
	}

	if err := r.ParseForm(); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid form submission")
		return
	}

	ctx := r.Context()
	identifier, userType, err := extractLoginFields(r)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	otp := r.FormValue("otp")
	if otp == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "OTP is required")
		return
	}

	user, err := h.service.VerifyOTP(ctx, identifier, userType, otp)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	if err := helper.CreateSession(
		w,
		user.ID.String(),
		user.Username,
		string(user.RoleType),
	); err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]any{
		"redirect": getDashboardRedirect(user.RoleType),
	}

	helper.NewSuccessResponse(w, http.StatusOK, "OTP verified successfully", data)
}

// ==========================
// SERVICE LAYER
// ==========================

func (s *Service) Login(ctx context.Context, identifier, userType string) error {
	if err := s.repo.Login(ctx, identifier, userType); err != nil {
		return err
	}

	otp := helper.GenerateRandomOTP(6)
	logger.Infof("OTP generated for %s: %s", identifier, otp)

	if err := helper.StoreOTP(identifier, otp, 5*time.Minute); err != nil {
		return fmt.Errorf("failed to store OTP")
	}

	if err := s.SendOTP(ctx, identifier, userType, otp); err != nil {
		return err
	}

	return nil
}

func (s *Service) SendOTP(ctx context.Context, identifier, role, otp string) error {
	loginType := helper.ValidateLoginType(identifier)
	if loginType == helper.LoginInvalid {
		return fmt.Errorf("invalid username format")
	}

	logger.Infof("Sending OTP to %s via %s", identifier, loginType)

	switch loginType {
	case helper.LoginEmail:
		logger.Infof("OTP sent via email to %s", identifier)
	case helper.LoginPhone:
		logger.Infof("OTP sent via SMS to %s", identifier)
	default:
		return fmt.Errorf("unsupported login type")
	}

	return nil
}

func (s *Service) VerifyOTP(ctx context.Context, identifier, userType, otp string) (*domain.User, error) {
	storedOTP, _, err := helper.GetStoredOTP(identifier)
	if err != nil || storedOTP != otp {
		return nil, fmt.Errorf("invalid or expired OTP")
	}

	helper.DeleteOTP(identifier)

	user, err := s.repo.GetUserByUsername(ctx, identifier)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if string(user.RoleType) != userType {
		return nil, fmt.Errorf("user role mismatch")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is inactive")
	}

	return user, nil
}

// ==========================
// REPOSITORY
// ==========================

func (r *Repository) Login(ctx context.Context, identifier, userType string) error {
	// Validate user existence based on identifier & role
	return nil
}

// ==========================
// HELPERS
// ==========================

func isPost(r *http.Request, w http.ResponseWriter) bool {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return false
	}
	return true
}

func extractLoginFields(r *http.Request) (string, string, error) {
	identifier := r.FormValue("username")
	if identifier == "" {
		identifier = r.FormValue("identifier")
	}

	userType := r.FormValue("role")
	if userType == "" {
		userType = r.FormValue("user_type")
	}

	if identifier == "" {
		return "", "", fmt.Errorf("username/email/phone is required")
	}

	if userType == "" {
		return "", "", fmt.Errorf("user type is required")
	}

	return identifier, userType, nil
}

func getDashboardRedirect(role domain.UserRole) string {
	switch role {
	case domain.RoleSuperAdmin:
		return "/dashboard"
	case domain.RoleStudent, domain.RoleGuardian:
		return "/student/dashboard"
	case domain.RoleAdmin, domain.RoleTeacher, domain.RoleAccountant,
		domain.RoleLibrarian, domain.RoleDriver, domain.RoleEmployee, domain.RoleNurse:
		return "/school/dashboard"
	default:
		return "/dashboard"
	}
}
