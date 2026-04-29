package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	skillRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	skillUsecase "github.com/heru-oktafian/cms-be/internal/usecase/skill"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type skillRequest struct {
	Name      string `json:"name"`
	Level     string `json:"level"`
	IconPath  string `json:"icon_path"`
	SortOrder int    `json:"sort_order"`
}

func (h *Handler) initSkillUsecase() {
	if h.skillUsecase == nil {
		skillRepository := skillRepo.NewSkillRepository(h.app.DB)
		h.skillUsecase = skillUsecase.NewUsecase(skillRepository)
	}
}

func (h *Handler) ListPublicSkills(c *fiber.Ctx) error {
	h.initSkillUsecase()
	skills, err := h.skillUsecase.List()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch skills", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", skills)
}

func (h *Handler) ListAdminSkills(c *fiber.Ctx) error {
	return h.ListPublicSkills(c)
}

func (h *Handler) GetAdminSkill(c *fiber.Ctx) error {
	h.initSkillUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid skill id", nil)
	}
	skill, err := h.skillUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch skill", nil)
	}
	if skill == nil {
		return response.JSON(c, http.StatusNotFound, "skill not found", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", skill)
}

func (h *Handler) CreateAdminSkill(c *fiber.Ctx) error {
	h.initSkillUsecase()
	skill, err := h.parseSkillRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	created, err := h.skillUsecase.Create(skill)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to create skill: %v", err), nil)
	}
	return response.JSON(c, http.StatusCreated, "skill created", created)
}

func (h *Handler) UpdateAdminSkill(c *fiber.Ctx) error {
	h.initSkillUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid skill id", nil)
	}
	existing, err := h.skillUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch skill", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "skill not found", nil)
	}
	payload, err := h.parseSkillRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	payload.ID = existing.ID
	payload.CreatedAt = existing.CreatedAt
	updated, err := h.skillUsecase.Update(payload)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to update skill: %v", err), nil)
	}
	return response.JSON(c, http.StatusOK, "skill updated", updated)
}

func (h *Handler) DeleteAdminSkill(c *fiber.Ctx) error {
	h.initSkillUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid skill id", nil)
	}
	existing, err := h.skillUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch skill", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "skill not found", nil)
	}
	if err := h.skillUsecase.Delete(uint(id)); err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to delete skill", nil)
	}
	return response.JSON(c, http.StatusOK, "skill deleted", nil)
}

func (h *Handler) parseSkillRequest(c *fiber.Ctx) (*entity.Skill, error) {
	var payload skillRequest
	if err := c.BodyParser(&payload); err != nil {
		return nil, err
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Level = strings.TrimSpace(payload.Level)
	payload.IconPath = strings.TrimSpace(payload.IconPath)

	if payload.Name == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "name is required")
	}

	return &entity.Skill{
		Name:      payload.Name,
		Level:     payload.Level,
		IconPath:  payload.IconPath,
		SortOrder: payload.SortOrder,
	}, nil
}
