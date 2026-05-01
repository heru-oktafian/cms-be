package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	toolusecase "github.com/heru-oktafian/cms-be/internal/usecase/tool"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type ToolHandler struct {
	useCase *toolusecase.UseCase
}

type toolRequest struct {
	Name      string `json:"name"`
	IconPath  string `json:"icon_path"`
	URL       string `json:"url"`
	SortOrder int    `json:"sort_order"`
	IsActive  bool   `json:"is_active"`
}

func NewToolHandler(useCase *toolusecase.UseCase) *ToolHandler {
	return &ToolHandler{useCase: useCase}
}

func (h *ToolHandler) GetPublic(c *fiber.Ctx) error {
	tools, err := h.useCase.GetAllActive()
	if err != nil {
		return response.JSON(c, fiber.StatusInternalServerError, "failed to fetch tools", nil)
	}
	return response.JSON(c, fiber.StatusOK, "ok", tools)
}

func (h *ToolHandler) GetAdmin(c *fiber.Ctx) error {
	tools, err := h.useCase.GetAll()
	if err != nil {
		return response.JSON(c, fiber.StatusInternalServerError, "failed to fetch tools", nil)
	}
	return response.JSON(c, fiber.StatusOK, "ok", tools)
}

func (h *ToolHandler) Create(c *fiber.Ctx) error {
	var req toolRequest
	if err := c.BodyParser(&req); err != nil {
		return response.JSON(c, fiber.StatusBadRequest, "invalid request body", nil)
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return response.JSON(c, fiber.StatusBadRequest, "name is required", nil)
	}
	tool := &entity.Tool{Name: req.Name, IconPath: strings.TrimSpace(req.IconPath), URL: strings.TrimSpace(req.URL), SortOrder: req.SortOrder, IsActive: req.IsActive}
	if err := h.useCase.Create(tool); err != nil {
		return response.JSON(c, fiber.StatusInternalServerError, "failed to create tool", nil)
	}
	return response.JSON(c, fiber.StatusCreated, "tool created", tool)
}

func (h *ToolHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.JSON(c, fiber.StatusBadRequest, "invalid id", nil)
	}
	var req toolRequest
	if err := c.BodyParser(&req); err != nil {
		return response.JSON(c, fiber.StatusBadRequest, "invalid request body", nil)
	}
	tool := &entity.Tool{BaseModel: entity.BaseModel{ID: uint(id)}, Name: strings.TrimSpace(req.Name), IconPath: strings.TrimSpace(req.IconPath), URL: strings.TrimSpace(req.URL), SortOrder: req.SortOrder, IsActive: req.IsActive}
	if err := h.useCase.Update(tool); err != nil {
		return response.JSON(c, fiber.StatusInternalServerError, "failed to update tool", nil)
	}
	return response.JSON(c, fiber.StatusOK, "tool updated", tool)
}

func (h *ToolHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.JSON(c, fiber.StatusBadRequest, "invalid id", nil)
	}
	if err := h.useCase.Delete(uint(id)); err != nil {
		return response.JSON(c, fiber.StatusInternalServerError, "failed to delete tool", nil)
	}
	return response.JSON(c, fiber.StatusOK, "tool deleted", nil)
}
