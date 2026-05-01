package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetPublicTools(c *fiber.Ctx) error {
	toolHandler := NewToolHandler(h.toolUsecase)
	return toolHandler.GetPublic(c)
}

func (h *Handler) GetAdminTools(c *fiber.Ctx) error {
	toolHandler := NewToolHandler(h.toolUsecase)
	return toolHandler.GetAdmin(c)
}

func (h *Handler) CreateAdminTool(c *fiber.Ctx) error {
	toolHandler := NewToolHandler(h.toolUsecase)
	return toolHandler.Create(c)
}

func (h *Handler) UpdateAdminTool(c *fiber.Ctx) error {
	toolHandler := NewToolHandler(h.toolUsecase)
	return toolHandler.Update(c)
}

func (h *Handler) DeleteAdminTool(c *fiber.Ctx) error {
	toolHandler := NewToolHandler(h.toolUsecase)
	return toolHandler.Delete(c)
}
