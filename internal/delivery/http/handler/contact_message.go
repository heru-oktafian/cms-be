package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	contactMessageUsecase "github.com/heru-oktafian/cms-be/internal/usecase/contact_message"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type createContactMessageRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type updateContactMessageRequest struct {
	Status        string `json:"status"`
	FollowUpNotes string `json:"follow_up_notes"`
}

func (h *Handler) CreatePublicContactMessage(c *fiber.Ctx) error {
	var payload createContactMessageRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid payload", nil)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Subject = strings.TrimSpace(payload.Subject)
	payload.Message = strings.TrimSpace(payload.Message)

	if payload.Name == "" {
		return response.JSON(c, http.StatusBadRequest, "name is required", nil)
	}
	if payload.Email == "" {
		return response.JSON(c, http.StatusBadRequest, "email is required", nil)
	}
	if payload.Message == "" {
		return response.JSON(c, http.StatusBadRequest, "message is required", nil)
	}

	message, err := h.contactMessageUsecase.Create(&entity.ContactMessage{
		Name:    payload.Name,
		Email:   payload.Email,
		Phone:   payload.Phone,
		Subject: payload.Subject,
		Message: payload.Message,
	})
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to save contact message", nil)
	}

	return response.JSON(c, http.StatusCreated, "contact message created", message)
}

func (h *Handler) ListAdminContactMessages(c *fiber.Ctx) error {
	messages, err := h.contactMessageUsecase.List()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch contact messages", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", messages)
}

func (h *Handler) GetAdminContactMessage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid contact message id", nil)
	}
	message, err := h.contactMessageUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch contact message", nil)
	}
	if message == nil {
		return response.JSON(c, http.StatusNotFound, "contact message not found", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", message)
}

func (h *Handler) UpdateAdminContactMessage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid contact message id", nil)
	}

	existing, err := h.contactMessageUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch contact message", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "contact message not found", nil)
	}

	var payload updateContactMessageRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid payload", nil)
	}

	payload.Status = strings.TrimSpace(payload.Status)
	payload.FollowUpNotes = strings.TrimSpace(payload.FollowUpNotes)

	if !contactMessageUsecase.IsValidStatus(payload.Status) {
		return response.JSON(c, http.StatusBadRequest, "invalid status", nil)
	}

	existing.Status = payload.Status
	existing.FollowUpNotes = payload.FollowUpNotes

	updated, err := h.contactMessageUsecase.Update(existing)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to update contact message", nil)
	}

	return response.JSON(c, http.StatusOK, "contact message updated", updated)
}
