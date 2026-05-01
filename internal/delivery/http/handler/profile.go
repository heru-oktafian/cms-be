package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	appctx "github.com/heru-oktafian/cms-be/internal/app"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	profileRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	authUsecase "github.com/heru-oktafian/cms-be/internal/usecase/auth"
	experienceUsecase "github.com/heru-oktafian/cms-be/internal/usecase/experience"
	profileUsecase "github.com/heru-oktafian/cms-be/internal/usecase/profile"
	projectUsecase "github.com/heru-oktafian/cms-be/internal/usecase/project"
	skillUsecase "github.com/heru-oktafian/cms-be/internal/usecase/skill"
	socialLinkUsecase "github.com/heru-oktafian/cms-be/internal/usecase/social_link"
	toolUsecase "github.com/heru-oktafian/cms-be/internal/usecase/tool"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type Handler struct {
	app               *appctx.App
	authUsecase       *authUsecase.Usecase
	profileUsecase    *profileUsecase.Usecase
	projectUsecase    *projectUsecase.Usecase
	skillUsecase      *skillUsecase.Usecase
	experienceUsecase *experienceUsecase.Usecase
	socialLinkUsecase *socialLinkUsecase.Usecase
	toolUsecase       *toolUsecase.UseCase
}

type upsertProfileRequest struct {
	FullName         string `json:"full_name"`
	Headline         string `json:"headline"`
	SubHeadline      string `json:"sub_headline"`
	HeroDescription  string `json:"hero_description"`
	AboutDescription string `json:"about_description"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Location         string `json:"location"`
	AvatarPath       string `json:"avatar_path"`
	ResumePath       string `json:"resume_path"`
}

func NewHandler(app *appctx.App) *Handler {
	profileRepository := profileRepo.NewProfileRepository(app.DB)
	adminUserRepository := profileRepo.NewAdminUserRepository(app.DB)
	toolRepository := profileRepo.NewToolRepository(app.DB)
	return &Handler{
		app:            app,
		authUsecase:    authUsecase.NewUsecase(adminUserRepository, app.Config),
		profileUsecase: profileUsecase.NewUsecase(profileRepository),
		toolUsecase:    toolUsecase.NewUseCase(toolRepository),
	}
}

func (h *Handler) GetPublicProfile(c *fiber.Ctx) error {
	profile, err := h.profileUsecase.Get()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch profile", nil)
	}
	if profile == nil {
		return response.JSON(c, http.StatusOK, "profile not set", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", profile)
}

func (h *Handler) GetAdminProfile(c *fiber.Ctx) error {
	return h.GetPublicProfile(c)
}

func (h *Handler) UpsertAdminProfile(c *fiber.Ctx) error {
	var payload upsertProfileRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid payload", nil)
	}

	payload.FullName = strings.TrimSpace(payload.FullName)
	payload.Headline = strings.TrimSpace(payload.Headline)
	payload.SubHeadline = strings.TrimSpace(payload.SubHeadline)
	payload.HeroDescription = strings.TrimSpace(payload.HeroDescription)
	payload.AboutDescription = strings.TrimSpace(payload.AboutDescription)
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Location = strings.TrimSpace(payload.Location)
	payload.AvatarPath = strings.TrimSpace(payload.AvatarPath)
	payload.ResumePath = strings.TrimSpace(payload.ResumePath)

	if payload.FullName == "" {
		return response.JSON(c, http.StatusBadRequest, "full_name is required", nil)
	}

	profile := &entity.Profile{
		FullName:         payload.FullName,
		Headline:         payload.Headline,
		SubHeadline:      payload.SubHeadline,
		HeroDescription:  payload.HeroDescription,
		AboutDescription: payload.AboutDescription,
		Email:            payload.Email,
		Phone:            payload.Phone,
		Location:         payload.Location,
		AvatarPath:       payload.AvatarPath,
		ResumePath:       payload.ResumePath,
	}

	savedProfile, err := h.profileUsecase.Upsert(profile)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to save profile", nil)
	}
	return response.JSON(c, http.StatusOK, "profile saved", savedProfile)
}
