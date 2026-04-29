package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

func (h *Handler) Health(c *fiber.Ctx) error {
	sqlDB, err := h.app.DB.DB()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "database handle error", fiber.Map{"status": "unhealthy"})
	}
	if err := sqlDB.Ping(); err != nil {
		return response.JSON(c, http.StatusServiceUnavailable, "database unreachable", fiber.Map{"status": "unhealthy"})
	}
	return response.JSON(c, http.StatusOK, "ok", fiber.Map{"status": "healthy", "database": "up"})
}

func (h *Handler) Placeholder(message string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return response.JSON(c, http.StatusOK, message, fiber.Map{})
	}
}
