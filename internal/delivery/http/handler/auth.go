package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginAdmin(c *fiber.Ctx) error {
	var payload loginRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid payload", nil)
	}

	result, err := h.authUsecase.Login(strings.TrimSpace(payload.Email), strings.TrimSpace(payload.Password))
	if err != nil {
		return response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSON(c, http.StatusOK, "login success", result)
}
