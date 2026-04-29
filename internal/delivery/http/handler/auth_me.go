package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

func (h *Handler) GetAuthMe(c *fiber.Ctx) error {
	claims, ok := c.Locals("auth_user").(map[string]any)
	if !ok {
		if jwtClaims, ok := c.Locals("auth_user").(fiber.Map); ok {
			return response.JSON(c, http.StatusOK, "ok", jwtClaims)
		}
		return response.JSON(c, http.StatusUnauthorized, "invalid auth context", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", claims)
}
