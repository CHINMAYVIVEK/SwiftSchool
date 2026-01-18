package common

import (
	"context"
	"encoding/json"
	"net/http"
	"swiftschool/domain"
	"swiftschool/helper"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

// ========================= CREATE NOTIFICATION =========================

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create and send a notification to a user
// @Tags Common - Notifications
// @Accept json
// @Produce json
// @Param notification body dto.CreateNotificationRequest true "Notification details"
// @Success 201 {object} dto.SuccessResponse{data=dto.NotificationResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security SessionAuth
// @Router /common/notifications/create [post]
func (h *Handler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.NewErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var notif domain.Notification
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	data, err := h.service.CreateNotification(r.Context(), notif)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusInternalServerError, "failed to create notification: "+err.Error())
		return
	}

	helper.NewSuccessResponse(w, http.StatusCreated, "notification created successfully", data)
}

//////////////////////////////////////////////////////
// ========================= CREATE NOTIFICATION =========================

// SERVICE
func (s *Service) CreateNotification(ctx context.Context, arg domain.Notification) (*domain.Notification, error) {
	notif, err := s.repo.CreateNotification(ctx, arg)
	if err != nil {
		return nil, err
	}
	return notif, nil
}

// REPOSITORY
func (r *Repository) CreateNotification(ctx context.Context, arg domain.Notification) (*domain.Notification, error) {
	return nil, nil
}
