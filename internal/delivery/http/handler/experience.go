package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	experienceRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	experienceUsecase "github.com/heru-oktafian/cms-be/internal/usecase/experience"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type experienceRequest struct {
	Company     string `json:"company"`
	Position    string `json:"position"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsCurrent   bool   `json:"is_current"`
	SortOrder   int    `json:"sort_order"`
}

func (h *Handler) initExperienceUsecase() {
	if h.experienceUsecase == nil {
		experienceRepository := experienceRepo.NewExperienceRepository(h.app.DB)
		h.experienceUsecase = experienceUsecase.NewUsecase(experienceRepository)
	}
}

func (h *Handler) ListPublicExperiences(c *fiber.Ctx) error {
	h.initExperienceUsecase()
	experiences, err := h.experienceUsecase.List()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch experiences", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", experiences)
}

func (h *Handler) ListAdminExperiences(c *fiber.Ctx) error {
	return h.ListPublicExperiences(c)
}

func (h *Handler) GetAdminExperience(c *fiber.Ctx) error {
	h.initExperienceUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid experience id", nil)
	}
	experience, err := h.experienceUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch experience", nil)
	}
	if experience == nil {
		return response.JSON(c, http.StatusNotFound, "experience not found", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", experience)
}

func (h *Handler) CreateAdminExperience(c *fiber.Ctx) error {
	h.initExperienceUsecase()
	experience, err := h.parseExperienceRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	created, err := h.experienceUsecase.Create(experience)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to create experience: %v", err), nil)
	}
	return response.JSON(c, http.StatusCreated, "experience created", created)
}

func (h *Handler) UpdateAdminExperience(c *fiber.Ctx) error {
	h.initExperienceUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid experience id", nil)
	}
	existing, err := h.experienceUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch experience", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "experience not found", nil)
	}
	payload, err := h.parseExperienceRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	payload.ID = existing.ID
	payload.CreatedAt = existing.CreatedAt
	updated, err := h.experienceUsecase.Update(payload)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to update experience: %v", err), nil)
	}
	return response.JSON(c, http.StatusOK, "experience updated", updated)
}

func (h *Handler) DeleteAdminExperience(c *fiber.Ctx) error {
	h.initExperienceUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid experience id", nil)
	}
	existing, err := h.experienceUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch experience", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "experience not found", nil)
	}
	if err := h.experienceUsecase.Delete(uint(id)); err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to delete experience", nil)
	}
	return response.JSON(c, http.StatusOK, "experience deleted", nil)
}

func (h *Handler) parseExperienceRequest(c *fiber.Ctx) (*entity.Experience, error) {
	var payload experienceRequest
	if err := c.BodyParser(&payload); err != nil {
		return nil, err
	}

	payload.Company = strings.TrimSpace(payload.Company)
	payload.Position = strings.TrimSpace(payload.Position)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.StartDate = strings.TrimSpace(payload.StartDate)
	payload.EndDate = strings.TrimSpace(payload.EndDate)

	if payload.Company == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "company is required")
	}
	if payload.Position == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "position is required")
	}

	return &entity.Experience{
		Company:     payload.Company,
		Position:    payload.Position,
		Description: payload.Description,
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		IsCurrent:   payload.IsCurrent,
		SortOrder:   payload.SortOrder,
	}, nil
}
