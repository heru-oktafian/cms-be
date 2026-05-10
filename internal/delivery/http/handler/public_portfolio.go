package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

func (h *Handler) GetPublicPortfolio(c *fiber.Ctx) error {
	h.initProjectUsecase()
	h.initSkillUsecase()
	h.initExperienceUsecase()
	h.initSocialLinkUsecase()

	profile, err := h.profileUsecase.Get()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio profile", nil)
	}

	skills, err := h.skillUsecase.ListActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio skills", nil)
	}

	tools, err := h.toolUsecase.GetAllActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio tools", nil)
	}

	projects, err := h.projectUsecase.ListActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio projects", nil)
	}

	experiences, err := h.experienceUsecase.ListActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio experiences", nil)
	}

	socialLinks, err := h.socialLinkUsecase.ListActive()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch portfolio social links", nil)
	}

	payload := fiber.Map{
		"profile":      profile,
		"skills":       skills,
		"tools":        tools,
		"projects":     projects,
		"experiences":  experiences,
		"social_links": socialLinks,
	}

	return response.JSON(c, http.StatusOK, "ok", payload)
}
