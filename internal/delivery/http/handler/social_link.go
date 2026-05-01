package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	socialLinkRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	socialLinkUsecase "github.com/heru-oktafian/cms-be/internal/usecase/social_link"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type socialLinkRequest struct {
	Platform  string `json:"platform"`
	Label     string `json:"label"`
	URL       string `json:"url"`
	IconPath  string `json:"icon_path"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}

func (h *Handler) initSocialLinkUsecase() {
	if h.socialLinkUsecase == nil {
		socialLinkRepository := socialLinkRepo.NewSocialLinkRepository(h.app.DB)
		h.socialLinkUsecase = socialLinkUsecase.NewUsecase(socialLinkRepository)
	}
}

func (h *Handler) ListPublicSocialLinks(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	socialLinks, err := h.socialLinkUsecase.ListActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch social links", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", socialLinks)
}

func (h *Handler) ListAdminSocialLinks(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	socialLinks, err := h.socialLinkUsecase.List()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch social links", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", socialLinks)
}

func (h *Handler) GetAdminSocialLink(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid social link id", nil)
	}
	socialLink, err := h.socialLinkUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch social link", nil)
	}
	if socialLink == nil {
		return response.JSON(c, http.StatusNotFound, "social link not found", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", socialLink)
}

func (h *Handler) CreateAdminSocialLink(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	socialLink, err := h.parseSocialLinkRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	created, err := h.socialLinkUsecase.Create(socialLink)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to create social link: %v", err), nil)
	}
	return response.JSON(c, http.StatusCreated, "social link created", created)
}

func (h *Handler) UpdateAdminSocialLink(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid social link id", nil)
	}
	existing, err := h.socialLinkUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch social link", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "social link not found", nil)
	}
	payload, err := h.parseSocialLinkRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	payload.ID = existing.ID
	payload.CreatedAt = existing.CreatedAt
	updated, err := h.socialLinkUsecase.Update(payload)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to update social link: %v", err), nil)
	}
	return response.JSON(c, http.StatusOK, "social link updated", updated)
}

func (h *Handler) DeleteAdminSocialLink(c *fiber.Ctx) error {
	h.initSocialLinkUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid social link id", nil)
	}
	existing, err := h.socialLinkUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch social link", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "social link not found", nil)
	}
	if err := h.socialLinkUsecase.Delete(uint(id)); err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to delete social link", nil)
	}
	return response.JSON(c, http.StatusOK, "social link deleted", nil)
}

func (h *Handler) parseSocialLinkRequest(c *fiber.Ctx) (*entity.SocialLink, error) {
	var payload socialLinkRequest
	if err := c.BodyParser(&payload); err != nil {
		return nil, err
	}

	payload.Platform = strings.TrimSpace(payload.Platform)
	payload.Label = strings.TrimSpace(payload.Label)
	payload.URL = strings.TrimSpace(payload.URL)
	payload.IconPath = strings.TrimSpace(payload.IconPath)

	if payload.Platform == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "platform is required")
	}
	if payload.URL == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "url is required")
	}

	return &entity.SocialLink{
		Platform:  payload.Platform,
		Label:     payload.Label,
		URL:       payload.URL,
		IconPath:  payload.IconPath,
		IsActive:  payload.IsActive,
		SortOrder: payload.SortOrder,
	}, nil
}
