package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type countsResponse struct {
	Skills      int `json:"skills"`
	Projects    int `json:"projects"`
	Experiences int `json:"experiences"`
	Messages    int `json:"messages"`
}

func (h *Handler) GetCounts(c *fiber.Ctx) error {
	if h.skillUsecase == nil {
		h.initSkillUsecase()
	}
	if h.projectUsecase == nil {
		h.initProjectUsecase()
	}
	if h.experienceUsecase == nil {
		h.initExperienceUsecase()
	}
	// contactMessageUsecase is initialized in NewHandler

	skillsCount, _ := h.skillUsecase.Count()
	projectsCount, _ := h.projectUsecase.Count()
	experiencesCount, _ := h.experienceUsecase.Count()
	messagesCount, _ := h.contactMessageUsecase.Count()

	counts := countsResponse{
		Skills:      skillsCount,
		Projects:    projectsCount,
		Experiences: experiencesCount,
		Messages:    messagesCount,
	}

	return response.JSON(c, http.StatusOK, "ok", counts)
}